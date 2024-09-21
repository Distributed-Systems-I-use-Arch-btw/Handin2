package main

import (
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	seq := make([]byte, 1)
	ack := make([]byte, 1)
	seq[0] = 200

	conn.Read(ack) // Read seq

	// Adds 1 to the seq recived
	ack[0] += 1

	// Send the ack and seq
	conn.Write(ack) // Sends ack
	conn.Write(seq) // Send seq

	// Reads the ack + 1 and seq + 1
	seqRecived := make([]byte, 1)
	ackRecived := make([]byte, 1)
	
	conn.Read(ackRecived)
	conn.Read(seqRecived)

	// Test that they are correct
	if ackRecived[0] == (seq[0]+1) && seqRecived[0] == ack[0] {
		fmt.Println("Correct seq and ack recived")
	    fmt.Println("Connection established")
	} else {
		fmt.Println("Not correct")
	}

    data := time.Now().String()

    // Send data to client
    conn.Write([]byte(data))
    
    receivedData := make([]byte, 1024)
    
    // Read data ack
    conn.Read(receivedData)
    
    
    if int(receivedData[0]) == len(data) + 1 {
        fmt.Println("Data was recived correctly")
    } else {
        fmt.Println("Data was not recived correctly")
    }
}
