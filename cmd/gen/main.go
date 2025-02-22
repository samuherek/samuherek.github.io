package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	cp "github.com/otiai10/copy"
	"samuherek/pages"
)

func main() {
	outputDir := "build"
	staticDir := "static"
	// pagesDir := "pages"

	_, err := os.Stat(outputDir)
	if os.IsExist(err) {
		if err := os.RemoveAll(outputDir); err != nil {
			log.Fatal("ERROR: Failed to remove %s directory", outputDir)
		}
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatal("ERROR: Failed to create %s directory", outputDir)
	}
	fmt.Println("Build directory exists")

	if err = cp.Copy(staticDir, path.Join(outputDir, staticDir)); err != nil {
		log.Fatal("ERROR: Failed copying static content")
	}

	f, err := os.Create(path.Join(outputDir, "index.html"))
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = pages.HomeWrapper("Home", pages.Home()).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}

	fmt.Println("Done generating static site")
	//
	// // Generate static files
	// pages := map[string]string{
	// 	"index.html": templates.Index(), // Assuming you have an `Index()` function from Templ
	// }
	//
	// for name, content := range pages {
	// 	path := filepath.Join(outputDir, name)
	// 	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
	// 		fmt.Println("Error writing file:", err)
	// 		return
	// 	}
	// 	fmt.Println("Generated:", path)
	// }
	//
	// fmt.Println("Static site generated successfully!")
}
