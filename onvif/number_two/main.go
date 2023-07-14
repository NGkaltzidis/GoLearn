package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// Structs for ONVIF request and response
type Envelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Body    Body
}

type Body struct {
	XMLName                 xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
	GetDeviceInformationRes GetDeviceInformationResponse
	GetCapabilitiesRes      GetCapabilitiesResponse
}

type GetDeviceInformationResponse struct {
	XMLName         xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetDeviceInformationResponse"`
	Manufacturer    string
	Model           string
	SerialNumber    string
	FirmwareVersion string
	HardwareId      string
}

type GetCapabilitiesResponse struct {
	XMLName      xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetCapabilitiesResponse"`
	Capabilities Capabilities
}

type Capabilities struct {
	XMLName        xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl Capabilities"`
	Analytics      bool     // Whether the device supports analytics
	Events         bool     // Whether the device supports events
	Imaging        bool     // Whether the device supports imaging
	IO             bool     // Whether the device supports I/O
	PTZ            bool     // Whether the device supports PTZ
	Recording      bool     // Whether the device supports recording
	Replay         bool     // Whether the device supports replay
	Security       bool     // Whether the device supports security features
	Streaming      bool     // Whether the device supports streaming
	DeviceIO       bool     // Whether the device supports device I/O
	VideoAnalytics bool     // Whether the device supports video analytics
	// Add other necessary fields based on your requirements
}

func main() {
	// Endpoint URL for GetDeviceInformation and GetCapabilities requests
	endpointURL := "http://192.168.50.92:80/onvif/device_service"

	// Create the SOAP request for GetDeviceInformation
	getDeviceInformationEnvelope := Envelope{
		Body: Body{
			GetDeviceInformationRes: GetDeviceInformationResponse{},
		},
	}
	getDeviceInformationRequestBody, err := xml.Marshal(getDeviceInformationEnvelope)
	if err != nil {
		fmt.Println("Error marshaling GetDeviceInformation SOAP request:", err)
		return
	}

	// Send the GetDeviceInformation SOAP request to the ONVIF endpoint
	getDeviceInformationResp, err := http.Post(endpointURL, "application/soap+xml", bytes.NewBuffer(getDeviceInformationRequestBody))
	if err != nil {
		fmt.Println("Error sending GetDeviceInformation SOAP request:", err)
		return
	}
	defer getDeviceInformationResp.Body.Close()

	// Read the GetDeviceInformation response body
	getDeviceInformationResponseBody, err := io.ReadAll(getDeviceInformationResp.Body)
	if err != nil {
		fmt.Println("Error reading GetDeviceInformation SOAP response:", err)
		return
	}

	// Print the raw GetDeviceInformation SOAP response
	fmt.Println("Raw GetDeviceInformation SOAP Response:")
	fmt.Println(string(getDeviceInformationResponseBody))

	// Parse the GetDeviceInformation SOAP response
	var getDeviceInformationResponse Envelope
	err = xml.Unmarshal(getDeviceInformationResponseBody, &getDeviceInformationResponse)
	if err != nil {
		fmt.Println("Error unmarshaling GetDeviceInformation SOAP response:", err)
		return
	}

	// Extract the device information from the GetDeviceInformation response
	deviceInfo := GetDeviceInformationResponse{}
	if getDeviceInformationResponse.Body.GetDeviceInformationRes.XMLName.Local == "GetDeviceInformationResponse" {
		deviceInfo = getDeviceInformationResponse.Body.GetDeviceInformationRes
	}

	// Print the device information
	fmt.Println("Manufacturer:", deviceInfo.Manufacturer)
	fmt.Println("Model:", deviceInfo.Model)
	fmt.Println("Serial Number:", deviceInfo.SerialNumber)
	fmt.Println("Firmware Version:", deviceInfo.FirmwareVersion)
	fmt.Println("Hardware ID:", deviceInfo.HardwareId)

	// Create the SOAP request for GetCapabilities
	getCapabilitiesEnvelope := Envelope{
		Body: Body{
			GetCapabilitiesRes: GetCapabilitiesResponse{},
		},
	}
	getCapabilitiesRequestBody, err := xml.Marshal(getCapabilitiesEnvelope)
	if err != nil {
		fmt.Println("Error marshaling GetCapabilities SOAP request:", err)
		return
	}

	// Send the GetCapabilities SOAP request to the ONVIF endpoint
	getCapabilitiesResp, err := http.Post(endpointURL, "application/soap+xml", bytes.NewBuffer(getCapabilitiesRequestBody))
	if err != nil {
		fmt.Println("Error sending GetCapabilities SOAP request:", err)
		return
	}
	defer getCapabilitiesResp.Body.Close()

	// Read the GetCapabilities response body
	getCapabilitiesResponseBody, err := io.ReadAll(getCapabilitiesResp.Body)
	if err != nil {
		fmt.Println("Error reading GetCapabilities SOAP response:", err)
		return
	}

	// Print the raw GetCapabilities SOAP response
	fmt.Println("Raw GetCapabilities SOAP Response:")
	fmt.Println(string(getCapabilitiesResponseBody))

	// Parse the GetCapabilities SOAP response
	var getCapabilitiesResponse Envelope
	err = xml.Unmarshal(getCapabilitiesResponseBody, &getCapabilitiesResponse)
	if err != nil {
		fmt.Println("Error unmarshaling GetCapabilities SOAP response:", err)
		return
	}

	// Extract the device capabilities from the GetCapabilities response
	capabilities := GetCapabilitiesResponse{}
	if getCapabilitiesResponse.Body.GetCapabilitiesRes.XMLName.Local == "GetCapabilitiesResponse" {
		capabilities = getCapabilitiesResponse.Body.GetCapabilitiesRes
	}
	fmt.Println("Analytics : ", capabilities.Capabilities.Analytics)
	fmt.Println("Analytics : ", capabilities.Capabilities.Streaming)
	// Process and use the device capabilities as needed
	// ...
}
