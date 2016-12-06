package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
	BotID string
)

type MyStruct struct {
	cardMap map[string]Card
}

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

var cards []Card

func main() {

	cardMap := make(map[string]Card)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	file, err := ioutil.ReadFile("./sw-destiny-awakenings.json")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	json.Unmarshal(file, &cards)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	// log.Printf("Result: %#v\n", msg)

	// for k := range cards {
	// cards[k].Text = strings.Replace(cards[k].Text, "<b>", "``", -1)
	// cards[k].Text = strings.Replace(cards[k].Text, "</b>", "``", -1)
	// cards[k].Text = strings.Replace(cards[k].Text, "<i>", "*", -1)
	// cards[k].Text = strings.Replace(cards[k].Text, "</i>", "*", -1)
	// cards[k].Text = strings.Replace(cards[k].Text, "[ranged]", ":zranged:", -1)
	// cards[k].Text = strings.Replace(cards[k].Text, "[melee]", ":zmelee", -1)
	// cards[k].Text = strings.Replace(cards[k].Text, "[special]", ":zspecial", -1)

	// for s := range cards[k].Sides {
	// 	cards[k].Sides[s] = strings.Replace(cards[k].Sides[s], "RD", ":zranged:", -1)
	// 	cards[k].Sides[s] = strings.Replace(cards[k].Sides[s], "R", ":zresource:", -1)
	// 	cards[k].Sides[s] = strings.Replace(cards[k].Sides[s], "MD", ":zmelee:", -1)
	// 	cards[k].Sides[s] = strings.Replace(cards[k].Sides[s], "Dc", ":zdiscard:", -1)
	// 	cards[k].Sides[s] = strings.Replace(cards[k].Sides[s], "Dr", ":zdisrupt:", -1)
	// 	cards[k].Sides[s] = strings.Replace(cards[k].Sides[s], "F", ":zfocus:", -1)
	// 	cards[k].Sides[s] = strings.Replace(cards[k].Sides[s], "Sh", ":zshield:", -1)
	// 	cards[k].Sides[s] = strings.Replace(cards[k].Sides[s], "Sp", ":zspecial", -1)

	// }
	// }

	for _, v := range cards {
		x := fmt.Sprintf("%v%03d", v.SetCode, v.Position)
		cardMap[x] = v
	}

	MyStruct := &MyStruct{cardMap: cardMap}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(MyStruct.messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func (my *MyStruct) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}
	var x string
	var search []Card
	// If the message is "ping" reply with "Pong!"
	if !strings.HasPrefix(m.Content, "!botcard ") {
		return
	}
	x = strings.TrimPrefix(m.Content, "!botcard ")

	// This block is probably not needed due to returning if the prefix doesn't exist.
	// if strings.HasPrefix(x, "!botcard ") {
	// 	x = strings.TrimPrefix(x, "!botcard ")
	// }

	if x == ":eggplant:" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "That is an eggplant, not a card!")
		return
	}

	for _, v := range cards {
		if strings.Contains(v.Name, x) || strings.Contains(v.Name, strings.Title(x)) {
			search = append(search, v)
		}

	}

	if len(search) <= 0 {
		search = append(search, my.cardMap[x])
	} else if len(search) >= 2 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Please try the following: ")
		for _, v := range search {
			if v.Subtitle == "" {
				_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v%03d - %v\n", v.SetCode, v.Position, v.Name))
			} else {
				_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v%03d - %v - %v\n", v.SetCode, v.Position, v.Name, v.Subtitle))
			}
		}
		return
	}

	fmt.Println(search[0].String())

	// if err != nil {
	// 	log.Fatalf("Something messed up: %v", err)
	// }
	_, _ = s.ChannelMessageSend(m.ChannelID, search[0].String())
}
