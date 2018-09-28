package blue

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	methodRoutes   []MethodRoute
	globalMidwares []HandlerFunc
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{Request: r, ResponseWriter: w, engine: e, index: -1}
	e.handleRequest(c)
}

func (e *Engine) handleRequest(c *Context) {
	path := c.Request.URL.Path
	method := c.Request.Method

	c.handlers = append(c.handlers, e.globalMidwares...)
	for _, methodRoute := range e.methodRoutes {
		if methodRoute.Method == method {
			handler, p, _ := methodRoute.Root.getValue(path)
			if handler != nil {
				c.Params = p
				c.handlers = append(c.handlers, handler)
				c.Next()
				return
			}
		}
	}

	//404
	e.handle404(c)
}

func (e *Engine) handle404(c *Context) {
	c.Next()
	c.ResponseWriter.WriteHeader(http.StatusNotFound)
	c.ResponseWriter.Write([]byte("404"))
}

func (e *Engine) Run(addr string) {

	DebugLog("start listening " + addr)
	http.ListenAndServe(addr, e)
}

func (e *Engine) AddRoute(method string, path string, handler HandlerFunc) {
	for index, methodRoute := range e.methodRoutes {
		if methodRoute.Method == method {
			e.methodRoutes[index].Root.addRoute(path, handler)
			return
		}
	}

	root := new(node)
	e.methodRoutes = append(e.methodRoutes, MethodRoute{Method: method, Root: root})
	root.addRoute(path, handler)
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

func (e *Engine) AddGlobalMidware(midware HandlerFunc) {
	e.globalMidwares = append(e.globalMidwares, midware)
}

type MethodRoute struct {
	Method string
	Root   *node
}

func NewEngine() *Engine {
	e := &Engine{globalMidwares: make([]HandlerFunc, 0), methodRoutes: []MethodRoute{}}

	e.AddGlobalMidware(LogMidware)
	return e
}
