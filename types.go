package main

type Domain struct {
	DomainName  string `json:"domain_name"`
	HasMX       bool   `json:"hasmx"`
	HasSPF      bool   `json:"haspf"`
	SpfRecord   string `json:"spfrecord"`
	HasDMARC    bool   `json:"hasdmarc"`
	DmarcRecord string `json:"dmarcRecord"`
}

type DomainUrl struct {
	DomainUrl string `json:"domain_url"`
}
