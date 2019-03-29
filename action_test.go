package rolele_test

import (
	"github.com/tanopwan/rolele"
	"testing"
)

var allowActions = []rolele.Action{{ID: 1, Name: "action_1"}, {ID: 2, Name: "action_2"}}

func TestApplyActionsSuccess(t *testing.T) {

	if rolele.ApplyActions(allowActions, []string{"action_1"}) != true {
		t.Fail()
	}

	if rolele.ApplyActions(allowActions, []string{"action_2"}) != true {
		t.Fail()
	}

	if rolele.ApplyActions(allowActions, []string{"action_3", "action_2"}) != true {
		t.Fail()
	}
}

func TestApplyActionsFail(t *testing.T) {
	if rolele.ApplyActions(allowActions, []string{"action_3"}) != false {
		t.Fail()
	}
}
