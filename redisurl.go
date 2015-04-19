package redisurl

import (
	"errors"
	"net/url"
	"strings"
)

type RedisUrl struct {
	Host     string
	Port     string
	Db       string
	Pass     string
	Hostname string
}

func Parse(uriString string) (*RedisUrl, error) {

	// Always try to parse as url
	if !strings.HasPrefix(uriString, "redis://") {
		uriString = "redis://" + uriString
	}

	// Try to parse url and return err with the proper format should be used
	url, err := url.Parse(uriString)

	if err != nil {
		return nil, errors.New("redisurl must be in the form redis://[:password@]hostname:port[/db_number]")
	}

	// Figure out Host and Port
	parts := strings.Split(url.Host, ":")

	host := parts[0]

	// Set default port
	port := "6379"

	if len(parts) > 1 {
		port = parts[1]
	}

	hostname := host + ":" + port

	// Preset db to 0 in case it's nil
	db := "0"

	if url.Path != "" {
		db = strings.Split(url.Path, "/")[1]
	}

	// check for password
	password := ""

	if url.User != nil {
		password, _ = url.User.Password()
	}

	ru := &RedisUrl{
		Host:     host,
		Port:     port,
		Hostname: hostname,
		Pass:     password,
		Db:       db,
	}

	return ru, nil
}
