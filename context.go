package blue

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Params         Params
	engine         *Engine
	index          int
	handlers       []HandlerFunc
}

func (c *Context) Json(data interface{}) {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		DebugLog(err)
	}
	c.ResponseWriter.Write(jsonBytes)
}

func (c *Context) String(s string) {
	c.ResponseWriter.Write([]byte(s))
}

func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}
