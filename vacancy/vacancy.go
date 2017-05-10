package vacancy

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/delphinus/moesia/util"
)

// Vacancy is a struct to store vacancies of hotels
type Vacancy struct {
	Hotel string
	Dates []*util.Time
}

func (v *Vacancy) String() (str string) {
	var strs []string
	for _, date := range v.Dates {
		strs = append(strs, fmt.Sprintf("%s: %s", v.Hotel, date.MoesiaFormat()))
	}
	str = strings.Join(strs, "\n")
	return
}

// Vacancies is an array of Vacancy
type Vacancies struct {
	List []Vacancy
}

// TemplateName specified a template filename for mail body
var TemplateName = "templates/mailBody.tmpl"

// MailBody returns body string for mail
func (vs *Vacancies) MailBody() (html string, err error) {
	tmplString := string(MustAsset(TemplateName))
	tmpl := template.Must(template.New(TemplateName).Parse(tmplString))
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, vs)
	if err != nil {
		err = fmt.Errorf("failed to execute template: %v", err)
		return
	}
	html = buf.String()
	return
}
