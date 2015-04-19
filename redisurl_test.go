package redisurl

import (
	"reflect"
	"testing"
)

var testUriCases = []struct {
	uriString   string
	expected    *RedisUrl
	description string
}{
	{
		uriString: "redis://localhost:6379",
		expected: &RedisUrl{
			Host:     "localhost",
			Port:     "6379",
			Hostname: "localhost:6379",
			Db:       "0",
			Pass:     "",
		},
		description: "typical case",
	},
	{
		uriString: "redis://redisproviders.com:12345/5",
		expected: &RedisUrl{
			Host:     "redisproviders.com",
			Port:     "12345",
			Hostname: "redisproviders.com:12345",
			Db:       "5",
			Pass:     "",
		},
		description: "host, port, db",
	},
	{
		uriString: "redis://10.0.0.1",
		expected: &RedisUrl{
			Host:     "10.0.0.1",
			Port:     "6379",
			Hostname: "10.0.0.1:6379",
			Db:       "0",
			Pass:     "",
		},
		description: "host-only ip (guess default port and db)",
	},
	{
		uriString: "redis://10.0.0.1",
		expected: &RedisUrl{
			Host:     "10.0.0.1",
			Port:     "6379",
			Hostname: "10.0.0.1:6379",
			Db:       "0",
			Pass:     "",
		},
		description: "with password (no username)",
	},
	{
		uriString: "redis://josh:click@redis.codechimp.io:999/1",
		expected: &RedisUrl{
			Host:     "redis.codechimp.io",
			Port:     "999",
			Hostname: "redis.codechimp.io:999",
			Db:       "1",
			Pass:     "click",
		},
		description: "everything URI (username should be ignored)",
	},
}

func TestParse(t *testing.T) {
	for _, tt := range testUriCases {
		expected := tt.expected
		actual, _ := Parse(tt.uriString)
		if reflect.DeepEqual(expected, actual) {
			t.Logf("PASS: %s", tt.description)
		} else {
			t.Errorf("FAIL: %s, expected: %+v, actual: %+v", tt.description, expected, actual)
		}
	}
}
