package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)


type Command struct {
	Name        string
	Description string
	Execute     func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}


var commands = []*Command{
	{
		Name:        "ping",
		Description: "Replies with Pong!",
		Execute: func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		},
	},
	{
		Name:        "echo",
		Description: "Echoes back the message.",
		Execute: func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
			if len(args) > 0 {
				s.ChannelMessageSend(m.ChannelID, strings.Join(args, " "))
			} else {
				s.ChannelMessageSend(m.ChannelID, "Please provide a message to echo.")
			}
		},
	},
}

func main() {
	Token := "YOUR_BOT_TOKEN" 

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddMessageCreate(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	if m.Author.ID == s.State.User.ID {
		return
	}

	
	if strings.HasPrefix(m.Content, "!") {
		args := strings.Fields(m.Content[1:]) 
		commandName := args[0]
		args = args[1:]

		for _, cmd := range commands {
			if cmd.Name == commandName {
				cmd.Execute(s, m, args)
				return
			}
		}

		s.ChannelMessageSend(m.ChannelID, "Unknown command: "+commandName)
	}
}
