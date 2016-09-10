package vacancy

import (
	"fmt"
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

func (vs *Vacancies) String() (str string) {
	for _, v := range vs.List {
		str += fmt.Sprintln(v.String())
	}
	return
}
