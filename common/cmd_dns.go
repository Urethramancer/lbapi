package common

import "github.com/Urethramancer/lbapi"

type DNSClient interface {
	DNSActive(string) bool

	//
	// Add section
	//
	AddARecord(domain, address, host string, ttl int64, six bool) error
	AddCNAME(domain, value, host string, ttl int64) error
	AddMX(domain, value, host string, ttl int64, priority uint16) error
	AddNS(domain, value, host string, ttl int64, priority uint16) error
	AddTXT(domain, value, host string, ttl int64, priority uint16) error
	AddSRV(domain, value, host string, ttl int64, priority, port, weight uint16) error

	//
	// Get section
	//
	GetDNSRecords(domain, value, host, t string, page int) (*lbapi.DNSRecordList, error)

	//
	// Edit section
	//
	EditARecord(domain, oldip, newip, host string, ttl int64, six bool) error
	EditCNAME(domain, oldip, newip, host string, ttl int64) error
	EditMX(domain, oldip, newip, host string, ttl int64, priority uint16) error
	EditNS(domain, oldip, newip, host string, ttl int64) error
	EditTXT(domain, oldip, newip, host string, ttl int64) error
	EditSRV(domain, oldval, newval, host string, ttl int64, priority, port, weight uint) error
	EditSOA(domain, person string, refresh, retry, expire, ttl int64) error

	//
	// Delete section
	//
	DeleteARecord(domain, value, host string, six bool) error
	DeleteCNAME(domain, value, host string) error
	DeleteMX(domain, value, host string) error
	DeleteNS(domain, value, host string) error
	DeleteTXT(domain, value, host string) error
	DeleteSRV(domain, value, host string, port, weight uint16) error
}

var client DNSClient

func SetClient(c DNSClient) {
	client = c
}

// DNSCmd arguments and sub-commands.
type DNSCmd struct {
	Status DNSStatusCmd `command:"status" description:"Shows DNS management status for a domain by order ID, and activates it if not yet enabled."`
	Get    DNSGetCmd    `command:"get" description:"Get one type of DNS record for a domain."`
	Add    DNSAddCmd    `command:"add" description:"Add a DNS record for a domain."`
	Delete DNSDeleteCmd `command:"delete" description:"Delete a DNS record from a domain." alias:"del" alias:"rm"`
	Edit   DNSEditCmd   `command:"edit" description:"Edit a DNS record." alias:"ch" alias:"change" alias:"ed" alias:"mod"`
	Nuke   DNSNukeCmd   `command:"nuke" description:"Clear out all DNS records except the primary A/AAAA for a domain." alias:"wipeout"`
}

// DNSAddArgs are the default arguments for some record-adding sub-commands.
type DNSAddArgs struct {
	Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	Value  string `required:"true" positional-arg-name:"VALUE" description:"IP address, or FQDN for CNAME."`
	Host   string `positional-arg-name:"HOST" description:"Host name."`
	TTL    int64  `positional-arg-name:"TTL" description:"Time to live (seconds)."`
}

// DNSAddArgsPri are the default arguments for record-adding sub-commands with priority.
type DNSAddArgsPri struct {
	Domain   string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	Value    string `required:"true" positional-arg-name:"VALUE" description:"IP address, or FQDN for CNAME."`
	Host     string `positional-arg-name:"HOST" description:"Host name."`
	TTL      int64  `positional-arg-name:"TTL" description:"Time to live (seconds)."`
	Priority uint16 `positional-arg-name:"PRIORITY" description:"Priority of record. Default is 0 (most important)."`
}

// DNSEditArgs are the default arguments for some record-editing sub-commands.
type DNSEditArgs struct {
	Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	Old    string `required:"true" positional-arg-name:"OLD VALUE" description:"IP address/name/data to change from."`
	New    string `required:"true" positional-arg-name:"NEW VALUE" description:"New IP address/name/data."`
	Host   string `positional-arg-name:"HOST" description:"Host name."`
	TTL    int64  `positional-arg-name:"TTL" description:"Time to live (seconds)."`
}

// DNSEditArgsPri are the default arguments for record-editing sub-commands with priority.
type DNSEditArgsPri struct {
	Domain   string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	Old      string `required:"true" positional-arg-name:"OLDIP" description:"IP address to change from."`
	New      string `required:"true" positional-arg-name:"NEWIP" description:"New IP address."`
	Host     string `positional-arg-name:"HOST" description:"Host name."`
	TTL      int64  `positional-arg-name:"TTL" description:"Time to live (seconds)."`
	Priority uint16 `positional-arg-name:"PRIORITY" description:"Priority of record. Default is 0 (most important)."`
}

// DNSGetArgs are the default arguments for record-fetching sub-commands.
type DNSGetArgs struct {
	Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	Value  string `positional-arg-name:"VALUE" description:"IP address, or FQDN for CNAME."`
	Host   string `positional-arg-name:"HOST" description:"Host name."`
}

// DNSDeleteArgs are the default args for most record-deletion sub-commands.
type DNSDeleteArgs struct {
	Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	Value  string `required:"true" positional-arg-name:"VALUE" description:"IP address, or FQDN for CNAME/MX."`
	Host   string `positional-arg-name:"HOST" description:"Host name."`
}

// DNSDeleteArgsAll are default args for some special record-deletion sub-commands.
type DNSDeleteArgsAll struct {
	Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	Value  string `required:"true" positional-arg-name:"VALUE" description:"IP address, or FQDN for CNAME/MX."`
	Host   string `required:"true" positional-arg-name:"HOST" description:"Host name."`
}

// DNSStatusCmd arguments.
type DNSStatusCmd struct {
	Args struct {
		ID string `required:"true" positional-arg-name:"ID" description:"Domain order ID."`
	} `positional-args:"true"`
}

// Execute the status command.
func (cmd *DNSStatusCmd) Execute(args []string) error {
	if client.DNSActive(cmd.Args.ID) {
		pr("DNS is active.")
	} else {
		pr("Non-existent or not yet activated.")
	}
	return nil
}
