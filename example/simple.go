package main

import (
	"fmt"

	"github.com/bgv/redisurl"
)

func main() {
	// Parse redis url
	url, err := redisurl.Parse("redis://:VerySecret@127.0.0.1:6379/5")

	// Check for error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("Got Redis Info - Host: %s, Port: %s, Password: %s, Db: %s\n", url.Host, url.Port, url.Pass, url.Db))
}
