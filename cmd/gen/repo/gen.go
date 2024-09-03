package repo

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"mvrp/config/dto"
	"mvrp/config/models"
	"mvrp/errors"
	"mvrp/util"
)

type TemplateData struct {
	Package              string
	ModelName            string
	PluralModelName      string
	HasSearchDTO         bool
	GroupQueryFields     []dto.Field
	SearchQueryStatement string
}

func Generate() error {
	rootDir, err := util.Util.FS.FindProjectRoot("go.mod")
	if err != nil {
		return errors.WrapError(errors.ErrTypeFileNotFound, fmt.Sprintf("Failed to find project root: %v", err))
	}

	config, err := models.GetConfig()
	if err != nil {
		return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to load config: %v", err))
	}

	// Load the template for generating the model files
	tmplModel, err := loadModelTemplate(rootDir)
	if err != nil {
		return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to load model template: %v", err))
	}

	// Load the template for generating the view files
	tmplView, err := loadViewTemplate(rootDir)
	if err != nil {
		return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to load view template: %v", err))
	}

	// Iterate over each package in the configuration
	count := 0
	for _, pkg := range config.Data {
		data := TemplateData{
			Package: pkg.Package,
		}
		err = generateRepoConstructor(data, rootDir)
		if err != nil {
			return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to generate repository constructor: %v", err))
		}

		for _, model := range pkg.Models {
			hasSearchDTO := checkForSearchDTO(model.Name)
			data = TemplateData{
				Package:         pkg.Package,
				ModelName:       model.Name,
				PluralModelName: model.PluralName,
				HasSearchDTO:    hasSearchDTO,
			}
			if hasSearchDTO {
				sdto, err := dto.GetDTO("Search" + model.Name + "DTO")
				if err != nil {
					return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to generate search DTO: %v", err))
				}
				data.GroupQueryFields = sdto.GroupQueryFields
				sqfs, err := dto.GetSearchQueryFields(sdto.Name)
				if err != nil {
					return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to get search query fields: %v", err))
				}
				data.SearchQueryStatement = createSearchQuery(sqfs)
			}
			err = generateRepo(data, rootDir, tmplModel)
			if err != nil {
				return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to generate repo: %v", err))
			}
			count++
		}

		// Generate repo files for views
		for _, view := range pkg.Views {
			hasSearchDTO := checkForSearchDTO(view.Name)
			data = TemplateData{
				Package:         pkg.Package,
				ModelName:       view.Name,
				PluralModelName: view.PluralName,
				HasSearchDTO:    hasSearchDTO,
			}
			err = generateRepo(data, rootDir, tmplView)
			if err != nil {
				return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to generate repo: %v", err))
			}
			count++
		}

	}
	fmt.Printf("%d Repo files generated\n", count)
	return nil
}

func generateRepo(data TemplateData, root string, tmpl *template.Template) error {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to execute template: %v", err))
	}
	outputPath := filepath.Join(
		root, "data", "repo", data.Package, fmt.Sprintf("repo_%s.go", util.Util.NC.ToSnakeCase(data.ModelName)))
	err = os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to create directories for output: %v", err))
	}
	err = os.WriteFile(outputPath, buf.Bytes(), 0644)
	if err != nil {
		return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to write file %s: %v", outputPath, err))
	}
	return nil
}

func generateRepoConstructor(data TemplateData, root string) error {
	tmplFilePath := filepath.Join(root, "cmd", "gen", "repo", "tpl", "repo.go.tpl")
	tmplName := filepath.Base(tmplFilePath)
	tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
		"ToPascalCase": util.Util.NC.ToPascalCase,
	}).ParseFiles(tmplFilePath)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return err
	}
	filename := filepath.Join(root, "data", "repo", data.Package, "repo.go")
	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return errors.WrapError(errors.ErrTypeInternal, fmt.Sprintf("Failed to create directories for output: %v", err))
	}
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

func loadModelTemplate(root string) (*template.Template, error) {
	tmplFilePath := filepath.Join(root, "cmd", "gen", "repo", "tpl", "repo_model.go.tpl")
	tmplName := filepath.Base(tmplFilePath)
	return template.New(tmplName).Funcs(template.FuncMap{
		"ToPascalCase": util.Util.NC.ToPascalCase,
	}).ParseFiles(tmplFilePath)
}

func loadViewTemplate(root string) (*template.Template, error) {
	tmplFilePath := filepath.Join(root, "cmd", "gen", "repo", "tpl", "repo_view.go.tpl")
	tmplName := filepath.Base(tmplFilePath)
	return template.New(tmplName).Funcs(template.FuncMap{
		"ToPascalCase": util.Util.NC.ToPascalCase,
	}).ParseFiles(tmplFilePath)
}

func checkForSearchDTO(modelName string) bool {
	return dto.IfDtoExists(fmt.Sprintf("Search%sDTO", modelName))
}

func createSearchQuery(searchQueryFields []string) string {
	var sqs string
	if len(searchQueryFields) > 0 {
		// Step 1: Initialize a slice to hold the conditions
		var conditions []string
		// Step 2: Iterate over the sqfds slice and append the condition for each Name field
		for _, sqf := range searchQueryFields {
			conditions = append(conditions, fmt.Sprintf("%s ILIKE ?", sqf))
		}
		// Step 3: Join the conditions with " or "
		conditionStr := strings.Join(conditions, " or ")
		// Step 4: Create a slice to hold the parameters
		params := make([]interface{}, len(searchQueryFields))
		for i := range params {
			params[i] = "dto.Keyword"
		}
		// Step 5: Combine the conditions and parameters into the And clause
		andClause := fmt.Sprintf("And(\n\t\t\t\"%s\",\n", conditionStr)
		for _, param := range params {
			percent := "\"%\""
			andClause += fmt.Sprintf("\t\t\t%s + %s + %s,\n", percent, param, percent)
		}
		andClause += "\t\t)"
		sqs = andClause
	}
	return sqs
}
