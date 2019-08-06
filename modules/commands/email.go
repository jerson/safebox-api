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
)

var tmpl = template.New("email")

func init() {
	var err error
	tmpl, err = tmpl.Parse(`
<h1>{{.name}}</h1>

<p>Hello {{.user.Username}}, we send you your last location:</p>

latitude: {{.location.Latitude}}
longitude: {{.location.Longitude}}


<img width="100%" src="https://static-maps.yandex.ru/1.x/?lang=en-US&ll={{.location.Latitude}},{{.location.Longitude}}&z=13&l=map&size=600,300&pt={{.location.Latitude}},{{.location.Longitude}},vkbkm"  alt="map"/>
`)
	if err != nil {
		panic(err)
	}
}

// EmailLocation ...
func EmailLocation(ctx context.Context, userID int64) error {

	locationRepo := repositories.NewUserLocationRepository(ctx)
	repo := repositories.NewUserRepository(ctx)
	user, err := repo.FindOneByID(userID)
	if err != nil {
		return err
	}
	if user.Email == "" {
		return errors.New("empty email")
	}

	location, err := locationRepo.FindOneByUserID(user.ID)
	if err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, map[string]interface{}{
		"name":     config.Vars.Name,
		"user":     user,
		"location": location,
	})
	if err != nil {
		return err
	}

	htmlContent := buf.String()

	from := mail.NewEmail(config.Vars.Name, config.Vars.SendGrid.From)
	to := mail.NewEmail(user.Username, user.Email)

	subject := fmt.Sprintf("Last location used in %s", config.Vars.Name)
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)

	client := sendgrid.NewSendClient(config.Vars.SendGrid.APIKey)
	_, err = client.Send(message)

	if err != nil {
		return err
	}

	return nil

}
