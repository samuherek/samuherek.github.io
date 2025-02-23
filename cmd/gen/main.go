package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"samuherek/pages"

	cp "github.com/otiai10/copy"
)

func hashCss(outputDir string) string {
	originalPath := path.Join(outputDir, "static/css/main.css")
	f, err := os.Open(originalPath)
	if err != nil {
		log.Fatal("Failed to open main.css")
	}
	defer f.Close()

	hash := sha256.New()
	if _, err = io.Copy(hash, f); err != nil {
		log.Fatal("Failed to read file to hash")
	}

	hashStr := fmt.Sprintf("%x", hash.Sum(nil))[:10]
	hashName := fmt.Sprintf("main-%s.css", hashStr)
	nextPath := path.Join(outputDir, fmt.Sprintf("static/css/%s", hashName))
	if err = os.Rename(originalPath, nextPath); err != nil {
		log.Fatal("Failed to rename main.css with hash")
	}

	return hashName
}

func processStaticFiles(outputDir string) string {
	staticDir := "static"

	fmt.Println("Processing static files")

	if err := cp.Copy(staticDir, path.Join(outputDir, staticDir)); err != nil {
		log.Fatal("ERROR: Failed copying static content")
	}

	return hashCss(outputDir)
}

func processIndex(outputDir, hashCss string) {
	fmt.Println("Process index")

	f, err := os.Create(path.Join(outputDir, "index.html"))
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = pages.HomeWrapper("Home", hashCss, pages.Home()).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func process404(outputDir, hashCss string) {
	fmt.Println("Process 404")

	f, err := os.Create(path.Join(outputDir, "404.html"))
	if err != nil {
		log.Fatal("Failed to create output file: %v", err)
	}

	err = pages.Error404Wrapper("404", hashCss, pages.Error404()).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func prepareBuildDir(outputDir string) {
	fmt.Println("Prepare build directory")

	_, err := os.Stat(outputDir)
	if os.IsExist(err) {
		if err := os.RemoveAll(outputDir); err != nil {
			log.Fatal("ERROR: Failed to remove %s directory", outputDir)
		}
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatal("ERROR: Failed to create %s directory", outputDir)
	}

	fmt.Println("Build directory is ready.")
}

func main() {
	outputDir := "build"

	prepareBuildDir(outputDir)
	hashCss := processStaticFiles(outputDir)
	processIndex(outputDir, hashCss)
	process404(outputDir, hashCss)

	fmt.Println("Done generating static site.")
}
