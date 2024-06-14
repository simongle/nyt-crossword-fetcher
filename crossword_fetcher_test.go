package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

func TestCrosswordFetcher(t *testing.T) {

	// Load environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	nytCookieString, ok := os.LookupEnv("NYT_COOKIE_STRING")

	if !ok {
		log.Fatal("missing NYT_COOKIE_STRING")
	}

	testInvokeData := CrossWordFetchEvent{
		baseUrl:         "https://www.nytimes.com/svc/crosswords/v2/puzzle/print/",
		crosswordDate:   "Jun0824",
		nytCookieString: nytCookieString,
	}

	var want []byte

	// Compare the first 20 bytes of the file array to confirm
	staticFile, _ := os.ReadFile("static/test.pdf")

	want = staticFile[:20]

	got := CrosswordFetcher(&testInvokeData)

	fmt.Printf("got %q want %q", got, want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}

}
