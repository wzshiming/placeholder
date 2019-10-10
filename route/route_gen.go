// Code generated; DO NOT EDIT.
// file ./route/route_gen.go

package route

import (
	"fmt"
	http "net/http"
	strconv "strconv"

	mux "github.com/gorilla/mux"
	ui "github.com/wzshiming/openapi/ui"
	redoc "github.com/wzshiming/openapi/ui/redoc"
	swaggereditor "github.com/wzshiming/openapi/ui/swaggereditor"
	swaggerui "github.com/wzshiming/openapi/ui/swaggerui"
	githubComWzshimingPlaceholder "github.com/wzshiming/placeholder"
)

// notFoundHandler Is the not found of handler
func notFoundHandler(w http.ResponseWriter, r *http.Request) {

	err := fmt.Errorf("Not found '%s %s'", r.Method, r.URL.Path)

	http.Error(w, err.Error(), 404)

}

// Router is all routing for package
// generated do not edit.
func Router() http.Handler {
	router := mux.NewRouter()

	// Placeholder Define the method scope
	var _placeholder githubComWzshimingPlaceholder.Placeholder
	RoutePlaceholder(router, &_placeholder)

	router = RouteOpenAPI(router)

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	return router
}

// RoutePlaceholder is routing for Placeholder
func RoutePlaceholder(router *mux.Router, _placeholder *githubComWzshimingPlaceholder.Placeholder, fs ...mux.MiddlewareFunc) *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}

	_routePlaceholder := router.PathPrefix("/placeholder").Subrouter()
	if len(fs) != 0 {
		_routePlaceholder.Use(fs...)
	}

	// Registered routing GET /placeholder/{width}x{height}
	var __operationGetPlaceholderWidthXHeight http.Handler
	__operationGetPlaceholderWidthXHeight = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_operationGetPlaceholderWidthXHeight(_placeholder, w, r)
	})
	_routePlaceholder.Methods("GET").Path("/{width}x{height}").Handler(__operationGetPlaceholderWidthXHeight)

	if router.NotFoundHandler == nil {
		router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	}
	return router
}

// _requestPathHeight Parsing the path for of height
func _requestPathHeight(w http.ResponseWriter, r *http.Request) (_height uint64, err error) {

	var _rawHeight = mux.Vars(r)["height"]

	if __height, err := strconv.ParseUint(_rawHeight, 0, 0); err == nil {
		_height = uint64(__height)
	}

	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	return
}

// _requestQueryText Parsing the query for of text
func _requestQueryText(w http.ResponseWriter, r *http.Request) (_text string, err error) {

	var _rawText = r.URL.Query()["text"]

	if len(_rawText) == 0 {
		return
	}

	__rawText_0 := _rawText[0]
	_text = string(__rawText_0)
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	return
}

// _requestPathWidth Parsing the path for of width
func _requestPathWidth(w http.ResponseWriter, r *http.Request) (_width uint64, err error) {

	var _rawWidth = mux.Vars(r)["width"]

	if __width, err := strconv.ParseUint(_rawWidth, 0, 0); err == nil {
		_width = uint64(__width)
	}

	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	return
}

// _operationGetPlaceholderWidthXHeight Is the route of Simple
func _operationGetPlaceholderWidthXHeight(s *githubComWzshimingPlaceholder.Placeholder, w http.ResponseWriter, r *http.Request) {
	// requests github.com/wzshiming/placeholder Placeholder.Simple.width
	var _width uint64
	// requests github.com/wzshiming/placeholder Placeholder.Simple.height
	var _height uint64
	// requests github.com/wzshiming/placeholder Placeholder.Simple.text
	var _text string
	// responses github.com/wzshiming/placeholder Placeholder.Simple.file
	var _file []byte
	// responses github.com/wzshiming/placeholder Placeholder.Simple.err
	var _err error

	// Parsing width.
	_width, _err = _requestPathWidth(w, r)
	if _err != nil {
		return
	}

	// Parsing height.
	_height, _err = _requestPathHeight(w, r)
	if _err != nil {
		return
	}

	// Parsing text.
	_text, _err = _requestQueryText(w, r)
	if _err != nil {
		return
	}

	// Call github.com/wzshiming/placeholder Placeholder.Simple.
	_file, _err = s.Simple(_width, _height, _text)

	// Response code 200 OK for file.
	if _file != nil {
		var __file []byte
		__file = _file

		w.Header().Set("Content-Type", "image/svg+xml")
		w.WriteHeader(200)
		w.Write(__file)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	var __file []byte
	__file = _file

	w.Header().Set("Content-Type", "image/svg+xml")
	w.WriteHeader(200)
	w.Write(__file)

	return
}

var OpenAPI4YAML = []byte(`openapi: 3.0.1
info:
  title: OpenAPI Demo
  description: Automatically generated
  contact:
    name: wzshiming
    url: https://github.com/wzshiming/gen
  version: 0.0.1
servers:
- url: /
- url: '{scheme}{host}{port}{path}'
  variables:
    host:
      enum:
      - localhost
      default: localhost
    path:
      enum:
      - /
      default: /
    port:
      enum:
      - ""
      default: ""
    scheme:
      enum:
      - http://
      - https://
      default: http://
paths:
  /placeholder/{width}x{height}:
    get:
      tags:
      - Placeholder
      summary: Simple style
      description: |-
        Simple style


        #route:"GET /{width}x{height}"#
      parameters:
      - $ref: '#/components/parameters/width_path'
      - $ref: '#/components/parameters/height_path'
      - $ref: '#/components/parameters/text_query'
      responses:
        "200":
          description: |2-


            #content:"image/svg+xml"#
          content:
            image/svg+xml:
              schema:
                type: string
                format: binary
        "400":
          description: Response code is 400
          content:
            text/plain:
              schema:
                type: string
                format: error
components:
  responses:
    err_body:
      description: Response code is 400
      content:
        text/plain:
          schema:
            type: string
            format: error
    file_body:
      description: |2-


        #content:"image/svg+xml"#
      content:
        image/svg+xml:
          schema:
            type: string
            format: binary
  parameters:
    height_path:
      name: height
      in: path
      required: true
      explode: false
      schema:
        type: integer
        format: uint64
    text_query:
      name: text
      in: query
      explode: false
      schema:
        type: string
    width_path:
      name: width
      in: path
      required: true
      explode: false
      schema:
        type: integer
        format: uint64
tags:
- name: Placeholder
  description: |-
    Placeholder Generate SVG placeholders


    #path:"/placeholder"#
`)
var OpenAPI4JSON = []byte(`{"openapi":"3.0.1","info":{"title":"OpenAPI Demo","description":"Automatically generated","contact":{"name":"wzshiming","url":"https://github.com/wzshiming/gen"},"version":"0.0.1"},"servers":[{"url":"/"},{"url":"{scheme}{host}{port}{path}","variables":{"host":{"enum":["localhost"],"default":"localhost"},"path":{"enum":["/"],"default":"/"},"port":{"enum":[""],"default":""},"scheme":{"enum":["http://","https://"],"default":"http://"}}}],"paths":{"/placeholder/{width}x{height}":{"get":{"tags":["Placeholder"],"summary":"Simple style","description":"Simple style\n\n\n#route:\"GET /{width}x{height}\"#","parameters":[{"$ref":"#/components/parameters/width_path"},{"$ref":"#/components/parameters/height_path"},{"$ref":"#/components/parameters/text_query"}],"responses":{"200":{"description":"\n\n#content:\"image/svg+xml\"#","content":{"image/svg+xml":{"schema":{"type":"string","format":"binary"}}}},"400":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}}}}}},"components":{"responses":{"err_body":{"description":"Response code is 400","content":{"text/plain":{"schema":{"type":"string","format":"error"}}}},"file_body":{"description":"\n\n#content:\"image/svg+xml\"#","content":{"image/svg+xml":{"schema":{"type":"string","format":"binary"}}}}},"parameters":{"height_path":{"name":"height","in":"path","required":true,"explode":false,"schema":{"type":"integer","format":"uint64"}},"text_query":{"name":"text","in":"query","explode":false,"schema":{"type":"string"}},"width_path":{"name":"width","in":"path","required":true,"explode":false,"schema":{"type":"integer","format":"uint64"}}}},"tags":[{"name":"Placeholder","description":"Placeholder Generate SVG placeholders\n\n\n#path:\"/placeholder\"#"}]}`)

// RouteOpenAPI
func RouteOpenAPI(router *mux.Router) *mux.Router {
	openapi := map[string][]byte{
		"openapi.json": OpenAPI4JSON,
		"openapi.yml":  OpenAPI4YAML,
		"openapi.yaml": OpenAPI4YAML,
	}
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger", ui.HandleWithFiles(openapi, swaggerui.Asset)))
	router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui", ui.HandleWithFiles(openapi, swaggerui.Asset)))
	router.PathPrefix("/swaggereditor/").Handler(http.StripPrefix("/swaggereditor", ui.HandleWithFiles(openapi, swaggereditor.Asset)))
	router.PathPrefix("/redoc/").Handler(http.StripPrefix("/redoc", ui.HandleWithFiles(openapi, redoc.Asset)))
	return router
}
