# blue
simple go web framework

## todo list

* Route
    - [x] Powerful route
    - [ ] Route group
    - [ ] Basic midware
* Request
    - [x] Query process
    - [x] Form process
    - [ ] Upload process
    - [ ] Data validataion
    - [ ] Data binding
* Output
    - [x] Json
    - [x] String
    - [ ] Template



## example 
```
func main() {
	e := blue.NewEngine()
	e.GET("/helloworld", HelloWorld)
	e.GET("/json", Json)
	e.GET("/query", Query)
	e.ANY("/user/:name", User)
	e.POST("/upload", Upload)

	e.AddGlobalMidware(CustomGlobalMidware())
	e.Run(":8082")
}

func HelloWorld(c *blue.Context) {
	c.String("helloworld")
}

func CustomGlobalMidware() blue.HandlerFunc {
	return func(c *blue.Context) {
		blue.DebugLog("custom midware before")
		c.Next()
		blue.DebugLog("custom midware after")
	}
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
	a := c.PostDefault("a", "")
	fmt.Printf("%v", a)
	c.String(name)
}

func Upload(c *blue.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		c.String(404, "")
		return
	}
	err = c.SaveUploadedFile(fh, "/tmp/tmp")
	if err != nil {
		c.String(404, "")
		return
	}

	c.String(http.StatusOK, "upload successfully")
}


```

output

    2018/09/28 - 23:45:34  global midware before GET /user/bill
    2018/09/28 - 23:45:34  custom midware before
    2018/09/28 - 23:45:34  custom midware after
    2018/09/28 - 23:45:34  global midware after GET /user/bill