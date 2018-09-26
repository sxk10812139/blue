package blue

func LogMidware(c *Context) {
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	method := c.Request.Method

	DebugLog("全局中间件", method, path, raw)
}
