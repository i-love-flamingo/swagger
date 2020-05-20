package swagger

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -fs -prefix "swagger-ui-dist/" -o swagger-ui-dist.go -pkg swagger swagger-ui-dist/

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"log"
	"net/http"
)

type Module struct {
}

func (s *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

type routes struct {}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.HandleAny("swag.api", web.WrapHTTPHandler(http.StripPrefix("/api-console", assetHandler())))
	//registry.HandleAny("swag.api", web.WrapHTTPHandler(assetHandler()))
	registry.Route("/api-console/*n", "swag.api")
	registry.Route("/api-console", "swag.api")

	registry.HandleAny("swag.api.swagger", web.WrapHTTPHandler(swaggerDefHandler()))
	registry.Route("/api/swagger.json", "swag.api.swagger")
}

func swaggerDefHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		http.ServeFile(rw, req, "docs/swagger.json")
	})
}

func assetHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println(req)
		http.FileServer(AssetFile()).ServeHTTP(rw, req)
	})
}
