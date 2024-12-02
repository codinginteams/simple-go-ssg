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

// Run generates a static site using content, templates, and an output directory.
func Run(contentDir, outputDir, templateDir string) error {
	if err := validateDirectories(contentDir, outputDir, templateDir); err != nil {
		return err
	}

	tmpl, err := loadTemplate(templateDir)
	if err != nil {
		return err
	}

	files, err := getMarkdownFiles(contentDir)
	if err != nil {
		return err
	}

	log.Printf("Found %d content files to process...\n", len(files))

	for _, file := range files {
		if err := processMarkdownFile(file, outputDir, tmpl); err != nil {
			return err
		}
	}

	log.Println("Static site generation completed successfully.")
	return nil
}

func validateDirectories(contentDir, outputDir, templateDir string) error {
	if contentDir == "" || outputDir == "" || templateDir == "" {
		return errors.New("contentDir, outputDir, and templateDir cannot be empty")
	}

	return os.MkdirAll(outputDir, 0755)
}

func loadTemplate(templateDir string) (*template.Template, error) {
	baseTemplatePath := filepath.Join(templateDir, "base.html")
	return template.ParseFiles(baseTemplatePath)
}

func getMarkdownFiles(contentDir string) ([]string, error) {
	return filepath.Glob(filepath.Join(contentDir, "*.md"))
}

func processMarkdownFile(file, outputDir string, tmpl *template.Template) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	filename := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
	outputFile := filepath.Join(outputDir, filename+".html")

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
	return nil
}
