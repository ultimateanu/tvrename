package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const findShowUrl string = "http://services.tvrage.com/feeds/search.php?show="

type Results struct {
	Shows []Show `xml:"show"`
}

type Show struct {
	ShowId  int64  `xml:"showid"`
	Name    string `xml:"name"`
	Country string `xml:"country"`
	Started int    `xml:"started"`
	Ended   int    `xml:"ended"`
	Seasons int    `xml:"seasons"`
}

func (s Show) String() string {
	if s.Ended == 0 {
		return fmt.Sprintf("%s [%d - ] (%s)", s.Name, s.Started, s.Country)
	}
	return fmt.Sprintf("%s [%d - %d] (%s)", s.Name, s.Started, s.Ended, s.Country)
}

func findShow(name string) ([]Show, error) {
	res, err := getResource(findShowUrl + url.QueryEscape(name))
	if err != nil {
		return nil, err
	}

	var r Results
	err = xml.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	return r.Shows, nil
}

func getResource(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
