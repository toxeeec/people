// Package people provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package people

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /login)
	PostLogin(ctx echo.Context) error

	// (GET /me/followers)
	GetMeFollowers(ctx echo.Context, params GetMeFollowersParams) error

	// (GET /me/followers/{handle})
	GetMeFollowersHandle(ctx echo.Context, handle HandleParam) error

	// (GET /me/following)
	GetMeFollowing(ctx echo.Context, params GetMeFollowingParams) error

	// (DELETE /me/following/{handle})
	DeleteMeFollowingHandle(ctx echo.Context, handle HandleParam) error

	// (GET /me/following/{handle})
	GetMeFollowingHandle(ctx echo.Context, handle HandleParam) error

	// (PUT /me/following/{handle})
	PutMeFollowingHandle(ctx echo.Context, handle HandleParam) error

	// (POST /posts)
	PostPosts(ctx echo.Context) error

	// (POST /refresh)
	PostRefresh(ctx echo.Context) error

	// (POST /register)
	PostRegister(ctx echo.Context) error

	// (GET /users/{handle}/followers)
	GetUsersHandleFollowers(ctx echo.Context, handle HandleParam, params GetUsersHandleFollowersParams) error

	// (GET /users/{handle}/following)
	GetUsersHandleFollowing(ctx echo.Context, handle HandleParam, params GetUsersHandleFollowingParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostLogin converts echo context to params.
func (w *ServerInterfaceWrapper) PostLogin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostLogin(ctx)
	return err
}

// GetMeFollowers converts echo context to params.
func (w *ServerInterfaceWrapper) GetMeFollowers(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetMeFollowersParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMeFollowers(ctx, params)
	return err
}

// GetMeFollowersHandle converts echo context to params.
func (w *ServerInterfaceWrapper) GetMeFollowersHandle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "handle" -------------
	var handle HandleParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "handle", runtime.ParamLocationPath, ctx.Param("handle"), &handle)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter handle: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMeFollowersHandle(ctx, handle)
	return err
}

// GetMeFollowing converts echo context to params.
func (w *ServerInterfaceWrapper) GetMeFollowing(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetMeFollowingParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMeFollowing(ctx, params)
	return err
}

// DeleteMeFollowingHandle converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteMeFollowingHandle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "handle" -------------
	var handle HandleParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "handle", runtime.ParamLocationPath, ctx.Param("handle"), &handle)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter handle: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteMeFollowingHandle(ctx, handle)
	return err
}

// GetMeFollowingHandle converts echo context to params.
func (w *ServerInterfaceWrapper) GetMeFollowingHandle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "handle" -------------
	var handle HandleParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "handle", runtime.ParamLocationPath, ctx.Param("handle"), &handle)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter handle: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMeFollowingHandle(ctx, handle)
	return err
}

// PutMeFollowingHandle converts echo context to params.
func (w *ServerInterfaceWrapper) PutMeFollowingHandle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "handle" -------------
	var handle HandleParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "handle", runtime.ParamLocationPath, ctx.Param("handle"), &handle)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter handle: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutMeFollowingHandle(ctx, handle)
	return err
}

// PostPosts converts echo context to params.
func (w *ServerInterfaceWrapper) PostPosts(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostPosts(ctx)
	return err
}

// PostRefresh converts echo context to params.
func (w *ServerInterfaceWrapper) PostRefresh(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRefresh(ctx)
	return err
}

// PostRegister converts echo context to params.
func (w *ServerInterfaceWrapper) PostRegister(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRegister(ctx)
	return err
}

// GetUsersHandleFollowers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersHandleFollowers(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "handle" -------------
	var handle HandleParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "handle", runtime.ParamLocationPath, ctx.Param("handle"), &handle)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter handle: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersHandleFollowersParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUsersHandleFollowers(ctx, handle, params)
	return err
}

// GetUsersHandleFollowing converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersHandleFollowing(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "handle" -------------
	var handle HandleParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "handle", runtime.ParamLocationPath, ctx.Param("handle"), &handle)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter handle: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersHandleFollowingParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUsersHandleFollowing(ctx, handle, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/login", wrapper.PostLogin)
	router.GET(baseURL+"/me/followers", wrapper.GetMeFollowers)
	router.GET(baseURL+"/me/followers/:handle", wrapper.GetMeFollowersHandle)
	router.GET(baseURL+"/me/following", wrapper.GetMeFollowing)
	router.DELETE(baseURL+"/me/following/:handle", wrapper.DeleteMeFollowingHandle)
	router.GET(baseURL+"/me/following/:handle", wrapper.GetMeFollowingHandle)
	router.PUT(baseURL+"/me/following/:handle", wrapper.PutMeFollowingHandle)
	router.POST(baseURL+"/posts", wrapper.PostPosts)
	router.POST(baseURL+"/refresh", wrapper.PostRefresh)
	router.POST(baseURL+"/register", wrapper.PostRegister)
	router.GET(baseURL+"/users/:handle/followers", wrapper.GetUsersHandleFollowers)
	router.GET(baseURL+"/users/:handle/following", wrapper.GetUsersHandleFollowing)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xZX2/bNhD/Kga3Rzmy0xbY/Jasy5at64IsxR4Co2Cks8RFIlWSWuMZ+u7DkZQp2Yot",
	"e3aCtXkJIurI3/F+91dekEjkheDAtSKTBSmopDlokOYppTzO4ArX8JFxMiEF1SkJCKc5kImTIAGR8Klk",
	"EmIy0bKEgKgohZzirpw+vAOe6JRMxm8CkjNeP74JiJ4XeIzSkvGEVFVAMpYz3Yb8VIKce0wjQVYgWF7m",
	"ZDIejQyCe1qez7iGBCQJyMMwEUO3WjKuDWZBE9gIiQJtxB0xKmshUPpcxAyMdc9KnX5QIM9FPMfnSHAN",
	"XOO/tCgyFlHNBA//UoLjmgf/VsKMTMg3oecutG9VWB/qMT0rVUCuhNFgN7xCigKkdmo3tjWoPf1u1OJ2",
	"vMotmkXQgg0jEUMCfAgPWtKhpok5dUbvUXihAE+PoFq9wO0SeNp9tRtxD1wd1Jj2yE5TmhVVCK6sVc5p",
	"fG0JPhj6j1IKx2MMKpKswEPIBLEGNVgVkAsh71gcAz8+sjHIgKlBTrOZkDnEqMB78YNHPa4C78WgxjLA",
	"+kKUPD4+7jUoUcoIBlzowcxgVgH5wGmpUyHZP/AEOpxFESg10EsOmFJ11nSbm2llPXJdst4tJ2+J2/iu",
	"WQXqKM5Aa5B8Mh5VRr2UqhSl9zlapeYIFjcOWMm1Lktfvu1Kvn1QSgXyI4v9DdQ9K1xpUOqzkAYcPZ5q",
	"Uw7cYjvlnXbZzmtz5Xf1SoU1iq2oK39mNFMQjE878uSSjaWa0yog1q3WfCIHpbC4rZGzemotOHVVZGNh",
	"2J3nenPQWQkCEkmgGuIz3SIiphqGmuV4202mR5GTGyvXQxeL9ZHqDoc4qh8WQuluPyxdRG/KHK7ut3kz",
	"p3nrejtOlzVznUpqUo152+EZCDCToNLHBFZUaEkjbHd6moksE59d67m9qepjT3/kukXtO9T3wHDW/Vbh",
	"fPLdO78+Wxp8LMG0zLu8+7QvlLkQ+oJZZhpy1c/Fl8FOpaRzWwAhKiXT8z9Q0PrTHVAJEsuhf7qoM8cv",
	"f97UDT0eZN/6LJJq7S7O+EwYkzOdmSwOoshgcHZ1SQLyN0hli/P4ZHQyQs1EAZwWjEzIK7MUmInJKBRm",
	"ImEmXgqXQNH9TY9wGePRQul3RqQ5Lcwfs0lroAhb08Rqd3o6Gh2lKW43KL//isCvLVa3xk6nsNEumy3j",
	"7VtazRaCW3+6JbiMQ0FAwhzCVhZJoMPMP4H+DS4artscem+71fAioR8Wq2CrcGOaraZHJMVG0QZOdjYw",
	"bnq1fZMfPppRaOzYjL/bKV6/ZixvJ4917sKFzTFVTxJ/9i3PLlQ2v2900PN6+/X96PNUVsYdvRRzo9HB",
	"aHGVchsbtqa9hNTzhZSpwWvctUIqhgw0rPP41qw3qPzKAuv740/x6FY4vdsPCSbW4v9EdNAnJF8S5IET",
	"ZG37ouxq5MoX2/ez/RMHHc0k0Hh+mMDDDIuNvNrc0V8ZkT06+uW3+mN28+ZTzjP38sesi5YgS5b7DrGZ",
	"rmsntAdhjV8g/ncD2M7275i+JCRMafdpZ4OBndRXNeR2W6xUzWGn3+xq+lJbU/YfYlvFJfhCG/Rdu4Ca",
	"IMNKxyjUydaWsWiNrX3moxe2erPlSrN/v6h/wzdhhwZxzzk0n/zualr9GwAA///XZS1tFiEAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
