package client

import (
	"fmt"
	"log"
	"strings"
)

type DnsRecord struct {
	Id      string `mikrotik:".id"`
	Name    string
	Ttl     int `mikrotik:"ttl,ttlToSeconds"`
	Address string
}

func (client Mikrotik) AddDnsRecord(name, address string, ttl int) (*DnsRecord, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := strings.Split(fmt.Sprintf("/ip/dns/static/add =name=%s =address=%s =ttl=%d", name, address, ttl), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	log.Printf("[DEBUG] /ip/dns/static/add returned %v", r)

	if err != nil {
		return nil, err
	}

	return client.FindDnsRecord(name)
}

func (client Mikrotik) FindDnsRecord(name string) (*DnsRecord, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := strings.Split(fmt.Sprintf("/ip/dns/static/print ?name=%s", name), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Found dns record: %v", r)

	record := DnsRecord{}
	err = Unmarshal(*r, &record)

	if err != nil {
		return nil, err
	}

	if record.Name == "" {
		return nil, NewNotFound(fmt.Sprintf("dns record `%s` not found", name))
	}

	return &record, nil
}

func (client Mikrotik) UpdateDnsRecord(id, name, address string, ttl int) error {
	c, err := client.getMikrotikClient()

	if err != nil {
		return err
	}
	cmd := strings.Split(fmt.Sprintf("/ip/dns/static/set =numbers=%s =name=%s =address=%s =ttl=%d", id, name, address, ttl), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)
	return err
}

func (client Mikrotik) DeleteDnsRecord(id string) error {
	c, err := client.getMikrotikClient()

	if err != nil {
		return err
	}
	cmd := strings.Split(fmt.Sprintf("/ip/dns/static/remove =numbers=%s", id), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)
	return err
}
