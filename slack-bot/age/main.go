package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println(event)
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5394475290484-5391945966659-fNlVK9bD0UBWsig5jjHKqHW2")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05BLE2UF6G-5394480490948-4c4d92c531f849f4e10be9071ff478d6f2d0e23be48d471dc41b6d3a142a9b7c")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "Calculate your age",
		Examples:    []string{"my yob is 1990"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Invalid year of birth")
				return
			}
			age := time.Now().Year() - yob
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)

		},
	})

	bot.Command("Ping!", &slacker.CommandDefinition{
		Description: "Ping!",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Pong!")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
