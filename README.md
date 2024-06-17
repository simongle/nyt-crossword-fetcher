# NYT Crossword Fetcher
This is a simple Go app that will fetch the newspaper version of the times crossword when provided with a datestring like `jun0624`. 

It will automatically fetch today's puzzle (the day the script is run). If you would like to get a specific date, set a line in your `.env` like:
`NYT_CROSSWORD_DATE=Jun1524`

## Setup
1. Log into NYT, in same browser visit the direct URL of a crossword PDF like https://www.nytimes.com/svc/crosswords/v2/puzzle/print/Jun0824.pdf
2. Open the network tab, refresh the page, copy request headers for PDF
3. Create .env file and add the following values:

```sh
NYT_COOKIE_STRING=""
NYT_BASE_URL=https://www.nytimes.com/svc/crosswords/v2/puzzle/print/
NYT_CROSSWORD_DATE=Jun0624
```
4. Update lines 85 and 100 to the absolute path of working directory on your machine
5. *optional* Add a [crontab](https://crontab.guru/) entry on your machine with `crontab -e` to run this automatically. This for instance fetches the Wed-Sat puzzles at 1am

```bash
0 1 * * 3,4,5,6 ~/path/to/crossword_fetcher 
```

## Building 
```bash
go build 

# testing
go test

# Run a single test
go test -run TestFormatTime

go run .
```


## TODOs
* Deploy lambda (maybe)
* Add integration to write to an s3 bucket or other cloud storage, currently stores them in a local directory and copies to icloud drive on device. 
* Update test to use mocked HTTP request instead of actually requesting from live domain.