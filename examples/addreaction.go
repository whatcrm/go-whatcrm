package main

import (
	"context"
	"fmt"
	gowhatcrm "github.com/whatcrm/go-whatcrm"
	"github.com/whatcrm/go-whatcrm/models"
)

func main() {

	redirectURI := "https://dev.whatcrm.net/instances/%s/reaction"

	header := ""
	token := ""
	headerChat := ""
	chatToken := ""
	chatKey := ""

	client, err := gowhatcrm.NewClient(redirectURI, header, token, &headerChat, &chatToken)
	if err != nil {
		fmt.Println("error: ", err)
	}

	ctx := context.Background()

	request := models.Reaction{
		MessageID: "",
		ChatID:    "",
		Reaction:  "❤",
	}

	res, err := client.AddReactionToMessage(ctx, chatKey, request)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(res)
}
