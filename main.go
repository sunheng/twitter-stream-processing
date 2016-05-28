package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"github.com/sunheng/twitter-stream-processing/twitter"
)

func main() {

	twitterAPI := twitter.NewAPI()
	defer twitterAPI.Close()
	v := url.Values{}
	v.Set("track", "#twitter")
	s := twitterAPI.PublicStreamFilter(v)
	defer s.Stop()

	for t := range s.C {
		switch v := t.(type) {
		case anaconda.Tweet:
			fmt.Printf("%-15s: %s\n Created At: %s", v.User.ScreenName, v.Text, v.CreatedAt)
			fmt.Println("=======")
		case anaconda.EventTweet:
			switch v.Event.Event {
			case "favorite":
				sn := v.Source.ScreenName
				tw := v.TargetObject.Text
				fmt.Printf("★ Favorited by %-15s: %s\n", sn, tw)
			case "unfavorite":
				sn := v.Source.ScreenName
				tw := v.TargetObject.Text
				fmt.Printf("★ UnFavorited by %-15s: %s\n", sn, tw)
			}
		}
	}
}