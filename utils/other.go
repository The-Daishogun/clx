package utils

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadStdIn() string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Println("Error getting stdin status:", err)
		return ""
	}

	if fi.Mode()&os.ModeCharDevice != 0 {
		return ""
	}
	reader := bufio.NewReader(os.Stdin)
	var input []byte
	buf := make([]byte, 1024)

	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			log.Println("Error reading input:", err)
			return ""
		}

		// Append read bytes to input
		input = append(input, buf[:n]...)

		// Break if end of file is reached
		if err == io.EOF {
			break
		}
	}

	return string(input)
}

func ReadDistroInfo() string {
	osReleaseFileBytes, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return ""
	}
	return string(osReleaseFileBytes)
}
