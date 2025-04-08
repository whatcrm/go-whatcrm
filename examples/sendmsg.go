package main

import (
	"context"
	"fmt"
	gowhatcrm "github.com/whatcrm/go-whatcrm"
	"github.com/whatcrm/go-whatcrm/models"
)

func main() {

	redirectURI := "https://dev.whatcrm.net/instances/%s/sendMessage"

	header := ""
	token := ""
	headerChat := ""
	chatToken := ""

	client, err := gowhatcrm.NewClient(redirectURI, header, token, &headerChat, &chatToken)
	if err != nil {
		fmt.Println("error: ", err)
	}

	ctx := context.Background()
	chatKey := ""
	chatID := ""
	body := ""

	msg := models.MessageInput{
		Body:   body,
		ChatID: chatID,
	}

	res, err := client.SendMessage(ctx, chatKey, msg)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(res)
}
