package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	colorize(color.FgGreen, "=> TBWords")

	const (
		defaultPort       = "8084"
		defaultPortUsage  = "default TextBar port (8084)"
		defaultSleep      = 2 * time.Second
		defaultSleepUsage = "default sleep between words (2s)"
	)

	port := flag.String("port", defaultPort, defaultPortUsage)
	sleep := flag.Duration("sleep", defaultSleep, defaultSleepUsage)

	flag.Parse()

	url := "http://127.0.0.1:" + *port

	colorize(color.FgCyan, "=> Port: ", *port)
	colorize(color.FgCyan, "=> Sleep: ", *sleep)
	colorize(color.FgYellow, "=> Connecting to TB ", url)

	lineScanner := bufio.NewScanner(os.Stdin)
	lineScanner.Split(bufio.ScanLines)

	for lineScanner.Scan() {
		line := lineScanner.Text()

		wordScanner := bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			word := wordScanner.Text()

			colorize(color.FgBlue, word)
			_, err := http.Post(url, "text/plain", strings.NewReader(word))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(*sleep)
		}
	}

	if err := lineScanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
