# blue
go web framework

## example 
```
func main() {
	e := blue.NewEngine()
	e.AddRoute("GET", "/helloworld", HelloWorld)
	e.AddRoute("GET", "/json", Json)
	e.AddRoute("GET", "/user/:name", User)

	e.AddGlobalMidware(CustomGlobalMidware)
	e.Run(":8080")
}

func CustomGlobalMidware(c *blue.Context) {
	blue.DebugLog("自定义全局中间件 请求前")
	c.Next()
	blue.DebugLog("自定义全局中间件 请求后")
}

func HelloWorld(c *blue.Context) {
	c.String("helloworld")
}

func Json(c *blue.Context) {
	var a interface{}
	a = struct {
		A string
		B string
	}{
		A: "nihao",
		B: "ssd",
	}
	c.Json(a)
}

func User(c *blue.Context) {
	name := c.Params.ByName("name")
	c.String(name)
}

```