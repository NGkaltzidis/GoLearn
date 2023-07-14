package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net"
	"time"
)

type Envelope struct {
	XMLName    xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	ProbeMatch ProbeMatch
}

type ProbeMatch struct {
	XMLName           xml.Name `xml:"http://schemas.xmlsoap.org/ws/2005/04/discovery ProbeMatch"`
	EndpointReference EndpointReference
	Types             string
	Scopes            string
	XAddrs            string
	MetadataVersion   int
}

type EndpointReference struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing EndpointReference"`
	Address string
}

func main() {
	// Set the ONVIF discovery message
	msg := []byte(`<?xml version="1.0" encoding="UTF-8"?>
		<e:Envelope xmlns:e="http://www.w3.org/2003/05/soap-envelope"
			xmlns:w="http://schemas.xmlsoap.org/ws/2004/08/addressing"
			xmlns:d="http://schemas.xmlsoap.org/ws/2005/04/discovery"
			xmlns:dn="http://www.onvif.org/ver10/network/wsdl">
			<e:Header>
				<w:MessageID>uuid:84ede3de-7dec-11d0-c360-f01234567890</w:MessageID>
				<w:To>urn:schemas-xmlsoap-org:ws:2005:04:discovery</w:To>
				<w:Action>http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe</w:Action>
			</e:Header>
			<e:Body>
				<d:Probe>
					<d:Types>dn:NetworkVideoTransmitter</d:Types>
				</d:Probe>
			</e:Body>
		</e:Envelope>`)

	// Create a UDP listener for receiving responses
	listenAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:0")
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}

	conn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen on UDP: %v", err)
	}
	defer conn.Close()

	// Set the broadcast address and port for sending the discovery message
	broadcastAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:3702")
	if err != nil {
		log.Fatalf("Failed to resolve broadcast address: %v", err)
	}

	// Send the discovery message
	_, err = conn.WriteToUDP(msg, broadcastAddr)
	if err != nil {
		log.Fatalf("Failed to send discovery message: %v", err)
	}

	fmt.Println("Waiting for ONVIF devices to respond...")

	// Set a timeout for receiving responses
	timeout := 10 * time.Second
	deadline := time.Now().Add(timeout)
	conn.SetReadDeadline(deadline)

	// Buffer to hold the response data
	buf := make([]byte, 4096)

	// Receive and process the responses
	for {
		// Read the response data
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			netErr, ok := err.(net.Error)
			if ok && netErr.Timeout() {
				// Timeout occurred, exit the loop
				break
			}
			log.Fatalf("Error reading UDP response: %v", err)
		}

		// Process the response data
		response := string(buf[:n])
		fmt.Println("Received response:", response)

		// Parse the XML response
		var envelope Envelope
		err = xml.Unmarshal([]byte(response), &envelope)
		if err != nil {
			fmt.Println("Failed to parse XML response:", err)
			continue
		}

		// Extract the device's address (XAddrs)
		fmt.Println("Device address:", envelope.ProbeMatch.EndpointReference.Address)
	}

	fmt.Println("ONVIF device discovery completed.")
	fmt.Println("-----------------------------------------------")
}
