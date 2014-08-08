# Gominatim - Go library to access nominatim geocoding services

## Geocoding? WTF?

If you want to determine the coordinates of a certain location by only having its
name, you can do this via a geocoding service. If you want to do this in go, you
probably want to use gominatim to do it.

## License

LGPLv3

## Features

The plan is to cover everything, this site documents:
[Nomiatim Wiki](http://wiki.openstreetmap.org/wiki/Nominatim)

 * [x] Search
 * [ ] Reverese Geocoding

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
```
