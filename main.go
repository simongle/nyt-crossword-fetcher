package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/net/publicsuffix"
)

type CrossWordFetchEvent struct {
	baseUrl         string `json:"base_url"`
	crosswordDate   string
	nytCookieString string
}

// 1. Log into NYT, in same browser visit the direct URL of a crossword PDF like https://www.nytimes.com/svc/crosswords/v2/puzzle/print/Jun0824.pdf
// 2. Open the network tab, refresh the page, copy request headers for PDF
// 3. Paste the string value of CookieString as the value of rawCookieString below

func CrosswordFetcher(event *CrossWordFetchEvent) []byte {

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("initiating HTTP client")

	c := &http.Client{
		Jar: jar,
	}

	urlObj, err := url.Parse(event.baseUrl + event.crosswordDate + ".pdf")

	if err != nil {
		log.Fatal("error constructing URL for fetch, check event parameters `crosswordDate` and `baseUrl`")
	}

	parsedCookies := []*http.Cookie{}

	for _, cookieString := range strings.Split(event.nytCookieString, "; ") {
		parsedCookie := strings.Split(cookieString, "=")

		cookie := &http.Cookie{
			Name:     parsedCookie[0],
			Value:    parsedCookie[1],
			Secure:   false,
			HttpOnly: false,
		}

		parsedCookies = append(parsedCookies, cookie)
	}

	c.Jar.SetCookies(urlObj, parsedCookies)

	var d string
	if event.crosswordDate != "" {
		d = event.crosswordDate
	} else {
		d = formatTime(time.Now())
	}
	fmt.Println("fetching puzzle from: ", d)

	// Fetch puzzle
	resp, err := c.Get(event.baseUrl + d + ".pdf")

	if err != nil {
		log.Fatal(err)
	}

	b, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	status := resp.Status
	log.Print(status)

	err = os.WriteFile("fetched/"+d+".pdf", b, 0666)

	if err != nil {
		log.Fatal(err)
	}

	// Return a slice of first 20 bytes of file to use for testing
	head := b[:20]
	return head
}

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	baseUrl, ok := os.LookupEnv("NYT_BASE_URL")

	if !ok {
		log.Fatal("missing NYT_BASE_URL")
	}

	nytCookieString, ok := os.LookupEnv("NYT_COOKIE_STRING")

	if !ok {
		log.Fatal("missing NYT_COOKIE_STRING")
	}

	crosswordDate, ok := os.LookupEnv("NYT_CROSSWORD_DATE")

	if !ok {
		fmt.Println("No date provided, fetching todays puzzle")
	}

	invokeData := CrossWordFetchEvent{
		baseUrl:         baseUrl,
		crosswordDate:   crosswordDate,
		nytCookieString: nytCookieString,
	}

	CrosswordFetcher(&invokeData)
}
