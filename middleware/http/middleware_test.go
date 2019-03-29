package http

import "testing"

func TestGetActionFromHeader(t *testing.T) {
	if getActionFromHeader("all_po")[0] != "all_po" {
		t.Errorf("all_po failed")
	}

	if getActionFromHeader("all_po;")[0] != "all_po" {
		t.Errorf("all_po; failed")
	}

	if getActionFromHeader(" all_po;")[0] != "all_po" {
		t.Errorf(" all_po; failed")
	}

	actions := getActionFromHeader(" all_po ; all_so")
	if actions[0] != "all_po" || actions[1] != "all_so" {
		t.Errorf(" all_po ; all_so failed")
	}
}
