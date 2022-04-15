package constants

var consts map[string]string = map[string]string{
	"name": "MyFirstApi",
	"get":  "/Get",
	"post": "/Post",
	"main": "/",
}

func Get(name string) string {
	return consts[name]
}
