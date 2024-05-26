package pii

var str string = "default value"

func Init(s string) {
	str = s
}

func Any(any interface{}) (string, error) {
	println("any...")
	return str, nil
}
