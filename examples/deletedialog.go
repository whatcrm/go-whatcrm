package main

import (
	"context"
	"fmt"
	gowhatcrm "github.com/whatcrm/go-whatcrm"
	"github.com/whatcrm/go-whatcrm/models"
)

func main() {
	redirectURI := "https://dev.whatcrm.net/instances/%s/removeChat"

	header := ""
	token := ""
	headerChat := ""
	chatToken := ""
	chatKey := ""
	chatID := ""

	client, err := gowhatcrm.NewClient(redirectURI, header, token, &headerChat, &chatToken)
	if err != nil {
		fmt.Println("error: ", err)
	}

	ctx := context.Background()

	msg := models.Dialog{
		ChatID: chatID,
	}

	res, err := client.DeleteDialog(ctx, chatKey, msg)
	if err != nil {
		fmt.Println("error: ", err)
	}

	if err != nil {
		fmt.Println("Error formatted:", err)
		return
	}
	fmt.Println(res)
}
