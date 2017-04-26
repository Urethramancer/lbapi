package api

const (
	// PathAuth accepts 'username' and 'password' parameters.
	PathAuth = "/auth"
	// PathInfo has no parameters and returns an InfoDump.
	PathInfo      = "/info"
	PathDNSStatus = "/dnsstatus"
	PathDNSGet    = "/dnsget"
	PathDNSAdd    = "/dnsadd"
	PathDNSEdit   = "/dnsedit"
	PathDNSDelete = "/dnsdelete"
	PathDNSNuke   = "/dnsnuke"
)
