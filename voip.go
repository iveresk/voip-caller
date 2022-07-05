package main

import (
	"log"
	"net"
	"time"
)

func makeCall(target, port, phrase string, debug bool, ch chan string) {
	payload := "INVITE sip:\"100@" + target + "\" SIP/2.0 Via: SIP/2.0/UDP To: sip:\"100@" + target + "\";tag=x From: sip:\"" + phrase + "\" Call-ID: \"27346824\" Cseq: 1 INVITE"
	targAddr := target + ":" + port

	conn, err := net.Dial("udp", targAddr)
	if err != nil {
		ch <- "The Dial was failed"
		return
	}

	if conn == nil {
		ch <- "SIP request was failed. Port 5060 is filtered or closed"
		return
	}
	if debug {
		log.Println("Sending payload to Init the Call: ", payload)
	}
	_, err = conn.Write([]byte(payload))

	if err != nil {
		ch <- "Payload first attempt delivery was failed"
		return
	}
	if debug {
		log.Println("Call duration = Minute")
	}
	time.Sleep(time.Minute)

	reply := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(15 * time.Second))
	_, err = conn.Read(reply)
	if err != nil {
		ch <- "We are blocked. The response was failed"
		return
	}
	conn.Close()
	ch <- "The call was made successfully"
	return
}
