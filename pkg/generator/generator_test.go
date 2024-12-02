package generator

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func createTempDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "generator_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	return dir
}

func createTempFile(t *testing.T, dir, filename, content string) string {
	path := filepath.Join(dir, filename)
	if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create temp file %s: %v", filename, err)
	}
	return path
}

func cleanup(t *testing.T, path string) {
	if err := os.RemoveAll(path); err != nil {
		t.Fatalf("Failed to clean up: %v", err)
	}
}

func TestRun(t *testing.T) {
	contentDir := createTempDir(t)
	defer cleanup(t, contentDir)

	outputDir := createTempDir(t)
	defer cleanup(t, outputDir)

	templateDir := createTempDir(t)
	defer cleanup(t, templateDir)

	createTempFile(t, templateDir, "base.html", "<html>{{.Content}}</html>")
	createTempFile(t, contentDir, "index.md", "Hello, World!")

	err := Run(contentDir, outputDir, templateDir)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	htmlFile := filepath.Join(outputDir, "index.html")
	if _, err := os.Stat(htmlFile); os.IsNotExist(err) {
		t.Errorf("Expected file %s to be generated, but it was not", htmlFile)
	}
}

func TestValidateDirectories(t *testing.T) {
	outputDir := createTempDir(t)
	defer cleanup(t, outputDir)

	err := validateDirectories("content", outputDir, "templates")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	err = validateDirectories("", outputDir, "templates")
	if err == nil {
		t.Error("Expected an error for empty contentDir, got nil")
	}
}

func TestLoadTemplate(t *testing.T) {
	templateDir := createTempDir(t)
	defer cleanup(t, templateDir)

	templatePath := createTempFile(t, templateDir, "base.html", "<html>{{.Content}}</html>")

	tmpl, err := loadTemplate(templateDir)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if tmpl == nil {
		t.Error("Expected template, got nil")
	}

	os.Remove(templatePath)
	_, err = loadTemplate(templateDir)
	if err == nil {
		t.Error("Expected an error for missing template, got nil")
	}
}

func TestGetMarkdownFiles(t *testing.T) {
	contentDir := createTempDir(t)
	defer cleanup(t, contentDir)

	createTempFile(t, contentDir, "test1.md", "Content 1")
	createTempFile(t, contentDir, "test2.md", "Content 2")

	files, err := getMarkdownFiles(contentDir)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("Expected 2 files, got: %d", len(files))
	}
}

func TestProcessMarkdownFile(t *testing.T) {
	outputDir := createTempDir(t)
	defer cleanup(t, outputDir)

	templateDir := createTempDir(t)
	defer cleanup(t, templateDir)

	createTempFile(t, templateDir, "base.html", "<html>{{.Content}}</html>")
	tmpl, err := loadTemplate(templateDir)
	if err != nil {
		t.Fatalf("Failed to load template: %v", err)
	}

	markdownFile := createTempFile(t, outputDir, "test.md", "Hello, World!")

	err = processMarkdownFile(markdownFile, outputDir, tmpl)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	htmlFile := filepath.Join(outputDir, "test.html")
	if _, err := os.Stat(htmlFile); os.IsNotExist(err) {
		t.Errorf("Expected file %s to be generated, but it was not", htmlFile)
	}
}
