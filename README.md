# Gominatim - Go library to access nominatim geocoding services

## Geocoding? WTF?

If you want to determine the coordinates of a certain location by only having its
name, you can do this via a geocoding service. If you want to do this in go, you
probably want to use gominatim to do it.

## License

LGPLv3

## Contributions

Are welcome

## Examples


```
package main

import(
    "github.com/grindhold/gominatim"
    "fmt"
)


func main(){
    gominatim.SetServer("http://nominatim.openstreetmap.org/search")

    //Get by a Querystring
    qry := new(gominatim.Query)
    qry.Q="Hamburg"
    resp, _:= qry.Get() // Returns []gominatim.Result
    fmt.Println(resp[0].Lat, resp[0].Lon)

    //Get by City
    qry = &gominatim.Query{
         City:"Berlin",
    }
    resp, _ = qry.Get()
    fmt.Println(resp[0].Lat, resp[0].Lon)
}
```
