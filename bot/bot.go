package bot

import (
	"fmt"
	"hello_bot/config"

	// for discord commands
	"github.com/bwmarrin/discordgo"
)

// responsible for sending message

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {
	//start the sesson

	// 	New creates a new Discord session with provided token. If the token is for a bot, it must be prefixed with "Bot "
	// e.g. "Bot ..."
	goBot, err := discordgo.New("Bot " + config.Token)

	config.ErrorCheck(err)

	// User returns the user details of the given userID userID : A user ID or "@me" which is a shortcut of current user ID
	user, err := goBot.User("@me")

	config.ErrorCheck(err)

	BotID = user.ID

	//form discord libraby(AddHandeler)
	goBot.AddHandler(messageHandeler)

	err = goBot.Open()

	config.ErrorCheck(err)

	fmt.Println("Bot is running")
}

func messageHandeler(s *discordgo.Session, m *discordgo.MessageCreate) {

	//commands are not confused as ans
	if m.Author.ID == BotID {
		return // don't do any thing
	}

	if m.Content == "hello" || m.Content == "Hello" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hello "+m.Author.Username)
	}
	if m.Content == "time" || m.Content == "Time" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Time is good now ")

	}
}

func Time(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	// presentTime := time.Now().String()
	// //fmt.Println("Time of now", presentTime)
	// if m.Content == "Tell the time" {
	// 	s.ChannelMessageSend(m.ChannelID, "Time is %v", presentTime)
	// }
}
