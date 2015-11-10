package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
)

func main() {

	const (
		CONSUMER_KEY = ""
		CONSUMER_SECRET = ""
		ACCESS_TOKEN = ""
		ACCESS_TOKEN_SECRET = ""	
	)

	args := os.Args

	action := args[1]

	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)

	api := anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

	switch action {
		case "tweet":

			if len(args[2:]) > 0 {
				api.PostTweet(args[2], nil)
				fmt.Println("Tweeted")
			} else {
				fmt.Println("No tweet defined :(")
			}
	
	}
}
