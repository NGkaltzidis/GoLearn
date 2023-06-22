package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	device := "/dev/disk3"
	// name = command, args = options and then can be followed by whatever
	cmd := exec.Command("diskutil", "info", device)

	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "could not find disk") {
			fmt.Println(device, "is not connected.")
		} else {
			if err.Error() == "exit status 1" {
				fmt.Println(device, "is not connected.")
			}
			fmt.Println("Error executing command:", err.Error())
		}
		return
	}

	fmt.Println(device, "is connected.")
}
