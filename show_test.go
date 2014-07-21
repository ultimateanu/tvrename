package main

import "testing"

var tests = []struct {
	query           string
	topShowResponse Show
}{
	{"house", Show{3908, "House", "US", 2004, 2012, 8}},
	{"avatar", Show{2680, "Avatar: The Last Airbender", "US", 2005, 2008, 3}},
}

func TestFindShow(t *testing.T) {
	for _, c := range tests {
		shows, err := findShow(c.query)
		if err != nil || len(shows) == 0 {
			t.Fatalf("No results for '%s'", c.query)
		}
		if shows[0] != c.topShowResponse {
			t.Fatalf("For query '%s'\nExp: %#v\nGot: %#v",
				c.query, c.topShowResponse, shows[0])
		}
	}
}
