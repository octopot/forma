package static_test

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/kamilsk/form-api/domain"
	"github.com/kamilsk/form-api/static"
	"github.com/stretchr/testify/assert"
)

var (
	update = flag.Bool("update", false, "update .golden files")
	pool   = sync.Pool{
		New: func() interface{} {
			blob := [1024]byte{}
			return bytes.NewBuffer(blob[:0])
		},
	}
)

func TestErrorTemplate(t *testing.T) {
	tpl := template.Must(template.New("error").Parse(must("./templates", "error.html")))

	tests := []struct {
		name   string
		data   func() static.ErrorPageContext
		golden string
	}{
		{"email subscription", func() static.ErrorPageContext {
			schema := domain.Schema{
				Language:     "en",
				Title:        "Email subscription",
				Action:       "https://kamil.samigullin.info/",
				Method:       "post",
				EncodingType: "application/x-www-form-urlencoded",
				Inputs: []domain.Input{
					{
						ID:          "email",
						Name:        "email",
						Type:        "email",
						Title:       "Email",
						Placeholder: "Your email...",
						MaxLength:   64,
						Required:    true,
					},
				},
			}
			_, err := schema.Apply(map[string][]string{"email": {"invalid"}})
			return static.ErrorPageContext{Schema: schema, Error: err, Delay: time.Hour, Redirect: schema.Action}
		}, "./fixtures/email_subscription.html.golden"},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if *update {
				file := writer(tc.golden)
				assert.NoError(t, closeAfter(file, func() error {
					return tpl.Execute(file, tc.data())
				}))
			}

			buf := pool.Get().(*bytes.Buffer)
			assert.NoError(t, tpl.Execute(buf, tc.data()))
			expected, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			assert.Equal(t, expected, buf.Bytes())
			buf.Reset()
			pool.Put(buf)
		})
	}
}

func closeAfter(file *os.File, action func() error) error {
	defer file.Close()
	if err := action(); err != nil {
		return err
	}
	return nil
}

func must(base, tpl string) string {
	b, err := static.LoadTemplate(base, tpl)
	if err != nil {
		panic(tpl)
	}
	return string(b)
}

func writer(file string) *os.File {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	return f
}
