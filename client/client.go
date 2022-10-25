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

func encrypt(m int64) (int64, int64) {
	//r = rand.Int31n(p - 1)
	c1 := int64(math.Mod(math.Pow(float64(g), float64(privateKey)), float64(p)))

	i := math.Mod(math.Pow(float64(othersPublicKey), float64(privateKey)), float64(p))
	c2 := int64(m) * int64(i) % int64(p)
	return c1, c2
}

func decrypt(c1 int64, c2 int64) int64 {
	s := int64(math.Pow(float64(c1), float64(privateKey))) % p
	inv := int64(math.Pow(float64(s), float64(p-2))) % p
	return c2 * inv % p
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
			formattedMsg := ch.formatArrayMessage()
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
		} else {
			log.Print("I didn't get that message. You can write the following exchange:\nmake commitment\nsend roll\nreveal commitment\nvalidate commitment\n")
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func (ch *clientHandle) formatArrayMessage() *Chat.ClientContent {
	log.Printf("My commitment was %d", commitment)
	enCom := formatTouple(encrypt(commitment))
	enRoll := formatTouple(encrypt(myRoll))
	enRan := formatTouple(encrypt(randCE))

	formattedBody := fmt.Sprintf("[%s, %s, %s]", enCom, enRoll, enRan)
	msg := &Chat.ClientContent{
		Name: ch.clientName,
		Body: formattedBody,
	}
	return msg
}

func formatTouple(i int64, j int64) string {
	return fmt.Sprintf("(%d,%d)", i, j)
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

func computeRoll() int64 {
	return int64(math.Mod(float64(myRoll)+float64(recievedRoll), 6.0) + 1.0)
}

func (ch *clientHandle) validate() {
	if math.Pow(float64(g), float64(recievedRoll))*math.Pow(float64(h), float64(randCE)) == float64(commitment) {
		log.Printf("I have validated the commitment and agreed. The shared roll is %d", computeRoll())
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

func unpackTouple(msg string) (int64, int64) {
	vals := strings.Split(msg[1:len(msg)-1], ",")
	c1, err := strconv.ParseInt(vals[0], 0, 32)
	if err != nil {
		log.Fatalf("can not unpack c1 %v", err)
	}
	c2, err := strconv.ParseInt(vals[1], 0, 32)
	if err != nil {
		log.Fatalf("can not unpack c2: %v", err)
	}
	return c1, c2
}

func (ch *clientHandle) receiveMessage() {

	for {
		resp, err := ch.stream.Recv()
		if err != nil {
			log.Fatalf("can not receive %v", err)
		}
		log.Printf("%s : %s", resp.Name, resp.Body)

		if strings.HasPrefix(resp.Body, "(") && strings.HasSuffix(resp.Body, ")") {
			c1, c2 := unpackTouple(resp.Body)
			log.Printf("Decrypted message: %d", decrypt(c1, c2))

		} else if strings.HasPrefix(resp.Body, "[") && strings.HasSuffix(resp.Body, "]") {
			trimmedS := resp.Body[1 : len(resp.Body)-1]
			touples := strings.Split(trimmedS, ", ")

			commitment = decrypt(unpackTouple(touples[0]))
			recievedRoll = decrypt(unpackTouple(touples[1]))
			randCE = decrypt(unpackTouple(touples[2]))
			log.Printf("Decrypted message: %d, %d, %d", commitment, recievedRoll, randCE)
		}
	}
}
