package vacancy

import (
	"github.com/delphinus35/moesia/util"
)

// Vacancy is a struct to store vacancies of hotels
type Vacancy struct {
	Hotel string
	Dates []*util.Time
}
