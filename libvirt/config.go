package libvirt

import (
	"log"

	libvirt "github.com/libvirt/libvirt-go"
)

// Config struct for the libvirt-provider
type Config struct {
	URI string
}

// Client libvirt
type Client struct {
	libvirt *libvirt.Connect
}

// Client libvirt, generate libvirt client given URI
func (c *Config) Client() (*Client, error) {
	var err error
	var client Client

	client.libvirt, err = libvirt.NewConnect(c.URI)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Created libvirt client with URI=%s", c.URI)

	return &client, nil
}
