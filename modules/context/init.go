package context

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"safebox.jerson.dev/api/modules/cache"
	"safebox.jerson.dev/api/modules/config"
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
func (r *template) SetUser(id int) {
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

	if config.Vars.Debug {
		log.SetLevel(logrus.DebugLevel)
	}
	r.Logger = log.WithFields(map[string]interface{}{
		"category": r.Category,
		"tag":      tag,
		"token":    r.Token,
		"session":  r.Session,
	})

	return r.Logger
}

// GetCacheResponse ...
func (r template) GetCacheResponse(key string) ([]byte, error) {

	if !config.Vars.Cache.Enable {
		return nil, errors.New("cache disabled")
	}

	client, err := cache.Connect()
	if err != nil {
		return nil, err
	}

	data, err := client.Get(key)
	if err != nil {
		return nil, err
	}

	return data, nil
}
