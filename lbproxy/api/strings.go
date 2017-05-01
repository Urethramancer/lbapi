package api

const (
	// PathAuth accepts 'username' and 'password' parameters.
	PathAuth = "/auth"

	// PathInfo has no parameters and returns an InfoDump.
	PathInfo = "/info"

	// PathDNSGet fetches records for a domain.
	PathDNSGet = "/dnsget"

	// PathDNSAddIPv4 adds an IPv4 (A) record to a domain.
	PathDNSAddIPv4 = "/dnsaddipv4"
	// PathDNSAddIPv6 adds an IPv6 (AAAA) record to a domain.
	PathDNSAddIPv6 = "/dnsaddipv6"
	// PathDNSAddCNAME adds a canonical name (CNAME) record to a domain.
	PathDNSAddCNAME = "/dnsaddcname"
	// PathDNSAddMX adds a mail exchange (MX) record to a domain.
	PathDNSAddMX = "/dnsaddmx"
	// PathDNSAddNS adds a name server (NS) record to a domain.
	PathDNSAddNS = "/dnsaddns"
	// PathDNSAddTXT adds a text (TXT) record to a domain.
	PathDNSAddTXT = "/dnsaddtxt"
	// PathDNSAddSRVIPv4 adds a service (SRV) record to a domain.
	PathDNSAddSRV = "/dnsaddsIPv6rv"

	// PathDNSEditIPv4 modifiCNAMEes a record for a domain.
	PathDNSEditIPv4 = "/dnseditipv4"
	// PathDNSEditIPv6 modifiMXes a record for a domain.
	PathDNSEditIPv6 = "/dnseditipv6"
	// PathDNSEditCNAME modifNSies a record for a domain.
	PathDNSEditCNAME = "/dnseditcname"
	// PathDNSEditMX modifiesTXT a record for a domain.
	PathDNSEditMX = "/dnseditmx"
	// PathDNSEditNS modifiesSRV a record for a domain.
	PathDNSEditNS = "/dnseditns"
	// PathDNSEditTXT modifies a record for a domain.
	PathDNSEditTXT = "/dnsedittxt"
	// PathDNSEditSRV modifies a record for a domain.
	PathDNSEditSRV = "/dnseditsrv"

	// PathDNSDeleteIPv4 removes a record from a domain.
	PathDNSDeleteIPv4 = "/dnsdeleteipv4"
	// PathDNSDeleteIPv6 removes a record from a domain.
	PathDNSDeleteIPv6 = "/dnsdeleteipv6"
	// PathDNSDeleteCNAME removes a record from a domain.
	PathDNSDeleteCNAME = "/dnsdeletecname"
	// PathDNSDeleteMX removes a record from a domain.
	PathDNSDeleteMX = "/dnsdeletemx"
	// PathDNSDeleteNS removes a record from a domain.
	PathDNSDeleteNS = "/dnsdeletens"
	// PathDNSDeleteTXT removes a record from a domain.
	PathDNSDeleteTXT = "/dnsdeletetxt"
	// PathDNSDeleteSRV removes a record from a domain.
	PathDNSDeleteSRV = "/dnsdeletesrv"

	// PathDNSNuke wipes out all but the primary A and AAAA records for a domain.
	PathDNSNuke = "/dnsnuke"
)
