package dto

import (
	"fmt"
	"mvrp/config/dto"
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

	config, err := dto.GetConfig()
	if err != nil {
		return err
	}

	// Load the template file
	tmplFilePath := filepath.Join(rootDir, "cmd", "gen", "dto", "tpl", "dto.go.tpl")
	tmplName := filepath.Base(tmplFilePath)
	tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
		"ToPascalCase":    util.Util.NC.ToPascalCase,
		"ToSnakeCase":     util.Util.NC.ToSnakeCase,
		"IsSearchDto":     isSearchDto,
		"IsObjectTypeDto": isObjectTypeDto,
	}).ParseFiles(tmplFilePath)
	if err != nil {
		return err
	}

	// Generate a separate file for each package
	for _, pkg := range config.Data {
		// Deduplicate imports before passing data to the template
		imports := deduplicateImports(pkg)

		fileName := fmt.Sprintf("dto_%s.go", pkg.Package)
		filePath := filepath.Join(rootDir, "domain", "dto", fileName)

		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Execute the template for each package
		err = tmpl.Execute(file, struct {
			Package string
			DTOs    []dto.DTO
			Imports []string
		}{
			Package: pkg.Package,
			DTOs:    pkg.DTOs,
			Imports: imports,
		})
		if err != nil {
			return err
		}
	}

	fmt.Printf("%d DTO files generated\n", len(config.Data))
	return nil
}

func deduplicateImports(pkg dto.Package) []string {
	importSet := make(map[string]bool)
	var imports []string

	for _, dto := range pkg.DTOs {
		for _, inherit := range dto.Inherits {
			importPath := fmt.Sprintf("mvrp/data/model/%s", inherit.Package)
			if !importSet[importPath] {
				importSet[importPath] = true
				imports = append(imports, importPath)
			}
		}
	}

	return imports
}

func isObjectTypeDto(modelName string) bool {
	// Check if the last 3 characters of the model name are "DTO"
	return modelName[len(modelName)-3:] == "DTO"
}

func isSearchDto(modelName string) bool {
	// Check if the first 6 characters of the model name are "Search"
	pascalCaseModelName := util.Util.NC.ToPascalCase(modelName)
	return pascalCaseModelName[:6] == "Search"
}
