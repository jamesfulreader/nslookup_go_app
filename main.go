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
	// Open the input file containing URLs
	inputUrls, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputUrls.Close()
	// Create the output file to save results
	outputUrls, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputUrls.Close()

	// Create counter for
	c1 := 0
	c2 := 0

	// Create a buffered writer for the output file
	writer := bufio.NewWriter(outputUrls)
	// Read the input file line by line
	scanner := bufio.NewScanner(inputUrls)

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}
		// Perform NS lookup for each URL
		isp, err := net.LookupCNAME(url)
		if err != nil {
			writer.WriteString(fmt.Sprintf("Error resolving %s: %v\n", url, err))
			continue
		}

		// add isp to count and display in output
		if isp == "string of isp to look at" {
			c1++
			writer.WriteString(fmt.Sprintf("%s: %v\n", url, isp))
		} else {
			c2++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		writer.Flush()
	}
	writer.WriteString(fmt.Sprintf("c1 count: %v\nc2 count: %v\n", c1, c2))
	writer.Flush()

	fmt.Println("NS lookup completed. Results saved to output.txt.")
}
