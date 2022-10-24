package main

import (
	Chat "Golang_Chat_System/Chat"
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

var g int64 = 666  // generator
var h int64 = 3    // also generator
var p int64 = 6661 // group prime

var myRoll int64
var recievedRoll int64
var commitment int64
var randCE int64

var privateKey int64 = 7
var othersPublicKey int64 = 637

type clientHandle struct {
	stream     Chat.ChattingService_JoinChatClient
	Id         int64
	clientName string
}

func main() {
	const serverID = "localhost:8007"

	log.Println("Connecting : " + serverID)
	conn, err := grpc.Dial(serverID, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect gRPC server :: %v", err)
	}
	defer conn.Close()

	client := Chat.NewChattingServiceClient(conn)

	ch := clientHandle{}
	ch.clientConfig()

	rand.Seed(time.Now().UnixNano())
	ch.Id = rand.Int63()

	var user = &Chat.User{
		Id:   ch.Id,
		Name: ch.clientName,
	}

	_stream, err := client.JoinChat(context.Background(), user)
	if err != nil {
		log.Fatalf("Failed to get response from gRPC server :: %v", err)
	}

	ch.stream = _stream

	go ch.sendMessage(client)
	go ch.receiveMessage()

	// block main
	bl := make(chan bool)
	<-bl
}

func (ch *clientHandle) clientConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Your Name : ")
	msg, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from console :: %v", err)
	}
	ch.clientName = strings.TrimRight(msg, "\r\n")
}

func (ch *clientHandle) sendMessage(client Chat.ChattingServiceClient) {

	for {
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')
		clientMessage = strings.TrimRight(clientMessage, "\r\n")
		if err != nil {
			log.Printf("Failed to read from console : %v", err)
			continue
		}

		if strings.HasPrefix(clientMessage, "reveal commitment") {
			formattedMsg := ch.formatMessageOfImportantInformation()
			_, err = client.SendContent(context.Background(), formattedMsg)
			if err != nil {
				log.Printf("Error while sending to server :: %v", err)
			}

		} else if strings.HasPrefix(clientMessage, "validate commitment") {
			ch.validate()

		} else if strings.HasPrefix(clientMessage, "make commitment") || strings.HasPrefix(clientMessage, "send roll") {
			var message int64
			if strings.HasPrefix(clientMessage, "make commitment") {
				message = ch.commitment()

			} else if strings.HasPrefix(clientMessage, "send roll") {
				message = ch.rollDice()
			}
			c1, c2 := encrypt(message)

			msg := &Chat.ClientEncrypted{
				Name: ch.clientName,
				C1:   c1,
				C2:   c2,
			}
			_, err = client.SendEncrypted(context.Background(), msg)
			if err != nil {
				log.Printf("Error while sending to server :: %v", err)
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func (ch *clientHandle) formatMessageOfImportantInformation() *Chat.ClientContent {
	log.Printf("My commitment was %d", commitment)
	formattedBody := fmt.Sprintf("[%d, %d, %d, %d]", commitment, myRoll, randCE, ch.computeRoll())
	msg := &Chat.ClientContent{
		Name: ch.clientName,
		Body: formattedBody,
	}
	return msg
}

func (ch *clientHandle) rollDice() int64 {
	roll := rand.Int63n(6) + 1
	log.Printf("I rolled %d!", roll)
	return roll
}

func (ch *clientHandle) commitment() int64 {
	myRoll = ch.rollDice()
	return generatePedersenCommitment(myRoll)
}

func (ch *clientHandle) computeRoll() int64 {
	return int64(math.Mod(float64(myRoll)+float64(recievedRoll), 6.0) + 1.0)
}

func (ch *clientHandle) validate() {
	if math.Pow(float64(g), float64(recievedRoll))*math.Pow(float64(h), float64(randCE)) == float64(commitment) {
		log.Print("I trust this result.")
	} else {
		log.Printf("I don't trust this result. %f is not the same as %f", math.Pow(float64(g), float64(recievedRoll))*math.Pow(float64(h), float64(randCE)), float64(commitment))
	}
}

func generatePedersenCommitment(m int64) int64 {
	randCE = rand.Int63n(p) // random element in the group
	log.Printf("Random element r: %d", randCE)
	commitment = int64(math.Pow(float64(g), float64(m)) * math.Pow(float64(h), float64(randCE)))
	log.Printf("Commitment c: %d", commitment)
	return commitment
}

func (ch *clientHandle) receiveMessage() {

	for {
		resp, err := ch.stream.Recv()
		if err != nil {
			log.Fatalf("can not receive %v", err)
		}

		if strings.HasPrefix(resp.Body, "(") && strings.HasSuffix(resp.Body, ")") {
			log.Printf("%s : %s", resp.Name, resp.Body)

			vals := strings.Split(resp.Body[1:len(resp.Body)-1], ",")
			c1, _ := strconv.ParseInt(vals[0], 0, 32)
			c2, err := strconv.ParseInt(vals[1], 0, 32)
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			decryptAndPrint(c1, c2)
		} else if strings.HasPrefix(resp.Body, "[") && strings.HasSuffix(resp.Body, "]") {

			unpackAndPrintInformation(resp.Name, resp.Body)

		}
	}
}

func unpackAndPrintInformation(name string, body string) {
	s := body[1 : len(body)-1]
	var revealed = strings.Split(s, ", ")

	c, _ := strconv.ParseInt(revealed[0], 10, 32)
	commitment = int64(c)

	rr, _ := strconv.ParseInt(revealed[1], 10, 32)
	recievedRoll = int64(rr)

	r0, _ := strconv.ParseInt(revealed[2], 10, 32)
	randCE = int64(r0)

	r1, _ := strconv.ParseInt(revealed[3], 10, 32)
	result := int(r1)

	log.Printf("%s : %d, %d, %d, %d", name, commitment, recievedRoll, randCE, result)
}

func encrypt(m int64) (int64, int64) {
	//r = rand.Int31n(p - 1)
	c1 := int64(math.Mod(math.Pow(float64(g), float64(privateKey)), float64(p)))

	i := math.Mod(math.Pow(float64(othersPublicKey), float64(privateKey)), float64(p))
	c2 := int64(m) * int64(i) % int64(p)
	return c1, c2
}

func decryptAndPrint(c1 int64, c2 int64) {
	s := int64(math.Pow(float64(c1), float64(privateKey))) % p
	inv := int64(math.Pow(float64(s), float64(p-2))) % p
	m := c2 * inv % p
	log.Printf("Decrypted message: %d", m)
}
