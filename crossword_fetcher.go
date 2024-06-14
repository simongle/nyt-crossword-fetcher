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

	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/net/publicsuffix"
)

type CrossWordFetchEvent struct {
	baseUrl string `json:"base_url"`
	date    string
}

// 1. Log into NYT, in same browser visit the direct URL of a crossword PDF like https://www.nytimes.com/svc/crosswords/v2/puzzle/print/Jun0824.pdf
// 2. Open the network tab, refresh the page, copy request headers for PDF
// 3. Paste the string value of CookieString as the value of rawCookieString below

const rawCookieString = ""

func CrosswordFetcher(event *CrossWordFetchEvent) []byte {

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("initiating client")

	c := &http.Client{
		Jar: jar,
	}

	urlObj, err := url.Parse(event.baseUrl + event.date + ".pdf")

	if err != nil {
		log.Fatal("error constructing URL for fetch, check event parameters `date` and `baseUrl`")
	}

	parsedCookies := []*http.Cookie{}

	for _, cookieString := range strings.Split(rawCookieString, "; ") {
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

	resp, err := c.Get(event.baseUrl + event.date + ".pdf")

	if err != nil {
		log.Fatal(err)
	}

	b, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	status := resp.Status
	log.Print(status)

	err = os.WriteFile("fetched/"+event.date+".pdf", b, 0666)
	if err != nil {
		log.Fatal(err)
	}

	head := b[:20]
	return head
}

func main() {
	lambda.Start(CrosswordFetcher)
}
