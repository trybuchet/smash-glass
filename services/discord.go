package services

import (
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

type DiscordService struct {
	clientId   string
	InActivity bool
}

func NewDiscordService(clientId string) *DiscordService {
	return &DiscordService{
		clientId:   clientId,
		InActivity: false,
	}
}

// Start rich presence thread
func (ds *DiscordService) StartActivity(players int, game string, link string) {

	err := client.Login(ds.clientId)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:   game,
		Details: "I'm hosting '" + game + "' on Soda Arcade",
		Party: &client.Party{
			ID:      "-1",
			Players: players,
		},
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			{
				Label: "Join",
				Url:   link,
			},
		},
	})
	if err != nil {
		fmt.Println("Error setting activity:", err)
		return
	} else {
		fmt.Println("Discord activity started")
	}

}

// Stop rich presence thread
func (ds *DiscordService) StopActivity() {

	client.Logout()
	fmt.Println("Discord activity stopped")

}
