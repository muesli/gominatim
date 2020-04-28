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
	"fmt"
	"strings"
	"testing"
)

func Test_CreateSearchQuery(t *testing.T) {
	defer SetServer("")
	SetServer("https://nominatim.openstreetmap.org")
	expectation := "q=Berlin"
	q := new(SearchQuery)
	q.Q = "Berlin"
	qstr, err := q.buildQuery()
	if !strings.Contains(qstr, expectation) {
		t.Error(fmt.Sprintf("resulting query should contain %s", expectation))
	}
	if err != nil {
		t.Error(fmt.Sprintf("triggered error that was not supposed to: %s", err.Error()))
	}
}

func Test_CreateSearchQueryWithParams(t *testing.T) {
	defer SetServer("")
	SetServer("https://nominatim.openstreetmap.org")
	expectations := []string{
		"city=Berlin",
		"street=Karl-Marx-Allee",
		"county=Berlin",
		"state=Germany",
		"postalcode=012345",
	}
	q := &SearchQuery{
		City:       "Berlin",
		Street:     "Karl-Marx-Allee",
		County:     "Berlin",
		State:      "Germany",
		Postalcode: "012345",
	}
	qstr, err := q.buildQuery()
	for i := range expectations {
		if !strings.Contains(qstr, expectations[i]) {
			t.Error(fmt.Sprintf("resulting query should contain %s", expectations[i]))
		}
	}
	if err != nil {
		t.Error(fmt.Sprintf("triggered error that was not supposed to: %s", err.Error()))
	}
}

func Test_SpecificFieldsUsed(t *testing.T) {
	defer SetServer("")
	SetServer("https://nominatim.openstreetmap.org")
	q1 := &SearchQuery{
		City:       "Berlin",
		Street:     "Karl-Marx-Allee",
		County:     "Berlin",
		State:      "Germany",
		Postalcode: "012345",
	}
	q2 := new(SearchQuery)
	q2.Q = "Berlin"
	if !q1.specificFieldsUsed() {
		t.Error("Q1 -> specific fields are used. should return true")
	}
	if q2.specificFieldsUsed() {
		t.Error("Q2 -> specific fields are not used. should return false")
	}
}

func Test_EmptySearchQuery(t *testing.T) {
	defer SetServer("")
	SetServer("https://nominatim.openstreetmap.org")
	q := new(SearchQuery)
	_, err := q.buildQuery()
	if err == nil {
		t.Error("Empty query should result in an error")
	}
}

func Test_DoubleSearchQuery(t *testing.T) {
	defer SetServer("")
	SetServer("https://nominatim.openstreetmap.org")
	q := &SearchQuery{
		City:       "Berlin",
		Street:     "Karl-Marx-Allee",
		County:     "Berlin",
		State:      "Germany",
		Postalcode: "012345",
		Q:          "Berlin",
	}
	expectations := []string{
		"city=Berlin",
		"street=Karl-Marx-Allee",
		"county=Berlin",
		"state=Germany",
		"postalcode=012345",
	}
	qstr, err := q.buildQuery()
	for i := range expectations {
		if strings.Contains(qstr, expectations[i]) {
			t.Error(fmt.Sprintf("query should not contain %s", expectations[i]))
		}
	}
	if !strings.Contains(qstr, "q=Berlin") {
		t.Error(fmt.Sprintf("query should contain q=Berlin"))
	}
	if err != nil {
		t.Error("should not throw error")
	}
}

func Test_LimitedSearchQuery(t *testing.T) {
	defer SetServer("")
	SetServer("https://nominatim.openstreetmap.org")
	expectation := "limit=123"
	q := new(SearchQuery)
	q.Q = "Berlin"
	q.Limit = 123
	qstr, err := q.buildQuery()
	if !strings.Contains(qstr, expectation) {
		t.Error(fmt.Sprintf("resulting query should contain %s", expectation))
	}
	if err != nil {
		t.Error(fmt.Sprintf("triggered error that was not supposed to: %s", err.Error()))
	}
}

func Test_AddressFields(t *testing.T) {
	defer SetServer("")
	SetServer("https://nominatim.openstreetmap.org")
	q := new(SearchQuery)
	q.Q = "Unter den Linden"
	resp, err := q.Get()
	if resp[0].Address.Suburb != "" {
		t.Error(fmt.Sprintf("Address should contain suburb"))
	}

	if resp[0].Address.State != "" {
		t.Error(fmt.Sprintf("Address should contain State"))
	}

	if err != nil {
		t.Error(fmt.Sprintf("triggered error that was not supposed to: %s", err.Error()))
	}
}
