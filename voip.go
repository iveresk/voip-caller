package main

import (
	"log"
	"net"
	"time"
)

func makeCall(target, port, phrase string, ch chan string) {
	payload := "INVITE sip:\"" + target + "\" SIP/2.0 ia: SIP/2.0/TCP To: sip:\"" + target + "\";tag=x From: sip:\"" + phrase + "\" Call-ID: \"27346824\" Cseq: 1 INVITE"
	targAddr := target + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", targAddr)
	if err != nil {
		ch <- "IP resolve was failed"
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		ch <- "The Dial was failed"
	}
	log.Println("Sending payload to Init the Call: ", payload)
	_, err = conn.Write([]byte(payload))

	if err != nil {
		ch <- "Payload delivery was failed"
	}

	log.Println("Call duration = Minute")
	time.Sleep(3 * time.Minute)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		ch <- "We are blocked. The response was failed"
	}
	conn.Close()
	ch <- "The call was made successfully"
}
