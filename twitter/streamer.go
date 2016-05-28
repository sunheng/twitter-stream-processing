package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"os"
)

var (
	consumerKey string = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret string = os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken string = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret string = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

func init() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
}

func NewAPI() *anaconda.TwitterApi {
	return anaconda.NewTwitterApi(accessToken, accessTokenSecret)
}