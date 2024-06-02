package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	// runs all the tests for all the packages
	cmd := exec.Command("go", "test", "-v", "./tests")
	cmd.Stdin = strings.NewReader("Some Input")
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	//  outputs the response from the cmd as a string
	if err != nil {
		log.Fatal("Error while running the tests : ", out.String())
	}
	fmt.Println("Triggered the test scripts for golang packages ", out.String())
}
