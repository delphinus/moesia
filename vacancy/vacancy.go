package vacancy

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/delphinus35/moesia/util"
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

// TemplateFilename specified a template filename for mail body
var TemplateFilename = "../templates/mailBody.tmpl"

var funcMap = template.FuncMap{}

// MailBody returns body string for mail
func (vs *Vacancies) MailBody() (html string, err error) {
	tmpl := template.Must(template.New("mailBody.tmpl").Funcs(funcMap).ParseFiles(TemplateFilename))
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string][]Vacancy{
		"list": vs.List,
	})
	if err != nil {
		err = fmt.Errorf("failed to execute template: %v", err)
		return
	}
	html = buf.String()
	return
}
