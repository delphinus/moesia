package vacancy

import (
	"testing"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/delphinus/moesia/util"
)

func sampleVacancy(hotel string) (v Vacancy) {
	v = Vacancy{Hotel: hotel}
	for i := 0; i < 3; i++ {
		date := time.Date(2016, 8, 20+i, 0, 0, 0, 0, util.JST)
		v.Dates = append(v.Dates, &util.Time{Time: date})
	}
	return
}

func sampleVacancies(hotels ...string) (vs Vacancies) {
	for _, hotel := range hotels {
		vs.List = append(vs.List, sampleVacancy(hotel))
	}
	return
}

func TestVacancyString(t *testing.T) {
	v := sampleVacancy("ほげホテル")
	expected := heredoc.Doc(`
		ほげホテル: 8/20 (土)
		ほげホテル: 8/21 (日)
		ほげホテル: 8/22 (月)`)
	if result := v.String(); result != expected {
		t.Errorf("result: %s; expected: %s;", result, expected)
	}
}

func TestVacanciesMailBody(t *testing.T) {
	vs := sampleVacancies("ほげホテル", "ふがホテル")
	if _, err := vs.MailBody(); err != nil {
		t.Errorf("failed to create mail body: %v", err)
	}
}
