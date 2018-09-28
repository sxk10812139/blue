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

func HelloWorld(c *blue.Context) {
	c.String("helloworld")
}

func CustomGlobalMidware(c *blue.Context) {
	blue.DebugLog("custom midware before")
	c.Next()
	blue.DebugLog("custom midware after")
}

func Json(c *blue.Context) {
	var a interface{}
	a = struct {
		A string  
		B string
	}{
		A: "a",
		B: "b",
	}
	c.Json(a)
}

func User(c *blue.Context) {
	name := c.Params.ByName("name")
	c.String(name)
}


```