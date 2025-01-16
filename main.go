package main

import "fmt"



func main() {
	domain := DomainUrl{"rohithbn27@gmail.com"}
	isValid := domain.isValidDomain()
	fmt.Printf("Domain: %s\n", isValid.DomainName)
	fmt.Printf("Has MX: %v\n", isValid.HasMX)
	fmt.Printf("Has SPF: %v\n", isValid.HasSPF)
	fmt.Printf("SPF Record: %s\n", isValid.SpfRecord)
	fmt.Printf("Has DMARC: %v\n", isValid.HasDMARC)
	fmt.Printf("DMARC Record: %s\n", isValid.DmarcRecord)
}
