package blue

func LogMidware(c *Context) {
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	method := c.Request.Method

	DebugLog("global midware before", method, path, raw)
	c.Next()
	DebugLog("global midware after", method, path, raw)
}
