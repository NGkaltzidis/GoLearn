package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func randomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func main() {
	// Generate random letters
	randomLetters := randomString(10)
	randomLetters2 := randomString(10)

	// Create the content string
	content := randomLetters + "Samsung_SSD_870_EVO_1TB_S6PUNF0R600209F" + randomLetters2

	// Write the content to the file
	err := ioutil.WriteFile(".licence", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(".licence file created successfully.")
	fmt.Println(content)
}
