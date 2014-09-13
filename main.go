package main

import (
	"net/http"

	"code.google.com/p/gorest"
)

func main() {
	gorest.RegisterService(new(GeoIpService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":3001", nil)
}

type GeoIpService struct {
	gorest.RestService `root:"/geoip" consumes:"application/json" produces:"application/json"`
	lookup             gorest.EndPoint `method:"GET" path:"/lookup/{ip:string}" output:"string"`
}

func (serv GeoIpService) Lookup(ip string) string {
	return countryLookup(ip)
}
