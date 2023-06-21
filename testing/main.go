package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	if drives, err := Detect(); err == nil {
		fmt.Printf("%d USB Devices Found\n", len(drives))
		for _, d := range drives {
			fmt.Println(d)
		}
	} else {
		fmt.Println(err)
	}
}

// Detect returns a list of file paths pointing to the root folder of
// USB storage devices connected to the system.
func Detect() ([]string, error) {
	var drives []string
	driveMap := make(map[string]bool)
	dfPattern := regexp.MustCompile("^(\\/[^ ]+)[^%]+%[ ]+(.+)$")

	cmd := "df"
	out, err := exec.Command(cmd).Output()

	if err != nil {
		log.Printf("Error calling df: %s", err)
	}

	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		line := s.Text()
		if dfPattern.MatchString(line) {
			device := dfPattern.FindStringSubmatch(line)[1]
			rootPath := dfPattern.FindStringSubmatch(line)[2]

			if ok := isUSBStorage(device); ok {
				driveMap[rootPath] = true
			}
		}
	}

	for k := range driveMap {
		_, err := os.Open(k)
		if err == nil {
			drives = append(drives, k)
		}
	}

	return drives, nil
}

func isUSBStorage(device string) bool {
	deviceVerifier := "ID_USB_DRIVER=usb-storage"
	cmd := "udevadm"
	args := []string{"info", "-q", "property", "-n", device}
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		log.Printf("Error checking device %s: %s", device, err)
		return false
	}

	if strings.Contains(string(out), deviceVerifier) {
		return true
	}

	return false
}
