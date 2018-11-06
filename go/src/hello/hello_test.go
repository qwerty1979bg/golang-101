package main

import (
    "fmt"
    "log"
    "os/exec"
    "strings"
    "os"
    "testing"
)

// Use the included GO testing funcionality
func TestHello(t *testing.T) {

fmt.Printf("Start\n")

// Could probably use a variable for the program name also
progname := os.ExpandEnv("${GOPATH}/bin/hello")

// Execute our compiled application
    out, err := exec.Command(progname).CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }

// Some debug info, just in case...
    debug, err := exec.Command("file", progname).CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
fmt.Printf("DEBUG:\n%s\n", debug)

// Turn the output into a string
out2 := string(out[:])

// Trim any trailing newline
out3 := strings.TrimRight(out2, "\n")

// Check if the output contains "hello"
if strings.Compare(out3, "hello") != 0 {
	fmt.Printf("The test failed.\nValue is '%s', while the expected was 'hello'\n", out3)
	os.Exit (1)
	}

fmt.Printf("The test suceeded.\nThe value '%s' matches the expected.\n", out3)

}
