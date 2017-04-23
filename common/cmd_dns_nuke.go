package common

import (
	"bufio"
	"os"
	"strings"

	"github.com/Urethramancer/lbapi"
)

// DNSNukeCmd arguments.
type DNSNukeCmd struct {
	Args struct {
		Domain string `required:"true" positional-arg-name:"DOMAIN" description:"Domain name to clean out."`
	} `positional-args:"true"`
}

// Execute the nuclear option.
func (cmd *DNSNukeCmd) Execute(args []string) error {
	pr("If you're absolutely certain you want to delete everything but the primary A/AAAA records, type 'NUKE IT!' below and press enter:")
	prn("> ")
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	text = strings.TrimSpace(text)
	if text != "NUKE IT!" {
		pr("Safe choice.")
		return nil
	}

	pr("OK, you asked for it!\n")
	var err error
	var dns lbapi.DNSRecords

	pr(ANSI_YELLOW + "Purging A records…" + ANSI_NORMAL)
	dns, err = getRecords(cmd.Args.Domain, "", "", "A")
	if len(dns) > 0 && err == nil {
		for _, r := range dns {
			if r.Host == "@" {
				pr(ANSI_GREEN+"Skipping"+ANSI_NORMAL+" %s (%s)", cmd.Args.Domain, r.Value)
			} else {
				err = client.DeleteARecord(cmd.Args.Domain, r.Value, r.Host, false)
				if err != nil {
					pr("Error deleting %s: %s", r.Host, err.Error())
				} else {
					pr(ANSI_RED+"Deleted"+ANSI_NORMAL+" %s.%s (%s)", r.Host, cmd.Args.Domain, r.Value)
				}
			}
		}
	}
	pr("")

	pr(ANSI_YELLOW + "Purging AAAA records…" + ANSI_NORMAL)
	dns, err = getRecords(cmd.Args.Domain, "", "", "AAAA")
	if len(dns) > 0 && err == nil {
		for _, r := range dns {
			if r.Host == "@" {
				pr(ANSI_GREEN+"Skipping"+ANSI_NORMAL+" %s (%s)", cmd.Args.Domain, r.Value)
			} else {
				err = client.DeleteARecord(cmd.Args.Domain, r.Value, r.Host, true)
				if err != nil {
					pr("Error deleting %s: %s", r.Host, err.Error())
				} else {
					pr(ANSI_RED+"Deleted"+ANSI_NORMAL+" %s.%s (%s)", r.Host, cmd.Args.Domain, r.Value)
				}
			}
		}
	}
	pr("")

	pr(ANSI_YELLOW + "Purging CNAME records…" + ANSI_NORMAL)
	dns, err = getRecords(cmd.Args.Domain, "", "", "CNAME")
	if len(dns) > 0 && err == nil {
		for _, r := range dns {
			err = client.DeleteCNAME(cmd.Args.Domain, r.Value, r.Host)
			if err != nil {
				pr("Error deleting %s: %s", r.Host, err.Error())
			} else {
				pr(ANSI_RED+"Deleted"+ANSI_NORMAL+" %s.%s", r.Host, r.Value)
			}
		}
	}
	pr("")

	pr(ANSI_YELLOW + "Purging MX records…" + ANSI_NORMAL)
	dns, err = getRecords(cmd.Args.Domain, "", "", "MX")
	if len(dns) > 0 && err == nil {
		for _, r := range dns {
			err = client.DeleteMX(cmd.Args.Domain, r.Value, r.Host)
			if err != nil {
				pr("Error deleting %s: %s", r.Value, err.Error())
			} else {
				prn(ANSI_RED+"Deleted"+ANSI_NORMAL+" %s (%d) for ", r.Value, r.Priority)
				if r.Host == "@" {
					pr(cmd.Args.Domain)
				} else {
					pr(r.Host)
				}
			}
		}
	}
	pr("")

	pr(ANSI_YELLOW + "Purging NS records…" + ANSI_NORMAL)
	dns, err = getRecords(cmd.Args.Domain, "", "", "NS")
	if len(dns) > 0 && err == nil {
		for _, r := range dns {
			err = client.DeleteNS(cmd.Args.Domain, r.Value, r.Host)
			if err != nil {
				pr("Error deleting %s: %s", r.Value, err.Error())
			} else {
				prn(ANSI_RED+"Deleted"+ANSI_NORMAL+" %s (%d) for ", r.Value, r.Priority)
				if r.Host == "@" {
					pr(cmd.Args.Domain)
				} else {
					pr(r.Host)
				}
			}
		}
	}
	pr("")

	pr(ANSI_YELLOW + "Purging TXT records…" + ANSI_NORMAL)
	dns, err = getRecords(cmd.Args.Domain, "", "", "TXT")
	if len(dns) > 0 && err == nil {
		for _, r := range dns {
			err = client.DeleteTXT(cmd.Args.Domain, r.Value, r.Host)
			if err != nil {
				pr("Error deleting %s: %s", r.Value, err.Error())
			} else {
				prn(ANSI_RED+"Deleted"+ANSI_NORMAL+" %s (%d) for ", r.Value, r.Priority)
				if r.Host == "@" {
					pr(cmd.Args.Domain)
				} else {
					pr(r.Host)
				}
			}
		}
	}
	pr("")

	pr(ANSI_YELLOW + "Purging SRV records…" + ANSI_NORMAL)
	dns, err = getRecords(cmd.Args.Domain, "", "", "SRV")
	if len(dns) > 0 && err == nil {
		for _, r := range dns {
			err = client.DeleteSRV(cmd.Args.Domain, r.Value, r.Host, r.Port, r.Weight)
			if err != nil {
				pr("Error deleting %s: %s", r.Host, err.Error())
			} else {
				pr(ANSI_RED+"Deleted"+ANSI_NORMAL+" %s:%d (%d, %d) for %s", r.Host, r.Port, r.Priority, r.Weight, r.Value)
				client.DeleteSRV(cmd.Args.Domain, r.Value, r.Host, r.Port, r.Weight)
			}
		}
	}
	pr("")

	return nil
}
