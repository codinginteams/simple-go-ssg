package generator

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Run is a public method that generates a static site.
func Run(contentDir, outputDir, templateDir string) error {
	// Validate inputs
	if contentDir == "" || outputDir == "" || templateDir == "" {
		return errors.New("contentDir, outputDir, and templateDir cannot be empty")
	}

	// Ensure the output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Load the base HTML template
	baseTemplatePath := filepath.Join(templateDir, "base.html")
	tmpl, err := template.ParseFiles(baseTemplatePath)
	if err != nil {
		return err
	}

	// Process all Markdown files in the content directory
	files, err := filepath.Glob(filepath.Join(contentDir, "*.md"))
	if err != nil {
		return err
	}

	log.Printf("Found %d content files to process...\n", len(files))

	for _, file := range files {
		// Read the Markdown file
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		// Extract the filename (without extension) for the output HTML
		filename := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		outputFile := filepath.Join(outputDir, filename+".html")

		// Render the template with content
		data := map[string]string{
			"Title":   strings.Title(filename),
			"Content": string(content),
		}
		output, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		defer output.Close()

		if err := tmpl.Execute(output, data); err != nil {
			return err
		}

		log.Printf("Generated %s\n", outputFile)
	}

	log.Println("Static site generation completed successfully.")
	return nil
}
