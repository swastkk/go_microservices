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
	fmt.Print("Enter a domain which you want to verify! \n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error: Could not read from input- %v", err)
		os.Exit(1)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDmarc bool
	var sprRecord, dmarcRecord string
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	textRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	for _, record := range textRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			sprRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDmarc = true
			dmarcRecord = record
			break
		}
	}
	fmt.Print("\n")
	fmt.Printf("Domain, hasMX, hasSPF, sprRecord, hasDmarc, dmarcRecord\n")
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, sprRecord, hasDmarc, dmarcRecord)
	fmt.Print("\n")
	fmt.Printf("So Email verifier tool has done its work!\n")
	os.Exit(1)
}
