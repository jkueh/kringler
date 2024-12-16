package kringler

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
)

// Generate messages for all givers using the specified template.
func (k *Kringler) CreateMessagesForGivers(
	templateFilePath string, // The path to the template file
	outputDir string, // The directory where the output files will be saved
) {
	for _, assignment := range k.Assignments {
		k.CreateMessageWithTemplate(assignment, templateFilePath, outputDir)
	}
}

// Writes a message to the giver using the specified template.
func (k *Kringler) CreateMessageWithTemplate(
	assignment KringlerAssignment, // The assignment to use in the template
	templateFilePath string, // The path to the template file
	outputDir string, // The directory where the output file will be saved
) {
	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		log.Fatalln("An error occurred while reading the template file:", err)
	}

	file, err := os.Create(
		path.Join(outputDir, fmt.Sprintf("%s.txt", assignment.Giver)),
	)

	err = tmpl.Execute(file, assignment)
	if err != nil {
		log.Fatalln("An error occurred while writing the message:", err)
	}

	file.Close()
}
