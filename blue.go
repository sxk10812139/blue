package blue

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	methodRoutes   []MethodRoute
	globalMidwares []Midware
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{Request: r, ResponseWriter: w, engine: e}
	e.handleRequest(c)
}

func (e *Engine) findHandler(c *Context) (HandlerFunc, error) {
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
		DebugLog(err)

		//重定向到404
		return
	}
	//处理全局中间件
	for _, globalMidware := range e.globalMidwares {
		globalMidware(c)
	}
	//处理局部中间件
	handler(c)
}

func (e *Engine) Run(addr string) {

	http.ListenAndServe(addr, e)
}

func (e *Engine) AddRoute(method string, path string, handler HandlerFunc) {
	for index, methodRoute := range e.methodRoutes {
		if methodRoute.Method == method {
			e.methodRoutes[index].RouteNodes = append(e.methodRoutes[index].RouteNodes, RouteNode{path: path, handler: handler})
			return
		}
	}

	routeNode := RouteNode{path: path, handler: handler}
	e.methodRoutes = append(e.methodRoutes, MethodRoute{Method: method, RouteNodes: []RouteNode{routeNode}})
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.AddRoute("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.AddRoute("POST", path, handler)
}

func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.AddRoute("DELETE", path, handler)
}

func (e *Engine) AddGlobalMidware(midware Midware) {
	e.globalMidwares = append(e.globalMidwares, midware)
}

type MethodRoute struct {
	Method     string
	RouteNodes []RouteNode
}

type RouteNode struct {
	path    string
	handler HandlerFunc
}

type Midware HandlerFunc

func NewEngine() *Engine {
	e := &Engine{globalMidwares: []Midware{}, methodRoutes: []MethodRoute{}}

	e.AddGlobalMidware(LogMidware)
	return e
}
