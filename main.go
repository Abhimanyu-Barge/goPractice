package main

import (
	"fmt"

	"github.com/Abhimanyu-Barge/goPractice/api"
)

func main() {
	// models.PracticeDemo()
	Port := 3000
	StartedPort, err := api.StartWebServer(Port)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("localhost:",StartedPort)

}
