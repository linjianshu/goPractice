package main

import "testing"

func TestS1(t *testing.T) {
	if S1("abcdefgh") != 9 {
		t.Error(`S1("abcdefgh")!=9`)
	}

	if S1("") != 0 {
		t.Error(`S1("")!=0`)
	}
}
