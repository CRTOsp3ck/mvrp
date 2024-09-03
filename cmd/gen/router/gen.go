package router

import (
	"bytes"
	"fmt"
	"mvrp/config/handlers"
	"mvrp/util"
	"os"
	"path/filepath"
	"text/template"
)

func Generate() error {
	rootDir, err := util.Util.FS.FindProjectRoot("go.mod")
	if err != nil {
		return err
	}

	config, err := handlers.GetConfig()
	if err != nil {
		return err
	}

	tmplPath := filepath.Join(rootDir, "cmd", "gen", "router", "tpl", "router.go.tpl")
	tmplName := filepath.Base(tmplPath)
	tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
		"ToPascalCase": util.Util.NC.ToPascalCase,
	}).ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	// Create the output file
	outputDir := filepath.Join(rootDir, "http", "router")
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	outputFile := filepath.Join(outputDir, "router.go")
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, config); err != nil {
		return err
	}

	if err := os.WriteFile(outputFile, buf.Bytes(), os.ModePerm); err != nil {
		return err
	}

	fmt.Printf("1 Router file generated with %d routes\n", getRoutesCount(config))
	return nil
}

func getRoutesCount(config *handlers.Root) int {
	count := 0
	for _, pkg := range config.Data {
		for _, handlerGroup := range pkg.Handlers {
			count += len(handlerGroup.Routes)
		}
	}
	return count
}
