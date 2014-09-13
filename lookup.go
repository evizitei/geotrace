package main

import (
	"log"
	"net"

	"code.google.com/p/gcfg"

	"github.com/oschwald/geoip2-golang"
)

type lookupService interface {
	Country(net.IP) (*geoip2.Country, error)
}

func countryLookup(ipString string) string {

	filename, confErr := getDbFile()
	if confErr != nil {
		log.Fatal(confErr)
	}

	db, geoErr := geoip2.Open(filename)
	if geoErr != nil {
		log.Fatal(geoErr)
	}

	defer db.Close()
	return ipToCountry(db, ipString)
}

func ipToCountry(db lookupService, ipString string) string {
	ip := net.ParseIP(ipString)
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	return record.Country.Names["en"]
}

type Config struct {
	Geotrace struct {
		Dbfile string
	}
}

func getDbFile() (string, error) {
	var config Config
	err := gcfg.ReadFileInto(&config, "geotrace.gcfg")
	if err != nil {
		return "", err
	}
	filename := config.Geotrace.Dbfile
	return filename, nil
}
