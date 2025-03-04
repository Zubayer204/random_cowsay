package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

// Check if a command is available in PATH
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// Run a command and return its output
func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// Run fortune piped to cowsay
func runFortuneWithCowsay(cow string) (string, error) {
	fortune, err := runCommand("fortune")
	if err != nil {
		return "", fmt.Errorf("error running fortune: %v", err)
	}

	cowsay, err := runCommand("cowsay", "-f", cow, fortune)
	if err != nil {
		return "", fmt.Errorf("error running cowsay: %v", err)
	}

	return cowsay, nil
}

// Run cowsay with a custom message
func runCowsayWithMessage(message string) (string, error) {
	return runCommand("cowsay", message)
}

// List available cowfiles
func listCowfiles() ([]string, error) {
	// First, let's find where cowfiles are stored
	output, err := runCommand("cowsay", "-l", "|", "tail", "-n", "+2")
	if err != nil {
		return nil, fmt.Errorf("error finding cowfiles: %v", err)
	}

	files := strings.Split(strings.TrimSpace(string(output)), "\n")

	// Extract just the cowfile names without path and extension
	var cowfiles []string
	for _, file := range files {
		if file == "" {
			continue
		}
		parts := strings.Split(file, "/")
		filename := parts[len(parts)-1]
		cowfiles = append(cowfiles, strings.TrimSuffix(filename, ".cow"))
	}

	return cowfiles, nil
}

// Get a random cowfile from the available list
func getRandomCowfile(cowfiles []string) string {
	if len(cowfiles) == 0 {
		return "default"
	}
	// Generate a random index
	randomIndex := rand.Intn(len(cowfiles))

	return cowfiles[randomIndex]
}

func main() {
	// Check if cowsay is installed
	if !commandExists("cowsay") {
		fmt.Println("cowsay is not installed. Please install it first.")
		fmt.Println("On Debian/Ubuntu: sudo apt-get install cowsay")
		fmt.Println("On macOS with Homebrew: brew install cowsay")
		os.Exit(1)
	}

	// Check if fortune is installed
	fortuneExists := commandExists("fortune")
	if !fortuneExists {
		fmt.Println("fortune is not installed. Please install it first.")
		fmt.Println("On Debian/Ubuntu: sudo apt-get install fortune")
		fmt.Println("On macOS with Homebrew: brew install fortune")
		os.Exit(1)
	}

	// List available cowfiles
	cowfiles, err := listCowfiles()
	if err != nil {
		fmt.Printf("Error listing cowfiles: %v\n", err)
	} else if len(cowfiles) > 0 {
		// Use a random cowfile for fun
		if fortuneExists && len(cowfiles) > 0 {
			randomCow := getRandomCowfile(cowfiles)
			output, err := runFortuneWithCowsay(randomCow)
			if err == nil {
				fmt.Println(output)
			}
		}
	}
}
