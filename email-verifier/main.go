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
	fmt.Printf("domain, hasMX, hasSPF, SPFRecord, hasDMARC, DMARCRecord\n")
	for scanner.Scan() {
		domain, hasMX, hasSPF, SPFRecord, hasDMARC, DMARCRecord := checkDomain(scanner.Text())
		fmt.Printf("Domain: %v, has MX records: %v\n", domain, hasMX)
		fmt.Printf("\tHas SPF: %v, SPF Record: %v\n", hasSPF, SPFRecord)
		fmt.Printf("\tHas DMARC: %v, DMARC Record: %v\n", hasDMARC, DMARCRecord)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) (string, bool, bool, string, bool, string) {
	var hasMX, hasSPF, hasDMARC bool
	var SPFRecord, DMARCRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	hasMX = len(mxRecords) > 0

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPFRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			DMARCRecord = record
			break
		}
	}

	return domain, hasMX, hasSPF, SPFRecord, hasDMARC, DMARCRecord

}
