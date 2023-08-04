package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Read(inputPath string) FriendsData {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var friendsData FriendsData
	err = json.Unmarshal(data, &friendsData)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	fmt.Println("Friends:")
	for _, friend := range friendsData.Friends {
		fmt.Printf("Name: %s\n", friend.Name)
		fmt.Printf("Description: %s\n", friend.Description)
		fmt.Printf("URL: %s\n", friend.URL)
		fmt.Println("------------------")
	}

	return friendsData
}
