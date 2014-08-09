# Gominatim - Go library to access nominatim geocoding services

![Buildstatus](https://travis-ci.org/grindhold/gominatim.svg)

## Geocoding? WTF?

If you want to determine the coordinates of a certain location by only having its
name, you can do this via a geocoding service. If you want to do this in go, you
probably want to use gominatim to do it.

## License

[LGPLv3](https://www.gnu.org/licenses/lgpl.html )

## Features

The plan is to cover everything, this site documents:
[Nominatim Wiki](http://wiki.openstreetmap.org/wiki/Nominatim)

 * [x] Search
 * [x] Reverese Geocoding

## Contributions

Are welcome if you want to implement the Reverse Geocoding-Part

## Examples


```go
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
```
