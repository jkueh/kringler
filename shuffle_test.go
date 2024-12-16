package kringler

import (
	"reflect"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	kringler := NewFromList(KringlersList)

	// Note to self: Setting this too high will increase the likelihood of a false positive
	//
	// Setting it too low will also increase the likelihood of a false positive
	numberOfShuffles := 5

	shuffledSlices := make([][]string, numberOfShuffles)

	for index := range numberOfShuffles {
		time.Sleep(1 * time.Nanosecond)
		shuffledSlices[index] = kringler.ShuffleKringlers()
	}

	// Go through the shuffled slices and make sure each one is unique
	for index, slice := range shuffledSlices {
		t.Log("Shuffle:\t", index, "\t", slice)
		for innerIndex, innerSlice := range shuffledSlices {
			if index == innerIndex {
				continue
			}
			if reflect.DeepEqual(slice, innerSlice) {
				t.Error("ShuffleKringlers() is potentially not shuffling the kringlers list")
			}
		}
	}
}
