package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type RequestLog struct {
	Address   string
	Timestamp time.Time
	HTTP      struct {
		Verb, Path, Version string
		StatusCode          uint16
	}
	ResponseTime time.Duration
	UserAgent    string
}

var requestRegex *regexp.Regexp = regexp.MustCompile(
	"(?P<address>\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}) - - " +
		"\\[(?P<timestamp>\\d{1,2}/\\w+/\\d{4}:\\d{2}:\\d{2}:\\d{2} [-+]\\d{4})\\] " +
		"\"(?P<verb>GET|POST|PUT|PATCH|DELETE) (?P<path>[^ ]+) HTTP/(?P<version>\\d+\\.\\d+)\" " +
		"(?P<status>\\d{3}) (?P<ms>\\d+) \"[^\"]*\" \"(?P<user_agent>[^\"]*)\"")

func parseRequest(line string) (r *RequestLog) {
	match := requestRegex.FindStringSubmatch(line)
	if match == nil {
		return
	}

	t, err := time.Parse("02/January/2006:15:04:05 -0700", match[requestRegex.SubexpIndex("timestamp")])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	status, err := strconv.Atoi(match[requestRegex.SubexpIndex("status")])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ms, err := strconv.Atoi(match[requestRegex.SubexpIndex("ms")])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r = new(RequestLog)
	r.Address = match[requestRegex.SubexpIndex("address")]
	r.Timestamp = t
	r.HTTP.Verb = match[requestRegex.SubexpIndex("verb")]
	r.HTTP.Path = match[requestRegex.SubexpIndex("path")]
	r.HTTP.Version = match[requestRegex.SubexpIndex("version")]
	r.HTTP.StatusCode = uint16(status)
	r.ResponseTime = time.Duration(ms) * time.Millisecond
	r.UserAgent = match[requestRegex.SubexpIndex("user_agent")]

	return
}

func main() {
	line := `38.99.236.50 - - [20/May/2015:21:05:31 +0000] "GET /favicon.ico HTTP/1.1" 200 3638 "-" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.107 Safari/537.36"`

	r := parseRequest(line)
	fmt.Printf("%+v\n", r)
}
