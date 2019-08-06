package commands

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"safebox.jerson.dev/api/modules/config"
	"safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
	"time"
)

var tmpl = template.New("email")

func init() {
	var err error
	tmpl, err = tmpl.Parse(`
<h1>{{.Name}}</h1>

<p>Hello {{.Username}}, we send you your last location:</p>

latitude: {{.Latitude}}
longitude: {{.Longitude}}
longitude: {{.Date}}


<img width="100%" src="https://static-maps.yandex.ru/1.x/?lang=en-US&ll={{.Latitude}},{{.Longitude}}&z=13&l=map&size=600,300&pt={{.Latitude}},{{.Longitude}},vkbkm"  alt="map"/>
`)
	if err != nil {
		panic(err)
	}
}

// EmailLocation ...
func EmailLocation(ctx context.Context, userID int64) error {

	log := ctx.GetLogger("EmailLocation")
	locationRepo := repositories.NewUserLocationRepository(ctx)
	repo := repositories.NewUserRepository(ctx)
	user, err := repo.FindOneByID(userID)
	if err != nil {
		return err
	}
	if user.Email == "" {
		return errors.New("empty email")
	}

	log.Info("find location")
	location, err := locationRepo.FindOneByUserID(user.ID)
	if err != nil {
		return err
	}
	buf := &bytes.Buffer{}

	log.Info("build template")
	err = tmpl.Execute(buf, map[string]interface{}{
		"Name":      config.Vars.Name,
		"Usename":   user.Username,
		"Latitude":  location.Latitude,
		"Longitude": location.Longitude,
		"Date":      location.Date.Format(time.RFC850),
	})
	if err != nil {
		return err
	}

	htmlContent := buf.String()

	log.Info("prepare email")
	from := mail.NewEmail(config.Vars.Name, config.Vars.SendGrid.From)
	to := mail.NewEmail(user.Username, user.Email)

	subject := fmt.Sprintf("Last location used in %s", config.Vars.Name)
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)

	log.Info("send email")
	client := sendgrid.NewSendClient(config.Vars.SendGrid.APIKey)
	response, err := client.Send(message)

	if err != nil {
		return err
	}
	log.Infof("body: %s", response.Body)

	return nil

}
