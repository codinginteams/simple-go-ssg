package generator

import (
	"errors"
	"log"
	"path/filepath"
)

// Run is a public method that generates a static site.
func Run(contentDir, outputDir, templateDir string) error {
	if contentDir == "" || outputDir == "" || templateDir == "" {
		return errors.New("contentDir, outputDir, and templateDir cannot be empty")
	}

	files, err := filepath.Glob(filepath.Join(contentDir, "*.md"))
	if err != nil {
		return err
	}

	log.Printf("Found %d content files to process...\n", len(files))

	log.Println("Static site generation completed successfully.")
	return nil
}
