package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	chat "github.com/arivlav/chat-server/pkg/chat_v1"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer func() {
		if conn.Close() != nil {
			fmt.Println("Error when closing:", err)
		}
	}()

	c := chat.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//create...
	createReq, createErr := c.Create(ctx, &chat.CreateRequest{Usernames: []string{gofakeit.Name(), gofakeit.Name(), gofakeit.Name()}})
	if createErr != nil {
		log.Fatalf("failed to create new chat: %v", createErr)
	}
	log.Printf(color.RedString("New chat have id:\n"), color.GreenString("%+v", createReq.GetId()))

	// sendMessage...
	_, messageErr := c.SendMessage(ctx, &chat.SendMessageRequest{From: gofakeit.Name(), Text: gofakeit.HipsterWord()})
	if messageErr != nil {
		log.Fatalf("failed to send message: %v", messageErr)
	}
	log.Printf(color.RedString("Message is send\n"))

	//delete...
	_, deleteErr := c.Delete(ctx, &chat.DeleteRequest{Id: gofakeit.Uint64()})
	if deleteErr != nil {
		log.Fatalf("failed to delete chat by id: %v", deleteErr)
	}
	log.Printf(color.RedString("Chat is deleted\n"))
}
