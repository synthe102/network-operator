package config

import (
	"context"
	"os"

	"github.com/paultyng/go-unifi/unifi"
)

func NewUnifiClient() (*unifi.Client, error) {
	user := os.Getenv("UNIFI_USER")
	password := os.Getenv("UNIFI_PASSWORD")
	url := os.Getenv("UNIFI_URL")

	var c unifi.Client
	if err := c.SetBaseURL(url); err != nil {
		return nil, err
	}
	if err := c.Login(context.Background(), user, password); err != nil {
		return nil, err
	}
	return &c, nil
}
