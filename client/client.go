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

var g int32 = 2 // generator
var h int32 = 3 // ???
var p int32 = 5 // group prime???
var myRoll int32
var recievedRoll int32
var commitment int32
var r int32

type clientHandle struct {
	stream     Chat.ChattingService_JoinChatClient
	Id         int32
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
	ch.Id = rand.Int31()

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

		if len(clientMessage) > 0 && len(clientMessage) <= 128 {

			if strings.HasPrefix(clientMessage, "compute") {

				log.Printf("My commitment was %d", commitment)

				msg := &Chat.ClientRevelation{
					Name:   ch.clientName,
					C:      commitment,
					M:      myRoll,
					R:      r,
					Result: ch.computeRoll(),
				}
				_, err = client.RevealAll(context.Background(), msg)
				if err != nil {
					log.Printf("Error while sending to server :: %v", err)
				}

			} else {
				var formattedMessage string

				if strings.HasPrefix(clientMessage, "roll and send commitment") {
					formattedMessage = fmt.Sprintf("My commitment is %d", ch.commitment())

				} else if strings.HasPrefix(clientMessage, "roll") {
					formattedMessage = fmt.Sprintf("I rolled %d", ch.rollDice())

				} else if strings.HasPrefix(clientMessage, "validate") {
					formattedMessage = ch.validate()

				} else {
					formattedMessage = clientMessage
				}

				msg := &Chat.ClientContent{
					Name: ch.clientName,
					Body: formattedMessage,
				}
				_, err = client.SendContent(context.Background(), msg)
				if err != nil {
					log.Printf("Error while sending to server :: %v", err)
				}
			}

			time.Sleep(500 * time.Millisecond)

		} else {
			log.Print("Your message must be between 1 and 128 characters!")
		}
	}
}

func (ch *clientHandle) rollDice() int32 {
	return rand.Int31n(6) + 1
}

func (ch *clientHandle) commitment() int32 {
	myRoll = ch.rollDice()

	log.Printf("I rolled %d!", myRoll)
	return generatePedersenCommitment(myRoll)
	//return generateHashBasedCommitment(roll)
}

func (ch *clientHandle) computeRoll() int32 {
	return int32(math.Mod(float64(myRoll)+float64(recievedRoll), 6.0) + 1.0)
}

func (ch *clientHandle) validate() string {
	if math.Pow(float64(g), float64(recievedRoll))*math.Pow(float64(h), float64(r)) == float64(commitment) {
		return "I trust this result."
	} else {
		return fmt.Sprintf("I don't trust this result. %f is not the same as %f", math.Pow(float64(g), float64(recievedRoll))*math.Pow(float64(h), float64(r)), float64(commitment))
	}
}

func generatePedersenCommitment(m int32) int32 {
	r = rand.Int31n(p) // random element in the group
	log.Printf("Random element r: %d", r)

	commitment = int32(math.Pow(float64(g), float64(m)) * math.Pow(float64(h), float64(r)))
	return commitment
}

// func generateHashBasedCommitment(roll int32) int {
// 	var r = rand.Uint64() //random bit
// 	//...
// }

func (ch *clientHandle) receiveMessage() {

	for {
		resp, err := ch.stream.Recv()
		if err != nil {
			log.Fatalf("can not receive %v", err)
		}

		if strings.HasPrefix(resp.Body, "[") && strings.HasSuffix(resp.Body, "]") {

			s := resp.Body[1 : len(resp.Body)-1]
			var revealed = strings.Split(s, ", ")

			c, _ := strconv.ParseInt(revealed[0], 10, 32)
			commitment = int32(c)

			rr, _ := strconv.ParseInt(revealed[1], 10, 32)
			recievedRoll = int32(rr)

			r0, _ := strconv.ParseInt(revealed[2], 10, 32)
			r = int32(r0)

			r1, _ := strconv.ParseInt(revealed[3], 10, 32)
			result := int(r1)

			log.Printf("%s : %d, %d, %d, %d", resp.Name, commitment, recievedRoll, r, result)

		} else {
			log.Printf("%s : %s", resp.Name, resp.Body)

			if strings.Contains(resp.Body, "I rolled") {
				arr := strings.Split(resp.Body, " ")
				rRoll, err := strconv.ParseInt(arr[len(arr)-1], 0, 32)
				if err != nil {
					log.Fatalf("can not receive %v", err)
				}
				recievedRoll = int32(rRoll)
			}
		}
	}
}
