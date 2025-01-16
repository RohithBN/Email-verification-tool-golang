package main

import (
	"log"
	"net"
	"strings"
)

func (domain *DomainUrl) isValidDomain() *Domain {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord string
	var dmarcRecord string
	mxRecord, err := net.LookupMX(domain.DomainUrl)
	if err != nil {
		log.Printf("Error:%v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}
	txtRecord, err := net.LookupTXT(domain.DomainUrl)
	if err != nil {
		log.Printf("Error:%v\n", err)
	}
	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain.DomainUrl)
	if err != nil {
		log.Printf("Error:%v\n", err)
	}
	for _, records := range dmarcRecords {
		if strings.HasPrefix(records, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = records
			break
		}
	}
	return &Domain{
		DomainName:  domain.DomainUrl,
		HasMX:       hasMX,
		HasSPF:      hasSPF,
		SpfRecord:   spfRecord,
		HasDMARC:    hasDMARC,
		DmarcRecord: dmarcRecord,
	}

}
