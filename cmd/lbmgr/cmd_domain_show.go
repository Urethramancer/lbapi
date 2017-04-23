package main

import (
	"time"

	"github.com/Urethramancer/lbapi/common"
)

// DomainShowCmd arguments.
type DomainShowCmd struct {
	Args struct {
		Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name."`
	} `positional-args:"true"`
}

// Execute the show command.
func (cmd *DomainShowCmd) Execute(args []string) error {
	d, err := client.Domain(cmd.Args.Domain)
	if err != nil {
		return err
	}

	pr(common.ANSI_YELLOW+"%s (%s) - "+okColour(d.Status == "Active")+"%s"+common.ANSI_NORMAL, d.Description, d.TypeName, d.Status)
	cid := d.CustomerID
	cust, err := client.CustomerByID(cid)
	prn("Owner: ")
	if err != nil {
		pr(err.Error())
	} else {
		pr("%s (%d) - %s\nTotal receipts: $%s", cust.Name, cust.ID, cust.Status, cust.TotalReceipts)
	}

	prn("Created: %v\nLast modified: %v\nExpires: ", d.CreationDT, d.Timestamp)
	pr(okColour(!time.Now().After(d.Endtime))+"%v"+common.ANSI_NORMAL, d.Endtime)

	pr("Autorenew: %v", d.Autorenew)
	pr("Reseller lock: "+okColour(d.ResellerLock)+"%v"+common.ANSI_NORMAL, d.ResellerLock)
	pr("Transfer lock: "+okColour(d.TransferLock)+"%v"+common.ANSI_NORMAL, d.TransferLock)
	return nil
}
