package context

import (
	"github.com/sirupsen/logrus"
)

// Context ...
type Context interface {
	Set(name string, value interface{})
	Get(name string) interface{}
	Close()
	GetToken() string
	SetUser(id int64)
	GetLogger(tag string) *logrus.Entry
}
