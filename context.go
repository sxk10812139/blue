package blue

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type C map[string]interface{}

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Params         Params
	engine         *Engine
	index          int
	handlers       []HandlerFunc
	parsed         bool
}

func (c *Context) Json(status int, data interface{}) {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		DebugLog(err)
	}
	c.ResponseWriter.Write(jsonBytes)
	c.ResponseWriter.WriteHeader(status)
}

func (c *Context) String(s string) {
	c.ResponseWriter.Write([]byte(s))
}

func (c *Context) Param(key string) string {
	for _, p := range c.Params {
		if p.Key == key {
			return p.Value
		}
	}
	return ""
}

func (c *Context) Get(key string) string {
	return c.GetDefault(key, "")
}

func (c *Context) Post(key string) string {
	return c.PostDefault(key, "")
}

func (c *Context) GetDefault(key string, defaultValue string) string {
	values := c.GetArray(key)
	if len(values) == 0 {
		return defaultValue
	}

	return values[0]
}

func (c *Context) GetArray(key string) []string {
	//这里每次都要Query()解析一下 稍后可以优化
	if values, ok := c.Request.URL.Query()[key]; ok && len(values) > 0 { //len(values)这里注意下
		return values

	}
	return []string{}
}

func (c *Context) PostDefault(key string, defaultValue string) string {
	values := c.PostArray(key)

	if len(values) == 0 {
		return defaultValue
	}

	return values[0]
}

func (c *Context) PostArray(key string) []string {
	c.parseForm()
	if values, ok := c.Request.PostForm[key]; ok {
		return values
	}
	return []string{}
}

func (c *Context) parseForm() {
	if !c.parsed {
		c.Request.ParseForm()
	}
}

func (c *Context) FormFile(file string) (*multipart.FileHeader, error) {
	_, fh, err := c.Request.FormFile(file)
	return fh, err
}

func (c *Context) SaveUploadedFile(fh *multipart.FileHeader, dst string) error {

	s, err := fh.Open()
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		return err
	}
	return nil
}

/**************flowcontrol************************/
func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}
