package router

import (
	"fmt"
	"mvrp/config/handlers"
	"mvrp/domain/dto"
	"mvrp/htresp"
	"mvrp/util"
	"net/http"
	"os"
	"path/filepath"

	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
)

func generateOpenAPISpec(rootDir string) error {
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}
	reflector.Spec.Info.
		WithTitle("MVRP API").
		WithVersion("1.0.0").
		WithDescription("API documentation for MVRP").
		WithContact(openapi3.Contact{
			Name:  util.Util.Ptr.StrPtr("SP3CK"),
			Email: util.Util.Ptr.StrPtr("regedits@gmail.com"),
		})

	// Main route operations
	config, err := handlers.GetConfig()
	if err != nil {
		return err
	}
	for _, pkg := range config.Data {
		for _, handler := range pkg.Handlers {
			dtoMap := make(map[string]interface{})

			searchDto := dto.FindDTO(fmt.Sprintf("Search%sDTO", util.Util.NC.ToPascalCase(handler.Name)))
			dtoMap["search"] = searchDto

			createDto := dto.FindDTO(fmt.Sprintf("Create%sDTO", util.Util.NC.ToPascalCase(handler.Name)))
			if err != nil {
				return err
			}
			dtoMap["create"] = createDto

			updateDto := dto.FindDTO(fmt.Sprintf("Update%sDTO", util.Util.NC.ToPascalCase(handler.Name)))
			dtoMap["update"] = updateDto

			err = addMainOperations(&reflector, pkg, handler, dtoMap)
			if err != nil {
				return err
			}
		}
	}

	// Enum route operations
	// err := addEnumOperations(&reflector)
	// if err != nil {
	// 	return err
	// }

	// Serialize and write the specification
	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		return err
	}

	filename := filepath.Join(rootDir, "http", "router", "openapi.yaml")
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	err = os.WriteFile(filename, schema, 0644)
	if err != nil {
		return err
	}

	return nil
}

func addMainOperations(reflector *openapi3.Reflector, pkg handlers.Package, handler handlers.Handler, dtoMap map[string]interface{}) error {
	type req struct {
		ID int `path:"id"`
	}

	// Add Search operation
	searchOp, err := reflector.NewOperationContext(http.MethodPost, fmt.Sprintf("/v1/main/%s/%s/search", pkg.Package, handler.Name))
	if err != nil {
		return err
	}
	searchOp.AddReqStructure(dtoMap["search"])
	searchOp.AddRespStructure(new(htresp.Response), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusOK })
	reflector.AddOperation(searchOp)

	// Add Get operation
	getOp, err := reflector.NewOperationContext(http.MethodGet, fmt.Sprintf("/v1/main/%s/%s/{id}", pkg.Package, handler.Name))
	if err != nil {
		return err
	}
	getOp.AddReqStructure(new(req))
	getOp.AddRespStructure(new(htresp.Response), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusOK })
	reflector.AddOperation(getOp)

	// Add Create operation
	createOp, err := reflector.NewOperationContext(http.MethodPost, fmt.Sprintf("/v1/main/%s/%s", pkg.Package, handler.Name))
	if err != nil {
		return err
	}
	createOp.AddReqStructure(dtoMap["create"])
	createOp.AddRespStructure(new(htresp.Response), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusCreated })
	reflector.AddOperation(createOp)

	// Add Update operation
	updateOp, err := reflector.NewOperationContext(http.MethodPut, fmt.Sprintf("/v1/main/%s/%s/{id}", pkg.Package, handler.Name))
	if err != nil {
		return err
	}
	updateOp.AddReqStructure(new(req))
	updateOp.AddReqStructure(dtoMap["update"])
	updateOp.AddRespStructure(new(htresp.Response), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusOK })
	reflector.AddOperation(updateOp)

	// Add Delete operation
	deleteOp, err := reflector.NewOperationContext(http.MethodDelete, fmt.Sprintf("/v1/main/%s/%s/{id}", pkg.Package, handler.Name))
	if err != nil {
		return err
	}
	deleteOp.AddReqStructure(new(req))
	deleteOp.AddRespStructure(new(htresp.Response), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusNoContent })
	reflector.AddOperation(deleteOp)

	// Add List operation
	listOp, err := reflector.NewOperationContext(http.MethodGet, fmt.Sprintf("/v1/main/%s/%s", pkg.Package, handler.Name))
	if err != nil {
		return err
	}
	listOp.AddRespStructure(new(htresp.Response), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusOK })
	reflector.AddOperation(listOp)
	return nil
}

// func addEnumOperations(reflector *openapi3.Reflector) error {
// 	// Add Get operation
// 	getOp, err := reflector.NewOperationContext(http.MethodGet, "/v1/enum")
// 	if err != nil {
// 		return err
// 	}
// 	getOp.AddRespStructure(new(dto.EnumDTO), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusOK })
// 	reflector.AddOperation(getOp)

// 	return nil
// }
