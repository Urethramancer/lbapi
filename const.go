package lbapi

const (
	// APIURL is the default LogicBoxes API URL.
	APIURL = "https://httpapi.com/api/"
	// APITESTURL is the LogicBoxes API URL used for test accounts.
	APITESTURL = "https://test.httpapi.com/api/"

	// apiCustomersAuthenticate fetches CustomerDetails if authentication is successful.
	apiCustomersAuthenticate = "api/customers/authenticate.json"
	// apiCustomersDetailsByID fetches CustomerDetails by ID.
	apiCustomersDetailsByID = "api/customers/details-by-id.json"
	// apiDomainsSearch is the source of all domain order information.
	apiDomainsSearch = "api/domains/search.json"
	apiDNSAddIPv4    = "api/dns/manage/add-ipv4-record.json"
	apiDNSAddIPv6    = "api/dns/manage/add-ipv6-record.json"
	apiDNSAddCNAME   = "api/dns/manage/add-cname-record.json"
	apiDNSAddMX      = "api/dns/manage/add-mx-record.json"
	apiDNSAddNS      = "api/dns/manage/add-ns-record.json"
	apiDNSAddTXT     = "api/dns/manage/add-txt-record.json"
	apiDNSAddSRV     = "api/dns/manage/add-srv-record.json"

	apiDNSDeleteIPv4  = "api/dns/manage/delete-ipv4-record.json"
	apiDNSDeleteIPv6  = "api/dns/manage/delete-ipv6-record.json"
	apiDNSDeleteCNAME = "api/dns/manage/delete-cname-record.json"
	apiDNSDeleteMX    = "api/dns/manage/delete-mx-record.json"
	apiDNSDeleteNS    = "api/dns/manage/delete-ns-record.json"
	apiDNSDeleteTXT   = "api/dns/manage/delete-txt-record.json"
	apiDNSDeleteSRV   = "api/dns/manage/delete-srv-record.json"

	apiDNSUpdateIPv4  = "api/dns/manage/update-ipv4-record.json"
	apiDNSUpdateIPv6  = "api/dns/manage/update-ipv6-record.json"
	apiDNSUpdateCNAME = "api/dns/manage/update-cname-record.json"
	apiDNSUpdateMX    = "api/dns/manage/update-mx-record.json"
	apiDNSUpdateNS    = "api/dns/manage/update-ns-record.json"
	apiDNSUpdateTXT   = "api/dns/manage/update-txt-record.json"
	apiDNSUpdateSRV   = "api/dns/manage/update-srv-record.json"
	apiDNSUpdateSOA   = "api/dns/manage/update-soa-record.json"
)

type maplist map[string]interface{}
