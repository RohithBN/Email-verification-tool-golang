# Email Domain Verification Tool

A Go-based tool for verifying email domain configurations by checking essential email authentication records (MX, SPF, and DMARC).

## Overview

This tool helps verify if a domain is properly configured for email services by checking three critical DNS records:
- MX (Mail Exchange) records
- SPF (Sender Policy Framework) records
- DMARC (Domain-based Message Authentication, Reporting, and Conformance) records

## Key Terms

### MX (Mail Exchange) Records
MX records specify the mail servers responsible for accepting email messages on behalf of a domain. They are crucial for email delivery and routing. Without valid MX records, a domain cannot receive emails.

### SPF (Sender Policy Framework)
SPF is an email authentication method that helps prevent email spoofing. It specifies which mail servers are authorized to send emails on behalf of your domain. SPF records are published in the domain's DNS as TXT records and typically start with "v=spf1".

### DMARC (Domain-based Message Authentication, Reporting, and Conformance)
DMARC is an email authentication protocol that builds upon SPF and DKIM. It helps domain owners prevent unauthorized use of their domain for email spoofing, phishing scams, and other malicious activities. DMARC records are published as TXT records under the "_dmarc" subdomain.

## Features

- Validates the existence of MX records
- Checks for valid SPF records and retrieves the record content
- Verifies DMARC configuration and retrieves the policy
- Returns structured data about the domain's email configuration



