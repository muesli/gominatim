/*
 *    Copyright (C) 2014 Daniel 'grindhold' Brendle
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Lesser General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Lesser General Public License for more details.
 *
 *    You should have received a copy of the GNU Lesser General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Daniel 'grindhold' Brendle <grindhold@skarphed.org>
 */

package gominatim

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "errors"
)

var (
    server string
)

type ResultList struct {
    Results []*Result
}

type Address struct {
    House string `json:"house_number,omitempty"`
    Road string `json:"road,omitempty"`
    Village string `json:"village,omitempty"`
    Town string `json:"town,omitempty"`
    City string `json:"city,omitempty"`
    County string `json:"state_district,omitempty"`
    Postcode string `json:"postcode,omitempty"`
    Country string `json:"country,omitempty"`
    CountryCode string `json:"country_code,omitempty"`
}

type Result struct {
    PlaceId string `json:"place_id,omitempty"`
    License string `json:"license,omitempty"`
    OsmType string `json:"osm_type,omitempty"`
    OsmId   string `json:"osm_id,omitempty"`
    Boundingbox []string `json:"boundingbox,omitempty"`
    Polygonpoints [][]string `json:"polygonpoints,omitempty"`
    Lat string `json:"lat,omitempty"`
    Lon string `json:"lon,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Class string `json:"class,omitempty"`
    Type string `json:"type,omitempty"`
    Address Address `json:"address,omitempty"`
    Importance float32 `json:"importance",omitempty"`
}

type Query struct {
    JsonCallback interface{}
    AcceptLanguage string
    Q string
    Street string
    City string
    County string
    State string
    Postalcode string
    Countrycodes []string
    Viewbox string
    Bounded bool
    Polygon bool
    Addressdetails bool
    Email string
    ExcludePlaceIds []string
    Limit int
    PolygonGeojson bool
    PolygonKml bool
    PolygonText bool
    PolygonSvg bool
}

func (q *Query) specificFieldsUsed () bool {
    return q.Street != "" || q.City != "" || q.County != "" || q.State != "" || q.Postalcode != ""
}

func (q *Query) buildQuery() (string, error) {
    s := server
    s = s+"?format=json"
    if q.JsonCallback != nil {
        cb, err := json.Marshal(q.JsonCallback)
        if err != nil {
            return "", err
        }
        s += "&json_callback="+string(cb)
    }
    if q.AcceptLanguage != "" {
        s +="&accept_language="+q.AcceptLanguage
    }
    if q.Q != "" {
        s +="&q="+q.Q
    } else {
        if q.specificFieldsUsed() {
            if q.Street != "" {
               s += "&street="+q.Street
            }
            if q.City != "" {
                s += "&city="+q.City
            }
            if q.County != "" {
                s += "&county="+q.County
            }
            if q.State != "" {
                s += "&state="+q.State
            }
            if q.Postalcode != "" {
                s += "&postalcode="+q.Postalcode
            }
        } else {
            return "", errors.New("You must use either Q or one or more of Street, City, County, State, Postalcode. The latter will be ignored if the further is used.")
        }
    }
    if q.Countrycodes!= nil && len(q.Countrycodes) > 0 {
        als := ""
        first := true
        for i := range(q.Countrycodes) {
            if !first {
                als = als+","
            }
            als = als+q.Countrycodes[i]
            if first {
                first=false
            }
        }
        s += "&countrycodes="+als
    }
    if q.Viewbox != "" {
        s += "&viewbox="+q.Viewbox
    }
    if q.Bounded {
        s += "&bounded=1"
    }else{
        s += "&bounded=0"
    }
    if q.Polygon {
        s += "&polygon=1"
    }else{
        s += "&polygon=0"
    }
    if q.Addressdetails {
        s += "&addressdetails=1"
    }else{
        s += "&addressdetails=0"
    }
    if q.Email != ""{
        s += "&email="+ q.Email
    }
    if q.ExcludePlaceIds != nil && len(q.ExcludePlaceIds) > 0 {
        als := ""
        first := true
        for i := range(q.ExcludePlaceIds) {
            if !first {
                als = als+","
            }
            als = als+q.ExcludePlaceIds[i]
            if first {
                first=false
            }
        }
        s += "&exclude_place_ids="+als
    }
    if q.Limit > 0 {
        s += "&limit="+string(q.Limit)
    }
    if q.PolygonGeojson {
        s += "&polygon_geojson=1"
    }else{
        s += "&polygon_geojson=0"
    }
    if q.PolygonKml {
        s += "&polygon_kml=1"
    }else{
        s += "&polygon_kml=0"
    }
    if q.PolygonSvg {
        s += "&polygon_svg=1"
    }else{
        s += "&polygon_svg=0"
    }
    if q.PolygonText {
        s += "&polygon_text=1"
    }else{
        s += "&polygon_text=0"
    }
    return s, nil
}

func (q *Query) Get () ([]Result, error) {
    querystring, err := q.buildQuery()
    if err != nil {
        return nil, err
    }
    resp, err := http.Get(querystring)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    result := make([]Result,0)
    err = json.Unmarshal(body, &result)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func SetServer(srv string) {
    server = srv
}
