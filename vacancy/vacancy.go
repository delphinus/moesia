package vacancy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/delphinus35/moesia/util"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// Vacancy is a struct to store vacancies of hotels
type Vacancy struct {
	Hotel string
	Dates []*util.Time
}

// Vacancies is an array of Vacancy
type Vacancies struct {
	List []Vacancy
}

func (v *Vacancy) String() (str string) {
	var strs []string
	for _, date := range v.Dates {
		strs = append(strs, fmt.Sprintf("%s: %s", v.Hotel, date.MoesiaFormat()))
	}
	str = strings.Join(strs, "\n")
	return
}

func (vs *Vacancies) String() (str string) {
	var strs []string
	for _, vacancy := range vs.List {
		strs = append(strs, vacancy.String())
	}
	str = strings.Join(strs, "\n")
	return
}

// TemplateName specified a template filename for mail body
var TemplateName = "templates/mailBody.tmpl"

// TemplateData is data struct for templates
type TemplateData struct {
	Vacancies *Vacancies
	DiffData  diffData
}

// MailBody returns body string for mail
func (vs *Vacancies) MailBody() (html string, err error) {
	var diff diffData
	if diff, err = vs.diff(); err != nil {
		err = fmt.Errorf("error occurred in calculating diff: %v", err)
		return
	}
	data := TemplateData{vs, diff}

	tmplString := string(MustAsset(TemplateName))
	tmpl := template.Must(template.New(TemplateName).Parse(tmplString))

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		err = fmt.Errorf("failed to execute template: %v", err)
		return
	}
	html = buf.String()
	return
}

type diffData struct {
	BeforeUpdatedAt *util.Time
	AfterUpdatedAt  *util.Time
	Diff            []diffmatchpatch.Diff
}

func (vs *Vacancies) diff() (data diffData, err error) {
	var cache *Cache
	if cache, err = NewCache(vs); err != nil {
		err = fmt.Errorf("failed to load cache: %v", err)
		return
	}
	data.BeforeUpdatedAt = *cache.UpdatedAt
	data.AfterUpdatedAt = *util.Time.Now()
	data.Diff = vs.rawDiff(cache.Vacancies.String(), vs.String())
	return
}

func (vs *Vacancies) rawDiff(old, new string) (diff []diffmatchpatch.Diff) {
	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars(old, new)
	diffs := dmp.DiffMain(a, b, false)
	diff = dmp.DiffCharsToLines(diffs, c)
	return
}

// CacheFilename is a cache filename
var CacheFilename = path.Join(os.Getenv("HOME"), ".cache/moesia/cache.json")

// Cache is a cache struct
type Cache struct {
	UpdatedAt time.Time
	Vacancies *Vacancies
}

// NewCache returns a new instance
func NewCache(vs *Vacancies) (c *Cache, err error) {
	if _, err = os.Stat(CacheFilename); err != nil {
		c = &Cache{
			UpdatedAt: time.Now(),
			Vacancies: vs,
		}
		err = c.save()
		return
	}
	var content []byte
	if content, err = ioutil.ReadFile(CacheFilename); err != nil {
		err = fmt.Errorf("cannot read file '%s': %v", CacheFilename, err)
		return
	}
	if err = json.Unmarshal(content, &c); err != nil {
		err = fmt.Errorf("cannot unmarshal JSON from '%s': %v", CacheFilename, err)
		return
	}
	return
}

func (c *Cache) save() (err error) {
	if err = os.MkdirAll(path.Dir(CacheFilename), 0700); err != nil {
		err = fmt.Errorf("cannot create parent dir for '%s': %v", CacheFilename, err)
		return
	}
	var content []byte
	if content, err = json.Marshal(c); err != nil {
		err = fmt.Errorf("cannot marshal the instance '%s': %v", c, err)
		return
	}
	if err = ioutil.WriteFile(CacheFilename, content, 0600); err != nil {
		err = fmt.Errorf("cannot write file to '%s': %v", CacheFilename, err)
		return
	}
	return
}
