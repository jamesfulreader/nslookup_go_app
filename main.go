package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Performing NS lookup...")

	inputUrls, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputUrls.Close()

	outputUrls, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputUrls.Close()

	writer := bufio.NewWriter(outputUrls)

	scanner := bufio.NewScanner(inputUrls)
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}

		isp, err := net.LookupCNAME(url)
		if err != nil {
			writer.WriteString(fmt.Sprintf("Error resolving %s: %v\n", url, err))
			continue
		}

		writer.WriteString(fmt.Sprintf("%s: %v\n", url, isp))

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		writer.Flush()
	}
	fmt.Println("NS lookup completed. Results saved to output.txt.")
}
