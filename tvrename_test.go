package main

import "testing"

func TestGetResource(t *testing.T) {
	t.Log("trying to get google robots")
	res, err := getResource("http://www.google.com/robots.txt")
	if err != nil || len(res) == 0 {
		t.Fatal("can't get goog robot")
	}
}
