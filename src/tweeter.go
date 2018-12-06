package main

import (
	"github.com/abiosoft/ishell"
	"github.com/danielacarrero/Twitter/src/service"
	"github.com/danielacarrero/Twitter/src/domain"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")
			user := "user"
			text := c.ReadLine()

			service.PublishTweet(domain.NewTweet(user, text))

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetLastTweet()

			c.Println(tweet.Text)

			return
		},
	})

	shell.Run()

}
