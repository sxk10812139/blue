package blue

import (
	"net/http"
)

type Engine struct {
	methodRoutes   []MethodRoute
	globalMidwares []Midware
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{Request: r, ResponseWriter: w, engine: e}
	e.handleRequest(c)
}

func (e *Engine) findHandler(c *Context) (func(c *Context), error) {
	path := c.Request.URL.Path
	method := c.Request.Method
	for _, methodRoute := range e.methodRoutes {
		if methodRoute.Method == method {
			for _, routeNode := range methodRoute.RouteNodes {
				if routeNode.path == path {
					return routeNode.handler, nil
				}
			}
		}
	}

	return func(c *Context) {
		c.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.ResponseWriter.Write([]byte("404"))
	}, nil
}

func (e *Engine) handleRequest(c *Context) {
	handler, err := e.findHandler(c)
	if err != nil {
		debugLog(err)
	}
	//处理全局中间件
	for _, globalMidware := range e.globalMidwares {
		globalMidware(c)
	}
	//处理局部中间件
	handler(c)
}

func (e *Engine) Run() {
	http.ListenAndServe(":8080", e)
}

func (e *Engine) AddRoute(method string, path string, handler func(c *Context)) {
	for index, methodRoute := range e.methodRoutes {
		if methodRoute.Method == method {
			e.methodRoutes[index].RouteNodes = append(e.methodRoutes[index].RouteNodes, RouteNode{path: path, handler: handler})
			return
		}
	}

	routeNode := RouteNode{path: path, handler: handler}
	e.methodRoutes = append(e.methodRoutes, MethodRoute{Method: method, RouteNodes: []RouteNode{routeNode}})
}

func (e *Engine) GET(path string, handler func(c *Context)) {
	e.AddRoute("GET", path, handler)
}

func (e *Engine) POST(path string, handler func(c *Context)) {
	e.AddRoute("POST", path, handler)
}

func (e *Engine) DELETE(path string, handler func(c *Context)) {
	e.AddRoute("DELETE", path, handler)
}

func (e *Engine) AddGlobalMidware(midware func(c *Context)) {
	e.globalMidwares = append(e.globalMidwares, midware)
}

type MethodRoute struct {
	Method     string
	RouteNodes []RouteNode
}

type RouteNode struct {
	path    string
	handler func(c *Context)
}

type Midware func(c *Context)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	engine         *Engine
}

func (c *Context) Json() {

}

func debugLog(e interface{}) {

}

func NewEngine() *Engine {
	return &Engine{globalMidwares: []Midware{}, methodRoutes: []MethodRoute{}}
}

func Run() {
	e := NewEngine()

	e.Run()
}
