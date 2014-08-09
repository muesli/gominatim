package main

import (
	"fmt"
	"github.com/grindhold/gominatim"
)

func main() {
	gominatim.SetServer("http://nominatim.openstreetmap.org/")

	//Get by a Querystring
	qry := new(gominatim.SearchQuery)
	qry.Q = "Hamburg"
	resp, _ := qry.Get() // Returns []gominatim.Result
	fmt.Printf("Found location: %s (%s, %s)\n", resp[0].DisplayName, resp[0].Lat, resp[0].Lon)

	//Get by City
	qry = &gominatim.SearchQuery{
		City: "Berlin",
	}
	resp, _ = qry.Get()
	fmt.Printf("Found location: %s (%s, %s)\n", resp[0].DisplayName, resp[0].Lat, resp[0].Lon)

	//Reverse Geocoding
	rqry := new(gominatim.ReverseQuery)
	rqry.Lat = "52.5170365"
	rqry.Lon = "13.3888599"
	rresp, _ := rqry.Get()
	fmt.Printf("Found %s\n", rresp.DisplayName)
}
