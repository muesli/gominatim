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
	"strings"
	"testing"
)

func Test_CreateReverseQuery(t *testing.T) {
	defer SetServer("")
	SetServer("http://nominatim.openstreetmap.org")
	rqry := new(ReverseQuery)
	rqry.Lat = "52.5170365"
	rqry.Lon = "13.3888599"
	qstr, _ := rqry.buildQuery()
	if !strings.Contains(qstr, "lat=52.5170365") || !strings.Contains(qstr, "13.3888599") {
		t.Error("query does not contain longitude and latitude")
	}
}

func Test_ReverseQueryWithoutServer(t *testing.T) {
	rqry := new(ReverseQuery)
	rqry.Lat = "52.5170365"
	rqry.Lon = "13.3888599"
	_, err := rqry.buildQuery()
	if err != nil {
		if !(err.Error() == "Server is not set. Set via gominatim.SetServer(srv string)") {
			t.Error("Expecting error about unset server. Received" + err.Error())
		}
	} else {
		t.Error("Expected error about unset server. Got none.")
	}
}

func Test_OSMType(t *testing.T) {
	defer SetServer("")
	SetServer("http://nominatim.openstreetmap.org")
	rqry := new(ReverseQuery)
	rqry.Lat = "52.5170365"
	rqry.Lon = "13.3888599"
	rqry.OsmType = "V"
	_, err := rqry.buildQuery()
	if err != nil {
		if !(err.Error() == "OsmType must be 'N', 'W' or 'R'") {
			t.Error("Expecting error about Wrong OSMType. Received" + err.Error())
		}
	} else {
		t.Error("Expecting error about Wrong OSMType. Got none.")
	}

	rqry = new(ReverseQuery)
	rqry.Lat = "52.5170365"
	rqry.Lon = "13.3888599"
	rqry.OsmType = "R"
	_, err = rqry.buildQuery()
	if err != nil {
		t.Error("Expecting no error. Got " + err.Error())
	}
}

func Test_LatLon(t *testing.T) {
	defer SetServer("")
	SetServer("http://nominatim.openstreetmap.org")
	rqry := new(ReverseQuery)
	rqry.Lon = "13.3888599"
	_, err := rqry.buildQuery()
	if err != nil {
		if !(err.Error() == "Cannot search without a latitude. Set field Lat") {
			t.Error("Expecting error about missing latitude. Received" + err.Error())
		}
	} else {
		t.Error("Expecting error about missing latitude. Got none.")
	}

	rqry = new(ReverseQuery)
	rqry.Lat = "52.5170365"
	_, err = rqry.buildQuery()
	if err != nil {
		if !(err.Error() == "Cannot search without a longitude. Set field Lon") {
			t.Error("Expecting error about missing longitude. Received" + err.Error())
		}
	} else {
		t.Error("Expecting error about missing longitude. Got none.")
	}
}

func Test_Zoom(t *testing.T) {
	defer SetServer("")
	SetServer("http://nominatim.openstreetmap.org")
	rqry := new(ReverseQuery)
	rqry.Lon = "13.3888599"
	rqry.Lat = "52.5170365"
	rqry.Zoom = 1337
	_, err := rqry.buildQuery()
	if err != nil {
		if !(err.Error() == "Zoom must be within 0 and 18. 1337 is out of range") {
			t.Error("Expecting error about wrong Zoomfactor. Received" + err.Error())
		}
	} else {
		t.Error("Expecting error about wrong Zoomfactor. Got none.")
	}

	rqry = new(ReverseQuery)
	rqry.Lon = "13.3888599"
	rqry.Lat = "52.5170365"
	rqry.Zoom = 13
	_, err = rqry.buildQuery()
	if err != nil {
		t.Error("Expecting no error. Got " + err.Error())
	}
}
