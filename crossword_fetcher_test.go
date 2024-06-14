package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestCrosswordFetcher(t *testing.T) {

	testInvokeEvent := CrossWordFetchEvent{
		baseUrl: "https://www.nytimes.com/svc/crosswords/v2/puzzle/print/",
		date:    "Jun1024",
	}

	var want []byte

	// Compare the first 20 bytes of the file array to confirm
	staticFile, _ := os.ReadFile("static/test.pdf")

	want = staticFile[:20]

	got := CrosswordFetcher(&testInvokeEvent)

	fmt.Printf("got %q want %q", got, want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}

}
