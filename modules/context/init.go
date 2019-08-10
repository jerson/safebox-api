package context

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// template ...
type template struct {
	Token    string
	Session  string
	Category string
	Logger   *logrus.Entry
}

// GetToken ...
func (r *template) GetToken() string {
	return r.Token
}

// SetSession ...
func (r *template) SetUser(id int64) {
	r.Session = fmt.Sprint(id)
}

// GetLogger ...
func (r *template) GetLogger(tag string) *logrus.Entry {
	if r.Logger != nil {
		return r.Logger.WithFields(map[string]interface{}{
			"tag":     tag,
			"session": r.Session,
		})
	}

	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)

	r.Logger = log.WithFields(map[string]interface{}{
		"category": r.Category,
		"tag":      tag,
		"token":    r.Token,
		"session":  r.Session,
	})

	return r.Logger
}
