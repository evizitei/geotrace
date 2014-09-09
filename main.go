package main

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

func countryLookup(ipString string) string {

	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP(ipString)
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	return record.Country.Names["en"]
}

func main() {
	// If you are using strings that may be invalid, check that ip is not nil
	//ip := "81.2.69.142"
	ip := "50.160.16.241"
	country := countryLookup(ip)
	fmt.Printf("English country name: %v\n", country)
}
