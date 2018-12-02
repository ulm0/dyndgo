package data

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/dnsimple/dnsimple-go/dnsimple"
	yaml "gopkg.in/yaml.v2"
)

// Data template for YAML file
type Data struct {
	Credentials map[string]string   `yaml:"credentials"`
	Zones       map[string][]string `yaml:"zones"`
}

// ReadData gets data from YAML file
func (d *Data) ReadData(dataFile string) (*Data, error) {
	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, d)
	return d, nil
}

// UpdateDomains updates the IP for the given domains
func (d *Data) UpdateDomains(ip string) error {
	token := d.Credentials["token"]
	client := dnsimple.NewClient(dnsimple.NewOauthTokenCredentials(token))
	user, err := client.Identity.Whoami()
	if err != nil {
		return fmt.Errorf("Unable to identify who you are")
	}
	fmt.Printf("Identified as %s\n", user.Data.Account.Email)
	accountID := strconv.Itoa(int(user.Data.Account.ID))

	for zone, domains := range d.Zones {
		zoneResp, err := client.Zones.ListRecords(accountID, zone, nil)
		if err != nil {
			return fmt.Errorf("Failed to get zone %s", zone)
		}
		for _, domainResp := range zoneResp.Data {
			for _, domain := range domains {
				if domainResp.Name == domain {
					if domainResp.Content == ip {
						fmt.Printf("- %s.%s is up to date\n", domain, zone)
					} else if domainResp.Type == "A" && domainResp.Content != ip {
						updatedRecord := dnsimple.ZoneRecord{
							ID:      domainResp.ID,
							Type:    domainResp.Type,
							Name:    domainResp.Name,
							Content: ip,
						}
						updateResp, err := client.Zones.UpdateRecord(accountID, zone, domainResp.ID, updatedRecord)
						if err != nil {
							return fmt.Errorf("error updating record %s", err)
						}
						fmt.Printf("- %s.%s updated from %s to %s\n", domain, zone, domainResp.Content, updateResp.Data.Content)
					}
				}
			}
		}
	}
	return nil
}
