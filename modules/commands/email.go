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

var tmplHTML = template.New("html")
var tmplText = template.New("text")

func init() {
	var err error
	tmplHTML, err = tmplHTML.Parse(`
<h1>{{.Name}}</h1>

<p>Hello {{.Username}}, we send you your last location:</p>

<ul>
<li>latitude: {{.Latitude}}</li>
<li>longitude: {{.Longitude}}</li>
<li>date: {{.Date}}</li>
<li>image: https://static-maps.yandex.ru/1.x/?lang=en-US&ll={{.Latitude}},{{.Longitude}}&z=13&l=map&size=600,300&pt={{.Latitude}},{{.Longitude}},vkbkm</li>
</ul>

<br/>
<img width="100%" style="width:100%,min-height:400px" src="https://static-maps.yandex.ru/1.x/?lang=en-US&ll={{.Latitude}},{{.Longitude}}&z=13&l=map&size=600,300&pt={{.Latitude}},{{.Longitude}},vkbkm"  alt="map"/>
`)
	if err != nil {
		panic(err)
	}

	tmplText, err = tmplText.Parse(`

Hello {{.Username}}, we send you your last location:

latitude: {{.Latitude}}
longitude: {{.Longitude}}
date: {{.Date}}


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

	log.Info("build template")

	params := map[string]interface{}{
		"Name":      config.Vars.Name,
		"Username":  user.Username,
		"Latitude":  location.Latitude,
		"Longitude": location.Longitude,
		"Date":      location.Date.Format(time.RFC850),
	}

	bufHTML := &bytes.Buffer{}
	err = tmplHTML.Execute(bufHTML, params)
	if err != nil {
		return err
	}
	htmlContent := bufHTML.String()

	bufText := &bytes.Buffer{}
	err = tmplText.Execute(bufText, params)
	if err != nil {
		return err
	}
	plainContent := bufText.String()

	log.Info("prepare email")
	from := mail.NewEmail(config.Vars.Name, config.Vars.SendGrid.From)
	to := mail.NewEmail(user.Username, user.Email)

	subject := fmt.Sprintf("Last location used in %s", config.Vars.Name)
	message := mail.NewSingleEmail(from, subject, to, plainContent, htmlContent)

	log.Info("send email")
	client := sendgrid.NewSendClient(config.Vars.SendGrid.APIKey)
	response, err := client.Send(message)

	if err != nil {
		return err
	}
	log.Infof("body: %s", response.Body)

	return nil

}
