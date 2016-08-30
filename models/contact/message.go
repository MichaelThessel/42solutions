package fortytwo

import (
	"github.com/astaxie/beego/validation"
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
