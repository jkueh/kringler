package kringler

import "log"

// CreateAssignments creates the assignments for the kringlers, by assigning each kringler to the next kringler in the
// list.
//
// Calling this function before the list has been shuffled will result in a predictable list of assignments.
func (k *Kringler) CreateLinearAssignments() []KringlerAssignment {
	kringlersLength := len(k.Kringlers)
	assignments := make([]KringlerAssignment, kringlersLength)
	for index, kringler := range k.Kringlers {
		target := kringler
		if index == kringlersLength-1 {
			target = k.Kringlers[0]
		} else {
			target = k.Kringlers[index+1]
		}
		if kringler == target {
			log.Fatalln("Unable to determine target for kringler:", kringler)
		}
		k.Assignments = append(k.Assignments, KringlerAssignment{Giver: kringler, Receiver: target})
	}
	copy(assignments, k.Assignments)
	return assignments
}
