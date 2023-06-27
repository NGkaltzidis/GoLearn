package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	DEVICES_API_USER     = "pavels"
	DEVICES_API_PASSWORD = "JPxr3tP6"
	HARDCODED_PASS       = "mT2baurZUrAybkva"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Drive name is required")
	}
	driveName := args[1]
	fmt.Printf("Drive name is %s\n", driveName)

	err := os.RemoveAll("/mnt/ssd_drive")
	if err != nil {
		log.Fatal("Failed to remove mount point directory: ", err)
	}

	if !clearPartitionTable(driveName) {
		log.Fatal("Cannot clear partition table")
	}

	if !createPartition(driveName) {
		log.Fatal("Cannot create partition")
	}

	if !registerDisk(driveName) {
		log.Fatal("Cannot register drive")
	}

	if putCryptedFile(driveName) {
		fmt.Println("Drive is initiated")
	} else {
		log.Fatal("Cannot put serial number on the drive")
	}
}

func clearPartitionTable(driveName string) bool {
	cmd := exec.Command("wipefs", "-a", fmt.Sprintf("/dev/%s", driveName))
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func createPartition(driveName string) bool {
	cmd := exec.Command("parted", "--script", fmt.Sprintf("/dev/%s", driveName),
		"mklabel", "gpt",
		"mkpart", "primary", "ext4", "1MB", "1000GB")
	err := cmd.Run()
	if err != nil {
		return false
	}

	cmd = exec.Command("mkfs.ext4", fmt.Sprintf("/dev/%s1", driveName))
	err = cmd.Run()
	if err != nil {
		return false
	}

	return true
}

func registerDisk(driveName string) bool {
	driveSerial := getDiskSerial(driveName)
	if driveSerial == "" {
		return false
	}

	authToken := getOAuthToken()

	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + authToken,
	}

	data := map[string]string{
		"drive_full_serial":  driveSerial,
		"drive_short_serial": strings.Split(driveSerial, "_")[1],
		"drive_size":         "931.5G",
	}

	resp, err := requests.Post("https://init.monitoreal.com/api/devices/ssd/init_serial/", data, headers)
	if err != nil || resp.StatusCode != 201 {
		return false
	}

	return true
}

func getOAuthToken() string {
	apiUser := DEVICES_API_USER
	apiPassword := DEVICES_API_PASSWORD
	tokenAPIURL := "https://init.monitoreal.com/api/o/token/"
	data := map[string]string{
		"username":      apiUser,
		"password":      apiPassword,
		"grant_type":    "password",
		"client_id":     "z3abhire2D0Qn2bKGID8dGPgtUqB2vhUzyrQGegN",
		"client_secret": "qnllGRFvTHS7oFRk22lOtoXUY63CZzEI37hT9pU6y9IztRuThxW9LWRmYxosvw2i9dcLWIn2YJc2eTNSK54ym46nDNetLLcuqCutvzU1XgVrzD5ZdAgKmHvMBMTYVp2B",
	}

	resp, err := requests.Post(tokenAPIURL, data, nil)
	if err != nil {
		log.Fatal("Failed to fetch OAuth token")
	}

	token := resp.Json()["access_token"].(string)
	return token
}

func getDiskSerial(driveName string) string {
	cmd := exec.Command("udevadm", "info", "--query=all", fmt.Sprintf("--name=/dev/%s", driveName))
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Failed to get disk serial")
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "ID_SERIAL=") {
			serial := strings.Split(line, "ID_SERIAL=")[1]
			return strings.Trim(serial, "\n")
		}
	}

	return ""
}

func putCryptedFile(driveName string) bool {
	driveSerial := getDiskSerial(driveName)
	if driveSerial == "" {
		return false
	}

	mountPoint := "/mnt/ssd_drive"
	err := os.MkdirAll(mountPoint, os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create mount point directory: ", err)
	}

	cmd := exec.Command("mount", fmt.Sprintf("/dev/%s1", driveName), mountPoint)
	err = cmd.Run()
	if err != nil {
		return false
	}

	filePath := filepath.Join(mountPoint, "licence.txt")
	content := generateRandomString(10) + driveSerial + generateRandomString(10)
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return false
	}

	licenceFile := filepath.Join(mountPoint, ".licence")
	if _, err := os.Stat(licenceFile); err == nil {
		os.Remove(licenceFile)
	}

	cmd = exec.Command("gpg", "-o", licenceFile, "--batch", "-c", "--passphrase-fd", "0", filePath)
	cmd.Stdin = strings.NewReader(HARDCODED_PASS)
	err = cmd.Run()
	if err != nil {
		return false
	}

	err = os.Remove(filePath)
	if err != nil {
		return false
	}

	return true
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
