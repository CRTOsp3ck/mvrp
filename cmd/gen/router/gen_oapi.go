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
	reflector.Spec.Servers = []openapi3.Server{
		{
			URL:         "http://localhost:6900",
			Description: util.Util.Ptr.StrPtr("MVRP Development server"),
		},
	}
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
			// Request structures
			reqMap := make(map[string]interface{})
			searchDto := dto.FindDTO(fmt.Sprintf("Search%sDTO", util.Util.NC.ToPascalCase(getCleanName(handler.Name))))
			reqMap["search"] = searchDto
			createDto := dto.FindDTO(fmt.Sprintf("Create%sDTO", util.Util.NC.ToPascalCase(getCleanName(handler.Name))))
			reqMap["create"] = createDto
			updateDto := dto.FindDTO(fmt.Sprintf("Update%sDTO", util.Util.NC.ToPascalCase(getCleanName(handler.Name))))
			reqMap["update"] = updateDto

			// Response structures
			// TODO: Nest the appropriate model inside the htresp.Response struct
			// respMap := make(map[string]interface{})
			// respMap["search"] = &htresp.Response{
			// 	Data: searchDto,
			// }

			err = addMainOperations(&reflector, pkg, handler, reqMap)
			if err != nil {
				return err
			}
		}
	}

	// Ext route operations
	err = addExtOperations(&reflector)
	if err != nil {
		return err
	}

	// Enum route operations
	err = addEnumOperations(&reflector)
	if err != nil {
		return err
	}

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

func addMainOperations(reflector *openapi3.Reflector, pkg handlers.Package, handler handlers.Handler, reqMap map[string]interface{}) error {
	type req struct {
		ID int `path:"id"`
	}

	// Add List operation
	listOp, err := reflector.NewOperationContext(http.MethodGet, fmt.Sprintf("/v1/main/%s/%s", pkg.Package, handler.Name))
	if err != nil {
		return err
	}
	listOp.AddRespStructure(new(htresp.Response), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusOK })
	reflector.AddOperation(listOp)

	// Add Search operation
	searchOp, err := reflector.NewOperationContext(http.MethodPost, fmt.Sprintf("/v1/main/%s/%s/search", pkg.Package, handler.Name))
	if err != nil {
		return err
	}
	searchOp.AddReqStructure(reqMap["search"])
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

	if !handler.IsView {
		// Add Create operation
		createOp, err := reflector.NewOperationContext(http.MethodPost, fmt.Sprintf("/v1/main/%s/%s", pkg.Package, handler.Name))
		if err != nil {
			return err
		}
		createOp.AddReqStructure(reqMap["create"])
		createOp.AddRespStructure(new(htresp.Response), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusCreated })
		reflector.AddOperation(createOp)

		// Add Update operation
		updateOp, err := reflector.NewOperationContext(http.MethodPut, fmt.Sprintf("/v1/main/%s/%s/{id}", pkg.Package, handler.Name))
		if err != nil {
			return err
		}
		updateOp.AddReqStructure(new(req))
		updateOp.AddReqStructure(reqMap["update"])
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
	}

	return nil
}

func addEnumOperations(reflector *openapi3.Reflector) error {
	// Add Get operation
	getOp, err := reflector.NewOperationContext(http.MethodGet, "/v1/enum")
	if err != nil {
		return err
	}
	getOp.AddRespStructure(new(dto.EnumsDTO), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusOK })
	reflector.AddOperation(getOp)

	return nil
}

func addExtOperations(reflector *openapi3.Reflector) error {
	type req struct {
		ID int `path:"id"`
	}

	// OPERATION - /v1/ext/inventory/inventory/exists_by_item_id/{id}
	getOp, err := reflector.NewOperationContext(http.MethodGet, "/v1/ext/inventory/inventory/exists_by_item_id/{id}")
	if err != nil {
		return err
	}
	getOp.AddReqStructure(new(req))
	getOp.AddRespStructure(new(htresp.Response), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusOK })
	reflector.AddOperation(getOp)

	// OPERATION - /v1/ext/inventory/inventory_transaction/search_all
	postOp, err := reflector.NewOperationContext(http.MethodPost, "/v1/ext/inventory/inventory_transaction/search_all")
	if err != nil {
		return err
	}
	postOp.AddReqStructure(new(dto.SearchInventoryTransactionDTO))
	postOp.AddRespStructure(new(htresp.Response), func(cu *openapi.ContentUnit) { cu.HTTPStatus = http.StatusOK })
	reflector.AddOperation(postOp)

	return nil
}

func getCleanName(name string) string {
	// remove last 5 characters if they are "_view"
	if len(name) > 5 && name[len(name)-5:] == "_view" {
		return name[:len(name)-5]
	}
	return name
}
