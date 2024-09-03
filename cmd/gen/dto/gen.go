package dto

import (
	"bytes"
	"fmt"
	"log"
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

	// --------------------------------------------------
	// GENERATE DTO LIST FILE
	// --------------------------------------------------
	err = generateList(config.Data, rootDir)
	if err != nil {
		log.Fatalf("Error generating DTO list file: %v\n", err)
	}
	fmt.Printf("1 DTO list file generated\n")

	return nil
}

func generateList(pkgs []dto.Package, root string) error {
	tmpl, err := template.ParseFiles(filepath.Join(root, "cmd", "gen", "dto", "tpl", "dto_list.go.tpl"))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, pkgs)
	if err != nil {
		return err
	}
	filename := filepath.Join(root, "domain", "dto", "dto_list.go")
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(filename, buf.Bytes(), 0644)
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
