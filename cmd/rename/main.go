package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// getCurrentModuleName reads `go.mod` and extracts the current module name.
func getCurrentModuleName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", fmt.Errorf("module name not found in go.mod")
}

// replaceInFiles replaces occurrences of oldModule with newModule in all project files.
func replaceInFiles(rootDir, oldModule, newModule string) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore directories like .git, binary files, and the rename command itself
		if info.IsDir() || strings.Contains(path, ".git") || strings.HasSuffix(path, "/cmd/rename") {
			return nil
		}

		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Replace all occurrences of the old module name with the new one
		newContent := strings.ReplaceAll(string(content), oldModule, newModule)

		// Write changes only if modifications were made
		if string(content) != newContent {
			err = os.WriteFile(path, []byte(newContent), info.Mode())
			if err != nil {
				return err
			}
			fmt.Println("Updated:", path)
		}

		return nil
	})
}

// runGoModTidy runs "go mod tidy" to clean up dependencies.
func runGoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run cmd/rename/main.go <new_project_name>")
		os.Exit(1)
	}

	newModule := os.Args[1]

	// Get the current module name from go.mod
	oldModule, err := getCurrentModuleName()
	if err != nil {
		fmt.Println("Error retrieving current module name:", err)
		os.Exit(1)
	}

	fmt.Println("Renaming project:", oldModule, "→", newModule)

	// Replace module name in all files
	err = replaceInFiles(".", oldModule, newModule)
	if err != nil {
		fmt.Println("Error replacing package names:", err)
		os.Exit(1)
	}

	// Update `go.mod` module line
	err = replaceInFiles(".", "module "+oldModule, "module "+newModule)
	if err != nil {
		fmt.Println("Error updating go.mod:", err)
		os.Exit(1)
	}

	// Run `go mod tidy`
	err = runGoModTidy()
	if err != nil {
		fmt.Println("Error running go mod tidy:", err)
		os.Exit(1)
	}

	fmt.Println("Project renamed successfully to:", newModule)
}

