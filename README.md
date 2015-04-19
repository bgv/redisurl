# redisurl

Parses a [Redis](http://redis.io/) url and returns result as a struct

## Installation

Install redisurl using the "go get" command:

    $ go get github.com/bgv/redisurl

The Go distribution is redisul's only dependency.

## Usage

Import `github.com/bgv/redisurl` and use as:

```go
url, err := redisurl.Parse("redis://:v3rys3cr37@127.0.0.1:6379/5")
if err != nil {
	log.Fatal(err)
}
```

which will return the following struct
```go
// Redis Info:
url.Host: 127.0.0.1
url.Port: 6379
url.Password: v3rys3cr37
url.Db: 5
url.Hostname: 127.0.0.1:6379
```