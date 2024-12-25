# Email-Verifier

check- domain, hasDMARC,dmarcRecord, hasMx ,hasSPF, spfRecord

DMARC (Domain-based Message Authentication, Reporting, and Conformance) is an email authentication protocol that helps protect a domain from being used in phishing and spoofing attacks. It builds on two existing email authentication techniques, SPF (Sender Policy Framework) and DKIM (DomainKeys Identified Mail), to provide domain owners with a mechanism to protect their domain from unauthorized use.

How DMARC Works:
1.Policy Specification:

The domain owner publishes a DMARC policy in their DNS records. The policy specifies how to handle emails that fail authentication checks.
The policy can include:
none (monitoring only),
quarantine (mark as spam or place in quarantine),
reject (block the email).

2.Authentication Checks:

SPF: Ensures the email is sent from an authorized server.
DKIM: Confirms that the email has not been altered in transit and is legitimately signed.

3.Alignment Requirement:

DMARC checks if the "From" address matches the domain authenticated by SPF or DKIM. This is called alignment.
Two types:
Relaxed alignment: Allows subdomains to match.
Strict alignment: Requires an exact match.

4.Action on Failure:

Emails failing the SPF/DKIM checks and alignment are treated according to the domain’s DMARC policy (none, quarantine, or reject).

5.Reporting:

DMARC provides two types of reports:
Aggregate reports: Summarized data about authentication results.
Failure reports: Detailed information about individual failures.
Benefits of DMARC:
Protects Your Brand: Prevents your domain from being used in phishing and spoofing.
Improves Deliverability: Authenticated emails are more likely to be trusted and delivered.
Provides Visibility: Offers detailed reports about how your domain is being used in email.
