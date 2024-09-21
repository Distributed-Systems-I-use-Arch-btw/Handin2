package main

import (
	"fmt"
	"log"
	"net"
)

func client() {
	dial, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	//Makes the byte to send
	seqSend := make([]byte, 1)
	seqSend[0] = 100

	// Send seq
	dial.Write(seqSend)

	//Make recivers
	ackReceived := make([]byte, 1)
	seqRecived := make([]byte, 1)

	//Reads recives values and prints them
	dial.Read(seqRecived)  // Reads seq + 1
	dial.Read(ackReceived) // Reads ack

	//Checks that the correct value is recived
	if seqRecived[0] == (seqSend[0] + 1) {
		fmt.Println("Correct seq recived")
	} else {
		fmt.Println("Not correct")
	}
	//Adds 1 to the ack
	ackReceived[0] += 1
	//Sends back the ack + 1 and recived to establish the connection
	dial.Write(seqRecived)
	dial.Write(ackReceived)
}
