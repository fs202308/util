// Package pubip gets your public IP address from several services
package xips

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jpillora/backoff"
)

// Version indicates the version of this package.
const Version = "1.0.0"

// MaxTries is the maximum amount of tries to attempt to one service.
const MaxTries = 3

// APIURIs is the URIs of the services.
var APIURIs = []string{
	"http://manager-customer.qixin007.com",
	"https://manager-customer.qixin007.com",
}

var client *http.Client

// Timeout sets the time limit of collecting results from different services.
var Timeout = 2 * time.Second

func GetIPBy(dest string) (net.IP, error) {
	b := &backoff.Backoff{
		Jitter: true,
	}
	if client == nil {
		return nil, fmt.Errorf("client not init yet")
	}
	req, err := http.NewRequest("GET", dest, nil)
	if err != nil {
		return nil, err
	}

	for tries := 0; tries < MaxTries; tries++ {
		resp, err := client.Do(req)
		if err != nil {
			d := b.Duration()
			time.Sleep(d)
			continue
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != 200 {
			return nil, errors.New(dest + " status code " + strconv.Itoa(resp.StatusCode) + ", body: " + string(body))
		}

		tb := strings.TrimSpace(string(body))
		ip := net.ParseIP(tb)
		if ip == nil {
			return nil, errors.New("IP address not valid: " + tb)
		}
		return ip, nil
	}

	return nil, errors.New("Failed to reach " + dest)
}

func GetIPStrBy(dest string) (string, error) {
	ip, err := GetIPBy(dest)
	return ip.String(), err
}

func worker(d string, r chan<- net.IP, e chan<- error) {
	ip, err := GetIPBy(d)
	if err != nil {
		e <- err
		return
	}
	r <- ip
}

func Client(c *http.Client) {
	client = c
}

func Get() ([]net.IP, []error) {
	var results []net.IP
	resultCh := make(chan net.IP, len(APIURIs))
	var errs []error
	errCh := make(chan error, len(APIURIs))

	for _, d := range APIURIs {
		go worker(d, resultCh, errCh)
	}
	for {
		select {
		case err := <-errCh:
			errs = append(errs, err)
		case r := <-resultCh:
			results = append(results, r)
		case <-time.After(Timeout):
			return results, errs
		}
	}
}
