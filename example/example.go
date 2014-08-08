package main

import (
    "fmt"
    "github.com/grindhold/gominatim"
)

func main() {
    gominatim.SetServer("http://nominatim.openstreetmap.org/search")

    //Get by a Querystring
    qry := new(gominatim.Query)
    qry.Q = "Hamburg"
    resp, _ := qry.Get() // Returns []gominatim.Result
    fmt.Printf("Found location: %s (%s, %s)\n", resp[0].DisplayName, resp[0].Lat, resp[0].Lon)

    //Get by City
    qry = &gominatim.Query{
        City: "Berlin",
    }
    resp, _ = qry.Get()
    fmt.Printf("Found location: %s (%s, %s)\n", resp[0].DisplayName, resp[0].Lat, resp[0].Lon)
}
