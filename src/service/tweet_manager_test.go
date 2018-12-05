package service_test


import (
  "testing"
  "github.com/danielacarrero/Twitter/src/service"
  "github.com/danielacarrero/Twitter/src/domain"
)

//Se debe empezar con Test en la función para indicar que es test
// Debe recibir como parametro *testing.T, importando del package testing el tipo T, este tipo es el definido para test
// t es un puntero al tipo T
func TestPublishedTweetIsSaved(t *testing.T){
  //Initialization
  var tweet *domain.Tweet
  user := "danielacarrero"
  text := "This is my first tweet"
  tweet = domain.NewTweet(user, text)

  //Operation
  service.PublishTweet(tweet)

  //Validation

  /*if service.Tweet != tweet {
    t.Error("Expected tweet is", tweet)
  }*/

  publishedTweet := service.GetTweet()
  if publishedTweet.User != user || publishedTweet.Text != text {
    t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
  }

  if publishedTweet.Date == nil {
    t.Error("Expected date can't be nil")
  }
}

func TestTweetWithoutUserIsNotPublished(t *testing.T){
  //Initialization
  var tweet *domain.Tweet

  var user string
  text := "This is my first tweet"

  tweet = domain.NewTweet(user, text)

  //Operation
  var err error
  err = service.PublishTweet(tweet)

  //Validation
  if err != nil && err.Error() != "user is required" {
    t.Error("Expected error is user is required")
  }
}

func TestTweetWithoutTextIsNotPublished(t *testing.T){
  //Initialization
  var tweet *domain.Tweet

  var text string
  user := "danielacarrero"

  tweet = domain.NewTweet(user, text)

  //Operation
  var err error
  err = service.PublishTweet(tweet)

  //Validation
  if err != nil && err.Error() != "text is required" {
    t.Error("Expected error is text is required")
  }
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T){
  //Initialization
  var tweet *domain.Tweet

  text := `asdasdasdslajflkdkkkkkkaslkdjalkdjlaskjdlkajdl
            kjsdljdlkajdlkasjlksajflksjflksajflkjlkajsdlk
            jlkajsdlkjaflkjlkjdlkajdlajdlajdlakjdlkajdlak
            jflkasjdlkjsdlkjsadlkafjlakdjlkasjflkasjdljflkdljlaa`
  user := "danielacarrero"

  tweet = domain.NewTweet(user, text)

  //Operation
  var err error
  err = service.PublishTweet(tweet)

  //Validation
  if err != nil && err.Error() != "tweet exceeding 140 characters" {
    t.Error("Expected error is tweet exceeding 140 characters")
  }
}
