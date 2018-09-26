package blue

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	engine         *Engine
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
