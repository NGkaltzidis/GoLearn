package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type DevicesMonitor struct {
	devicesPlugged   map[string]string
	activatedDevices map[string]string
	mutex            sync.Mutex
}

func NewDevicesMonitor() *DevicesMonitor {
	return &DevicesMonitor{
		devicesPlugged:   make(map[string]string),
		activatedDevices: make(map[string]string),
	}
}

func (dm *DevicesMonitor) getConnectedDrivesList() {
	cmd := exec.Command("bash", "-c", "lsblk | awk -v col1=1 -v col2=4 -v col3=6 '{print $col1, $col2, $col3}'")

	output, err := cmd.Output()
	if err != nil {
		log.Println("Error executing lsblk command:", err)
		return
	}
	stdOut := string(output)
	lines := strings.Split(strings.TrimSpace(stdOut), "\n")
	dm.mutex.Lock()
	dm.devicesPlugged = make(map[string]string)
	dm.mutex.Unlock()
	for _, line := range lines {
		if strings.Contains(line, "disk") && strings.Contains(line, "sd") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				driveName := fields[0]
				driveSize := fields[1]
				driveSerial := getDiskSerial(driveName)
				if driveSerial != "" {
					dm.mutex.Lock()
					if _, ok := dm.activatedDevices[driveName]; !ok {
						activated, err := dm.checkDriveActivation(driveName, driveSerial)
						if err != nil {
							log.Println("Error checking drive activation:", err)
						}
						if activated {
							dm.activatedDevices[driveName] = fmt.Sprintf("%s %s", driveSize, driveSerial)
							dm.mutex.Unlock()
							return
						}
					}
					dm.devicesPlugged[driveName] = fmt.Sprintf("%s %s", driveSize, driveSerial)
					dm.mutex.Unlock()
				}
			}
		}
	}
}

func (dm *DevicesMonitor) checkDriveActivation(driveName, driveSerial string) (bool, error) {
	cmd := exec.Command("bash", "-c", "lsblk | awk -v col1=1 -v col2=4 -v col3=6 '{print $col1, $col2, $col3}'")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	stdOut := string(output)
	lines := strings.Split(strings.TrimSpace(stdOut), "\n")
	for _, line := range lines {
		if strings.Contains(line, driveName) && strings.Contains(line, "part") {
			mountPoint := "/mnt/ssd_drive"
			err := os.MkdirAll(mountPoint, os.ModePerm)
			if err != nil {
				return false, err
			}
			err = exec.Command("mount", fmt.Sprintf("/dev/%s1", driveName), mountPoint).Run()
			if err != nil {
				return false, err
			}
			defer exec.Command("umount", mountPoint).Run()
			if _, err := os.Stat("/tmp/licence.txt"); err == nil {
				os.Remove("/tmp/licence.txt")
			}
			if _, err := os.Stat("/tmp/.licence"); err == nil {
				os.Remove("/tmp/.licence")
				defer os.Remove(fmt.Sprintf("%s/.licence", mountPoint))
			}
			if _, err := os.Stat(fmt.Sprintf("%s/.licence", mountPoint)); err == nil {
				err := copyFile(fmt.Sprintf("%s/.licence", mountPoint), "/tmp/.licence")
				if err != nil {
					return false, err
				}
				defer os.Remove("/tmp/.licence")
				result, err := exec.Command("sh", "-c", fmt.Sprintf("echo %s | gpg -o /tmp/licence.txt -d --batch --passphrase-fd 0 /tmp/.licence", os.Getenv("HARDCODED_PASS"))).CombinedOutput()
				if err != nil {
					log.Println("Error decrypting licence:", err)
					log.Println("Output:", string(result))
					return false, err
				}
				licence, err := readLicenseFile("/tmp/licence.txt")
				if err != nil {
					log.Println("Error reading licence file:", err)
					return false, err
				}
				if licence == driveSerial {
					fmt.Println("Licence OK")
					return true, nil
				} else {
					fmt.Println("Licence BAD")
					return false, nil
				}
			} else {
				return false, nil
			}
		}
	}
	return false, nil
}

func (dm *DevicesMonitor) runMain() {
	for {
		time.Sleep(time.Second)
		dm.getConnectedDrivesList()
	}
}

func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func readLicenseFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text(), nil
}

func clearPartitionTable(diskName string) bool {
	cmd := exec.Command("wipefs", "-a", fmt.Sprintf("/dev/%s", diskName))
	err := cmd.Run()
	if err != nil {
		log.Println("Error clearing partition table:", err)
		return false
	}
	return true
}

func createPartition(diskName string) bool {
	cmd := exec.Command("parted", "--script", fmt.Sprintf("/dev/%s", diskName), "mklabel", "gpt", "mkpart", "primary", "ext4", "1MB", "1000GB")
	err := cmd.Run()
	if err != nil {
		log.Println("Error creating partition:", err)
		return false
	}
	cmd = exec.Command("mkfs.ext4", fmt.Sprintf("/dev/%s1", diskName))
	err = cmd.Run()
	if err != nil {
		log.Println("Error formatting partition:", err)
		return false
	}
	return true
}

func putCryptedFile(driveName string) bool {
	driveSerial := getDiskSerial(driveName)
	if driveSerial != "" {
		mountPoint := "/mnt/ssd_drive"
		err := os.MkdirAll(mountPoint, os.ModePerm)
		if err != nil {
			log.Println("Error creating mount point:", err)
			return false
		}
		err = exec.Command("mount", fmt.Sprintf("/dev/%s1", driveName), mountPoint).Run()
		if err != nil {
			log.Println("Error mounting drive:", err)
			return false
		}
		defer exec.Command("umount", mountPoint).Run()
		licenceFilePath := fmt.Sprintf("%s/licence.txt", mountPoint)
		file, err := os.Create(licenceFilePath)
		if err != nil {
			log.Println("Error creating licence file:", err)
			return false
		}
		defer os.Remove(licenceFilePath)
		defer file.Close()
		rand.Seed(time.Now().UnixNano())
		licence := generateRandomString(10) + driveSerial + generateRandomString(10)
		_, err = file.WriteString(licence)
		if err != nil {
			log.Println("Error writing licence to file:", err)
			return false
		}
		if _, err := os.Stat(fmt.Sprintf("%s/.licence", mountPoint)); err == nil {
			err := os.Remove(fmt.Sprintf("%s/.licence", mountPoint))
			if err != nil {
				log.Println("Error removing existing .licence file:", err)
				return false
			}
		}
		err = exec.Command("sh", "-c", fmt.Sprintf("echo %s | gpg -o %s/.licence --batch -c --passphrase-fd 0 %s/licence.txt", os.Getenv("HARDCODED_PASS"), mountPoint, mountPoint)).Run()
		if err != nil {
			log.Println("Error encrypting licence file:", err)
			return false
		}
		if _, err := os.Stat(fmt.Sprintf("%s/licence.txt", mountPoint)); err == nil {
			err := os.Remove(fmt.Sprintf("%s/licence.txt", mountPoint))
			if err != nil {
				log.Println("Error removing plaintext licence file:", err)
				return false
			}
		}
		return true
	}
	return false
}

func getDiskSerial(diskName string) string {
	url := fmt.Sprintf("http://localhost:12000/system/drive/serial?name=%s", diskName)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error requesting drive serial:", err)
		return ""
	}
	defer resp.Body.Close()
	var result struct {
		Serial string `json:"serial"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error decoding drive serial:", err)
		return ""
	}
	return result.Serial
}

func registerDisk(driveName string) bool {
	driveSerial := getDiskSerial(driveName)
	if driveSerial != "" {
		authToken := getOAuthToken()
		headers := make(map[string]string)
		deviceURL := "https://init.monitoreal.com/api/devices/ssd/init_serial/"
		headers["Accept"] = "application/json"
		headers["Authorization"] = "Bearer " + authToken
		data := map[string]string{
			"drive_full_serial":  driveSerial,
			"drive_short_serial": strings.Split(driveSerial, "_")[1],
			"drive_size":         "931.5G",
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Error marshaling register disk data:", err)
			return false
		}
		req, err := http.NewRequest("POST", deviceURL, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Error creating register disk request:", err)
			return false
		}
		for key, value := range headers {
			req.Header.Set(key, value)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error registering disk:", err)
			return false
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Println("Disk registration request returned non-OK status code:", resp.StatusCode)
			return false
		}
		return true
	}
	return false
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func getOAuthToken() string {
	authTokenURL := "https://init.monitoreal.com/api/auth/login/"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	data := map[string]string{
		"email":    "admin@example.com",
		"password": "password123",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshaling auth token data:", err)
		return ""
	}
	resp, err := http.Post(authTokenURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error getting auth token:", err)
		return ""
	}
	defer resp.Body.Close()
	var result struct {
		Token string `json:"access"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error decoding auth token:", err)
		return ""
	}
	return result.Token
}

func main() {
	dm := NewDevicesMonitor()
	go dm.runMain()

	for {
		time.Sleep(time.Second)
		dm.mutex.Lock()
		for driveName, driveInfo := range dm.devicesPlugged {
			if driveInfo == "931.5G" {
				fmt.Println("Drive plugged:", driveName)
				if clearPartitionTable(driveName) {
					if createPartition(driveName) {
						if putCryptedFile(driveName) {
							if registerDisk(driveName) {
								dm.activatedDevices[driveName] = driveInfo
								delete(dm.devicesPlugged, driveName)
							} else {
								log.Println("Error registering disk:", driveName)
							}
						} else {
							log.Println("Error putting crypted file on drive:", driveName)
						}
					} else {
						log.Println("Error creating partition on drive:", driveName)
					}
				} else {
					log.Println("Error clearing partition table on drive:", driveName)
				}
			}
		}
		dm.mutex.Unlock()
	}
}
