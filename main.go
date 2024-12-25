package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the domain name (or 'exit' to quit): ")

	for scanner.Scan() {
		domain := scanner.Text()
		if strings.ToLower(domain) == "exit" {
			break
		}
		checkDomain(domain)
		fmt.Println("\nEnter another domain name (or 'exit' to quit): ")
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading input: %v", err)
	}
}

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Check MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Printf("Error checking MX records: %v\n", err)
	} else {
		hasMx = len(mxRecords) > 0
	}

	// Check SPF records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("Error checking SPF records: %v\n", err)
	} else {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	}

	// Check DMARC records
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		fmt.Printf("Error checking DMARC records: %v\n", err)
	} else {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	}

	// Print results
	fmt.Println("\nResults for domain:", domain)
	fmt.Println("-------------------------------------------------")
	fmt.Println("Has MX record:", hasMx)
	fmt.Println("Has SPF record:", hasSPF)
	if hasSPF {
		fmt.Println("SPF record:", spfRecord)
	}
	fmt.Println("Has DMARC record:", hasDMARC)
	if hasDMARC {
		fmt.Println("DMARC record:", dmarcRecord)
	}
	fmt.Println("-------------------------------------------------")
}
