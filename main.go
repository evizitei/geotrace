package main

import (
	"net/http"

	"code.google.com/p/gorest"
)

func main() {
	gorest.RegisterService(new(GeoIpService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":3001", nil)
	//ip := "81.2.69.142"
	//ip := "50.160.16.241"
}

type GeoIpService struct {
	gorest.RestService `root:"/geoip" consumes:"application/json" produces:"application/json"`
	lookup             gorest.EndPoint `method:"GET" path:"/lookup/{ip:string}" output:"string"`
}

func (serv GeoIpService) Lookup(ip string) string {
	return countryLookup(ip)
}
