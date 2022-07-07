package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"os"

	"github.com/joho/godotenv"
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

	// load .env file from path
	err_env := godotenv.Load(".env")
	if err_env != nil {
		log.Fatalf("Error loading .env file")
	}

	// creating new slacker bot and getting env variables SLACK_BOT_TOKEN and SLACK_APP_TOKEN
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// printing all commands that take place to console (optional)
	go printCommandEvents(bot.CommandEvents())

	// when bot sees this command, it will handle as follows
	bot.Command("my dob is <BirthDate>", &slacker.CommandDefinition{
		Description: "age calculator",
		Example:     "my dob is 2000-07-11",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {

			// store inputted birthdate
			birthdate := request.Param("BirthDate")

			// parse and format birthdate
			b_date, err := time.Parse("2006-01-02", birthdate)
			if err != nil {
				println("error parsing birthdate")
			}
			b_year, b_month, b_day := b_date.Date()
			b_date = time.Date(b_year, b_month, b_day, 0, 0, 0, 0, time.UTC)

			// store and format current date
			c_date := time.Now()
			c_year, c_month, c_day := c_date.Date()
			c_date = time.Date(c_year, c_month, c_day, 0, 0, 0, 0, time.UTC)

			// returns true if current date is present before birthdate
			if c_date.Before(b_date) {
				return
			}

			// calculate age
			age := c_year - b_year

			// calculate anniversary
			anniversary := b_date.AddDate(age, 0, 0)
			if anniversary.After(c_date) {
				age--
			}

			// output age
			r := fmt.Sprintf("your age is %d", age)
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
