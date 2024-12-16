package kringler

import (
	"log"
	"os"
	"strings"
)

func NewFromFile(filePath string) *Kringler {
	kringler := New()
	kringler.ReadKringlersFromFile(filePath)
	return kringler
}

// Reads the Kringlers from a file, seperated by a single newline character (\n)
func (k *Kringler) ReadKringlersFromFile(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	for _, kringler := range strings.Split(string(data), "\n") {
		if kringler == "" {
			continue
		}
		k.AddKringler(kringler)
	}
	return k.Kringlers
}
