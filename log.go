package blue

func LogMidware(c *Context) {
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	method := c.Request.Method

	DebugLog("全局中间件 请求前", method, path, raw)
	c.Next()
	DebugLog("全局中间件 请求后", method, path, raw)
}
