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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type reverseAPIResult struct {
	ReverseResult
	Error string `json:"error"`
}

type ReverseResult struct {
	PlaceId     int64   `json:"place_id"`
	License     string  `json:"license"`
	OsmType     string  `json:"osm_type"`
	OsmId       int64   `json:"osm_id"`
	Lat         string  `json:"lat"`
	Lon         string  `json:"lon"`
	DisplayName string  `json:"display_name"`
	Address     Address `json:"address"`
}

type ReverseQuery struct {
	JsonCallback   interface{}
	AcceptLanguage string
	OsmType        string
	OsmId          string
	Lat            string
	Lon            string
	Zoom           int
	AddressDetails bool
	Email          string
}

func (r *ReverseQuery) buildQuery() (string, error) {
	if server == "" {
		return "", errors.New("Server is not set. Set via gominatim.SetServer(srv string)")
	}
	s := server
	s = s + "/reverse?format=json"
	if r.AcceptLanguage != "" {
		s = s + "&accept-language=" + r.AcceptLanguage
	}
	if r.OsmType != "" {
		if r.OsmType != "N" && r.OsmType != "W" && r.OsmType != "R" {
			return "", errors.New("OsmType must be 'N', 'W' or 'R'")
		}
		s = s + "&osm_type=" + r.OsmType
	}
	if r.OsmId != "" {
		s = s + "&osm_id=" + r.OsmType
	}
	if r.Lat == "" {
		return "", errors.New("Cannot search without a latitude. Set field Lat")
	}
	s = s + "&lat=" + r.Lat
	if r.Lon == "" {
		return "", errors.New("Cannot search without a longitude. Set field Lon")
	}
	s = s + "&lon=" + r.Lon
	if r.Zoom > 18 || r.Zoom < 0 {
		return "", errors.New(fmt.Sprintf("Zoom must be within 0 and 18. %d is out of range", r.Zoom))
	}
	s = s + fmt.Sprintf("&zoom=%d", r.Zoom)
	if r.AddressDetails {
		s = s + "&addressdetails=1"
	} else {
		s = s + "&addressdetails=0"
	}
	if r.Email != "" {
		s = s + "&email=" + url.QueryEscape(r.Email)
	}
	return s, nil
}

func (r *ReverseQuery) Get() (*ReverseResult, error) {
	querystring, err := r.buildQuery()
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
	result := new(reverseAPIResult)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return &result.ReverseResult, nil
}
