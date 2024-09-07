package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"mvrp/config/handlers"
	"mvrp/util"
	"os"
	"path/filepath"
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

	count := 0
	for _, pkg := range config.Data {
		for _, handlerGroup := range pkg.Handlers {
			name := util.Util.NC.ToPascalCase(handlerGroup.Name)
			templateData := map[string]interface{}{
				"Package": pkg.Package,
				"Name":    name,
				"Routes":  handlerGroup.Routes,
				"IsView":  handlerGroup.IsView,
			}

			// Load and parse the template
			tmplPath := filepath.Join(rootDir, "cmd", "gen", "handler", "tpl", "handler.go.tpl")
			tmplName := filepath.Base(tmplPath)
			tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
				"ToPascalCase": util.Util.NC.ToPascalCase,
			}).ParseFiles(tmplPath)
			if err != nil {
				return err
			}

			// Create the output file
			outputDir := filepath.Join(rootDir, "http", "handler", pkg.Package)
			if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
				return err
			}

			outputFile := filepath.Join(outputDir, fmt.Sprintf("handler_%s.go", handlerGroup.Name))
			var buf bytes.Buffer
			if err := tmpl.Execute(&buf, templateData); err != nil {
				return err
			}

			if err := os.WriteFile(outputFile, buf.Bytes(), os.ModePerm); err != nil {
				return err
			}
			count++
		}
	}

	fmt.Printf("%d Handler files generated\n", count)

	return nil
}
