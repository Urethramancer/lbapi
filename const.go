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
	APIDomainsSearch = "api/domains/search.json"
	APIDNSAddIPv4    = "api/dns/manage/add-ipv4-record.json"
	APIDNSAddIPv6    = "api/dns/manage/add-ipv6-record.json"
	APIDNSAddCNAME   = "api/dns/manage/add-cname-record.json"
	APIDNSAddMX      = "api/dns/manage/add-mx-record.json"
	APIDNSAddNS      = "api/dns/manage/add-ns-record.json"
	APIDNSAddTXT     = "api/dns/manage/add-txt-record.json"
	APIDNSAddSRV     = "api/dns/manage/add-srv-record.json"

	APIDNSDeleteIPv4  = "api/dns/manage/delete-ipv4-record.json"
	APIDNSDeleteIPv6  = "api/dns/manage/delete-ipv6-record.json"
	APIDNSDeleteCNAME = "api/dns/manage/delete-cname-record.json"
	APIDNSDeleteMX    = "api/dns/manage/delete-mx-record.json"
	APIDNSDeleteNS    = "api/dns/manage/delete-ns-record.json"
	APIDNSDeleteTXT   = "api/dns/manage/delete-txt-record.json"
	APIDNSDeleteSRV   = "api/dns/manage/delete-srv-record.json"

	APIDNSUpdateIPv4  = "api/dns/manage/update-ipv4-record.json"
	APIDNSUpdateIPv6  = "api/dns/manage/update-ipv6-record.json"
	APIDNSUpdateCNAME = "api/dns/manage/update-cname-record.json"
	APIDNSUpdateMX    = "api/dns/manage/update-mx-record.json"
	APIDNSUpdateNS    = "api/dns/manage/update-ns-record.json"
	APIDNSUpdateTXT   = "api/dns/manage/update-txt-record.json"
	APIDNSUpdateSRV   = "api/dns/manage/update-srv-record.json"
	APIDNSUpdateSOA   = "api/dns/manage/update-soa-record.json"
)

type maplist map[string]interface{}
