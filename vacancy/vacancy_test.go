package vacancy

import (
	"testing"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/delphinus35/moesia/util"
)

func TestString(t *testing.T) {
	v := Vacancy{Hotel: "ほげホテル"}
	for i := range make([]int, 3) {
		date := time.Date(2016, 8, 20+i, 0, 0, 0, 0, util.JST)
		v.Dates = append(v.Dates, &util.Time{Time: date})
	}
	expected := heredoc.Doc(`
		ほげホテル: 8/20 (土)
		ほげホテル: 8/21 (日)
		ほげホテル: 8/22 (月)`)
	if result := v.String(); result != expected {
		t.Errorf("result: %s; expected: %s;", result, expected)
	}
}
