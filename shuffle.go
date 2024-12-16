package kringler

import (
	"time"

	"golang.org/x/exp/rand"
)

// ShuffleKringlers shuffles the kringlers list, and returns a copy of the now-shuffled list.
//
// https://stackoverflow.com/a/12267471
func (k *Kringler) ShuffleKringlers() []string {
	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	returnSlice := make([]string, len(k.Kringlers))
	for i := range k.Kringlers {
		ri := r.Intn(len(k.Kringlers))
		k.Kringlers[i], k.Kringlers[ri] = k.Kringlers[ri], k.Kringlers[i]
	}
	copy(returnSlice, k.Kringlers)
	return returnSlice
}
