package kringler

import (
	"testing"
)

// Ensure that givers appear once and only once
func TestUniqueGivers(t *testing.T) {
	assignments := NewFromList(KringlersList).CreateLinearAssignments()
	for index, assignment := range assignments {
		for innerIndex, innerAssignment := range assignments {
			if index == innerIndex { // Skip comparing to self
				continue
			}
			if assignment.Giver == innerAssignment.Giver {
				t.Error("Giver appears more than once:", assignment.Giver)
			}
		}
		t.Log(assignment.Giver, "\t->\t", assignment.Receiver)
	}
}

// Ensure that receivers appear once and only once
func TestUniqueReceivers(t *testing.T) {
	assignments := NewFromList(KringlersList).CreateLinearAssignments()
	for index, assignment := range assignments {
		for innerIndex, innerAssignment := range assignments {
			if index == innerIndex { // Skip comparing to self
				continue
			}
			if assignment.Receiver == innerAssignment.Receiver {
				t.Error("Receiver appears more than once:", assignment.Receiver)
			}
		}
		t.Log(assignment.Giver, "\t->\t", assignment.Receiver)
	}
}

// Ensure that no one gets themselves as their Kris Kringle™
func TestDifferentReceiver(t *testing.T) {
	assignments := NewFromList(KringlersList).CreateLinearAssignments()
	for _, assignment := range assignments {
		if assignment.Giver == assignment.Receiver {
			t.Error("Giver is also a receiver:", assignment.Giver)
		}
		t.Log(assignment.Giver, "\t->\t", assignment.Receiver)
	}
}
