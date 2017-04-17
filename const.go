package lbapi

const (
	// APIURL is the default LogicBoxes API URL.
	APIURL = "https://httpapi.com/api/"
	// APITESTURL is the LogicBoxes API URL used for test accounts.
	APITESTURL = "https://test.httpapi.com/api/"

	API_CUSTOMERS_DETAILS_BY_ID = "api/customers/details-by-id.json"
	// API_DOMAINS_SEARCH is the source of all domain order information.
	API_DOMAINS_SEARCH = "api/domains/search.json"
)

type maplist map[string]interface{}
