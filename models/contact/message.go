package fortytwo

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"github.com/astaxie/beego/validation"
	"net/textproto"
)

type Message struct {
	Name      string
	Email     string
	Telephone string
	Message   string
}

func (m *Message) Validate() (bool, map[string][]string) {
	v := validation.Validation{}

	v.Required(m.Name, "name")
	v.Required(m.Message, "message")
	v.Required(m.Email, "email")

	v.Email(m.Email, "email")

	v.MaxSize(m.Name, 100, "name")
	v.MaxSize(m.Email, 100, "email")
	v.MaxSize(m.Telephone, 100, "telephone")
	v.MaxSize(m.Message, 1000, "message")

	errorList := make(map[string][]string)
	for _, e := range v.Errors {
		errorList[e.Key] = append(errorList[e.Key], e.Message)
	}

	return !v.HasErrors(), errorList
}

func (m *Message) Send() bool {
	e := new(utils.Email)

	e.Headers = textproto.MIMEHeader{}
	e.Username = beego.AppConfig.String("smtp::username")
	e.Password = beego.AppConfig.String("smtp::password")
	e.Host = beego.AppConfig.String("smtp::host")
	e.Port, _ = beego.AppConfig.Int("smtp::port")

	e.From = beego.AppConfig.String("contact::from")
	e.To = []string{beego.AppConfig.String("contact::to")}

	e.Subject = "42Solutions Contact Request"
	e.Text = fmt.Sprintf("Name: %s\nEmail: %s\nTelephone: %s\n\n%s", m.Name, m.Email, m.Telephone, m.Message)

	err := e.Send()

	if err != nil {
		// TODO: Logging
		fmt.Println(err.Error())
		return false
	}

	return true
}
