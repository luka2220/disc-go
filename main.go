package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/luka2220/discGo/services"
	"github.com/luka2220/discGo/utils"
)

// Global State
var (
	CityGlobal string
)

func main() {
	err := godotenv.Load()

	var (
		TOKEN      = os.Getenv("TOKEN")
		GUILD_ID   = os.Getenv("GUILD_ID")
		CHANNEL_ID = os.Getenv("CHANNEL_ID")
	)

	dg, err := discordgo.New("Bot " + TOKEN)

	if err != nil {
		fmt.Println("Error creating discord session: ", err)
		return
	}

	dg.ChannelMessageSend(CHANNEL_ID, "Bot Online")

	// Create the commands
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "weather",
			Description: "Get weather from location",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "city",
					Description: "Tempurature of city",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
	}

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"weather": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var city string

			// Loop through the options to find the "city" option
			for _, option := range i.ApplicationCommandData().Options {
				if option.Name == "city" {
					city = option.Value.(string)

					// Set the global state CityWeather
					utils.SetCityWeather(city)

					// Check if services package recieved state update
					services.TestCity()

					break
				}
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Getting the weather for %s...", city),
				},
			})
		},
	}

	// store middlewares
	middleware := SendMessageHandler(CHANNEL_ID)
	newMessage := SendMessage(CHANNEL_ID)

	// Add message event handler
	dg.AddHandler(middleware) // middleware function with channel ID passed in
	dg.AddHandler(newMessage)
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	// Open websocket connection to discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection to discord: ", err)
		return
	}

	// Initialize the commands
	_, err = dg.ApplicationCommandCreate(dg.State.User.ID, GUILD_ID, commands[0])
	if err != nil {
		fmt.Println("Error occured creating slash command: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

// Middleware function to return an event handler
func SendMessage(chan_id string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore message created by the bot
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "ping" {
			s.ChannelMessageSend(chan_id, "pong")
		}

	}
}

// Middleware function to return an event handler
func SendMessageHandler(chan_id string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore message created by the bot
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "check" {
			s.ChannelMessageSend(chan_id, "middleware function executed")
		}

	}
}