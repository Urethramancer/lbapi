package api

const (
	// PathAuth accepts 'username' and 'password' parameters.
	PathAuth = "/auth"
	// PathInfo has no parameters and returns an InfoDump.
	PathInfo = "/info"
	// PathDNSGet fetches records of one or more types for a domain.
	PathDNSGet = "/dnsget"
	// PathDNSAdd adds a record to a domain.
	PathDNSAdd = "/dnsadd"
	// PathDNSEdit modifies a record for a domain.
	PathDNSEdit = "/dnsedit"
	// PathDNSDelete removes a record from a domain.
	PathDNSDelete = "/dnsdelete"
	// PathDNSNuke wipes out all but the primary A and AAAA records for a domain.
	PathDNSNuke = "/dnsnuke"
)
