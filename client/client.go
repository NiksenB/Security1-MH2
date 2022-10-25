package main

import (
	Chat "Golang_Chat_System/Chat"
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
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

var privateKey int64 = 66
var othersPublicKey int64 = 2227

type clientHandle struct {
	stream     Chat.ChattingService_JoinChatClient
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

	var user = &Chat.User{
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

func SHA256FromInt(message int64) []byte {
	s := fmt.Sprintf("%d", message)
	return SHA256FromString(s)
}

func SHA256FromString(message string) []byte {
	h := sha256.New()
	h.Write([]byte(message))
	return h.Sum(nil)
}

func validateSHAFromInt(decrMsg int64, sig []byte) int {
	digest := SHA256FromInt(decrMsg)
	return bytes.Compare(digest, sig)
}

func validateSHAFromString(decrMsg string, sig []byte) int {
	digest := SHA256FromString(decrMsg)
	return bytes.Compare(digest, sig)
}

func encryptMsg(m int64) (int64, int64) {
	//r := rand.Int63n(p - 1)
	c1 := modPow(g, privateKey, p)
	k := modPow(othersPublicKey, privateKey, p) // shared secret
	c2 := m * k % p
	return c1, c2
}

func decryptMsg(c1 int64, c2 int64) int64 {
	s := modPow(c1, privateKey, p)
	inv := modPow(s, p-2, p)
	m := c2 * inv % p
	return m
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
			valueString := fmt.Sprintf("[%d, %d, %d]", commitment, myRoll, randCE)
			log.Printf("I'm sending these three values (c, roll, r): %s", valueString)
			signature := SHA256FromString(valueString)
			enCom := formatTouple(encryptMsg(commitment))
			enRoll := formatTouple(encryptMsg(myRoll))
			enRan := formatTouple(encryptMsg(randCE))
			formattedBody := fmt.Sprintf("[%s, %s, %s]", enCom, enRoll, enRan)
			msg := &Chat.ClientEncrypted{
				Name:      ch.clientName,
				Message:   formattedBody,
				Signature: signature,
			}
			_, err = client.SendEncrypted(context.Background(), msg)
			if err != nil {
				log.Printf("Error while sending to server :: %v", err)
			}

		} else if strings.HasPrefix(clientMessage, "make commitment") || strings.HasPrefix(clientMessage, "send roll") {
			var message int64
			if strings.HasPrefix(clientMessage, "make commitment") {
				message = ch.commitment()

			} else if strings.HasPrefix(clientMessage, "send roll") {
				message = ch.rollDice()
			}
			c1, c2 := encryptMsg(message)
			m := fmt.Sprintf("(%d,%d)", c1, c2)
			bs := SHA256FromInt(message)

			msg := &Chat.ClientEncrypted{
				Name:      ch.clientName,
				Message:   m,
				Signature: bs,
			}
			_, err = client.SendEncrypted(context.Background(), msg)
			if err != nil {
				log.Printf("Error while sending to server :: %v", err)
			}

		} else if strings.HasPrefix(clientMessage, "validate commitment") {
			ch.validate()
		} else {
			log.Print("I didn't understand that message. You should only write the following exchange:\nA: make commitment\nB: send roll\nA: reveal commitment\nB: validate commitment\n\n")
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func (ch *clientHandle) receiveMessage() {

	for {
		resp, err := ch.stream.Recv()
		if err != nil {
			log.Fatalf("can not receive %v", err)
		}
		fmt.Printf("\n%s : %s (signed)\n", resp.Name, resp.Body)

		sig := resp.Signature
		var sigValidation int

		if strings.HasPrefix(resp.Body, "(") && strings.HasSuffix(resp.Body, ")") {
			c1, c2 := unpackTouple(resp.Body)
			msgContent := decryptMsg(c1, c2)
			sigValidation = validateSHAFromInt(msgContent, sig)
			log.Printf("Decrypted message: %d", msgContent)

		} else if strings.HasPrefix(resp.Body, "[") && strings.HasSuffix(resp.Body, "]") {
			trimmedS := resp.Body[1 : len(resp.Body)-1]
			touples := strings.Split(trimmedS, ", ")
			commitment = decryptMsg(unpackTouple(touples[0]))
			recievedRoll = decryptMsg(unpackTouple(touples[1]))
			randCE = decryptMsg(unpackTouple(touples[2]))
			msgContent := fmt.Sprintf("[%d, %d, %d]", commitment, recievedRoll, randCE)
			sigValidation = validateSHAFromString(msgContent, sig)
			log.Printf("Decrypted message (c, roll, r): %d, %d, %d", commitment, recievedRoll, randCE)
		}

		if sigValidation == 0 {
			log.Printf("Signature has been validated.\n\n")
		} else {
			log.Printf("Signature could not be validated\n\n")
		}
	}
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
	//comCheck := math.Pow(float64(g), float64(recievedRoll))*math.Pow(float64(h), float64(randCE))
	comCheck := modPow(g, recievedRoll, p) * modPow(h, randCE, p) % p
	if comCheck == commitment {
		log.Printf("I have validated the commitment and agreed. The shared roll is then: %d", computeRoll())
	} else {
		log.Printf("I don't trust this result. %f is not the same as %f", math.Pow(float64(g), float64(recievedRoll))*math.Pow(float64(h), float64(randCE)), float64(commitment))
	}
}

func generatePedersenCommitment(m int64) int64 {
	randCE = rand.Int63n(p) // random element in the group
	log.Printf("Random element r: %d", randCE)
	//commitment = int64(math.Pow(float64(g), float64(m)) * math.Pow(float64(h), float64(randCE)))
	commitment = (modPow(g, m, p) * modPow(h, randCE, p)) % p
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

func modPow(x int64, r int64, p int64) int64 {
	sum := float64(x)
	for i := int64(1); i < r; i++ {
		sum = math.Mod(float64(x)*sum, float64(p))
	}
	return int64(sum)
}
