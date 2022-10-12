package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// look up mx
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error in MX lookup: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// look up spf records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error in SPF lookup: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// look up dmarc records
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error in DMARC lookup: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	// print out results
	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
func main() {

	// create scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	// listens for text in stdin  (i.e. advance to next token in stdin)
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	// log scanner error
	if err := scanner.Err(); err != nil {
		log.Printf("Error - could not read from stdin: %v\n", err)
	}
}
