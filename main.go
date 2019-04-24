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
	// Read Flags
	colorize(color.FgGreen, "=> Reading Config")

	const (
		defaultPort       = "8084"
		defaultPortUsage  = "default TextBar port (8084)"
		defaultSleep      = 2000
		defaultSleepUsage = "default sleep between words milliseconds (2000ms)"
	)

	port := flag.String("port", defaultPort, defaultPortUsage)
	sleep := flag.Duration("sleep", defaultSleep, defaultSleepUsage)

	flag.Parse()

	url := "http://127.0.0.1:" + *port

	colorize(color.FgCyan, "=> Port: ", *port)
	colorize(color.FgCyan, "=> Sleep: ", *sleep)
	colorize(color.FgYellow, "=> Connecting to TB ", url)

	// conn, err := net.Dial("udp", "127.0.0.1:"+*port)
	// if err != nil {
	// 	fmt.Printf("Error connecting to TextBar (127.0.0.1:%s) %v", *port, err)
	// 	return
	// }
	// defer conn.Close()

	// colorize(color.FgGreen, "=> Connected to TB 127.0.0.1:", *port)
	// fmt.Fprintf(conn, "TBWords")

	// Read File
	lineScanner := bufio.NewScanner(os.Stdin)
	lineScanner.Split(bufio.ScanLines)

	for lineScanner.Scan() {
		line := lineScanner.Text()

		wordScanner := bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			word := wordScanner.Text()

			// Send to TB
			colorize(color.FgBlue, word)
			// fmt.Fprintf(conn, word)
			_, err := http.Post(url, "text/plain", strings.NewReader(word))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(*sleep * time.Millisecond)
		}
	}

	if err := lineScanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
