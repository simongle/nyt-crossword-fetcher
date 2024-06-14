# NYT Crossword Fetcher
This will fetch the newspaper version of the crossword when provided with a datestring like `jun0624`. Its handy if you share an account with a partner who loves doing them in the app in the middle of the night before you've had a chance to log in.

## Setup
1. Log into NYT, in same browser visit the direct URL of a crossword PDF like https://www.nytimes.com/svc/crosswords/v2/puzzle/print/Jun0824.pdf
2. Open the network tab, refresh the page, copy request headers for PDF
3. Create .env file and add the following values:

```sh
NYT_COOKIE_STRING=""
NYT_BASE_URL=https://www.nytimes.com/svc/crosswords/v2/puzzle/print/
NYT_CROSSWORD_DATE=Jun0624
```

```bash
go build 

# testing
go test

# TODO add event invoke data with fixture file
go run .
```

## TODOs
* Create github repo
* Deploy lambda (maybe)
* Add integration to write to an s3 bucket or other cloud storage, currently stores them in a local directory. 
* Refactor to automatically construct date based on time.Now()
* Update test to use mocked HTTP request