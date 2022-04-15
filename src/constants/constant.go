package constants

var consts map[string]string = map[string]string{
	"token": "someToken",
}

func Get(name string) string {
	return consts[name]
}
