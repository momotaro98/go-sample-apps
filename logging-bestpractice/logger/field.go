package logger

import (
	"github.com/sirupsen/logrus"
)

// Fields is ...
type Fields []Field

// Field is ...
type Field struct {
	key   string
	value interface{}
}

// F is ...
func F(key string, value interface{}) Field {
	return Field{
		key:   key,
		value: value,
	}
}

// E is ...
func E(err error) Field {
	return Field{
		key:   logrus.ErrorKey,
		value: err,
	}
}
