package pii

var str string = "default value"

func Init(s string) {
	str = s
}

func Any(any interface{}) (string, error) {
	return str, nil
}
