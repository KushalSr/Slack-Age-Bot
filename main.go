package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5683131141239-5683160671895-hAHAJ9smo9QXJODTNtKia1Bq")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05M7BCFC1W-5694761051925-586470eabc07b39e30134386e154bb9b626fdd83b143382b5e5721f0a32968f5")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My Year of Birth is <year>", &slacker.CommandDefinition{
		Description: "Year of Birth Calculator",
		//Example: "My Year of Birth is 2001",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {

			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("Error")
			}
			age := 2023 - yob
			r := fmt.Sprintf("Age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
