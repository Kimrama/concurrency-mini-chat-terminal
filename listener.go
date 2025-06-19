package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Kimrama/concurrency-mini-chat-terminal/models"
)

func listenToRoom(roomname string) {
	r := models.Manager.GetRoom(roomname)
	if r == nil {
		fmt.Println("Invalid!! room does not exist")
		return
	}
	fmt.Println("Listen to room", roomname)
	for msg := range r.Message {
		fmt.Println(msg)
	}

}

func interactive() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter room to listen (or `exit`): ")
		scanner.Scan()
		roomname := scanner.Text()
		if roomname == "exit" {
			break
		}
		go listenToRoom(roomname)
	}
}
