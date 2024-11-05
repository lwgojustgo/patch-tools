package xtemplate

import (
	"io"
	"io/fs"
	"text/template"
)

type Template struct {
	t *template.Template
}
type FuncMap template.FuncMap

func (t *Template) Funcs(funcMap FuncMap) *Template {
	t.t = t.t.Funcs(template.FuncMap(funcMap))
	return t
}

func New(name string) *Template {
	t := &Template{}
	t.t = template.New(name)
	return t
}

func (t *Template) Option(opt ...string) *Template {
	t.t = t.t.Option(opt...)
	return t
}

func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
	x, e := t.t.ParseFiles(filenames...)
	return &Template{t: x}, e
}

func (t *Template) ParseFS(fsys fs.FS, patterns ...string) (*Template, error) {
	x, e := t.t.ParseFS(fsys, patterns...)
	return &Template{t: x}, e
}

func (t *Template) Clone() (*Template, error) {
	x, e := t.t.Clone()
	return &Template{t: x}, e
}

func (t *Template) New(name string) *Template {
	x := t.t.New(name)
	return &Template{t: x}
}

func (t *Template) Execute(wr io.Writer, data any) error {
	return t.t.Execute(wr, data)
}

func (t *Template) Parse(text string) (*Template, error) {
	x, e := t.t.Parse(text)
	return &Template{t: x}, e
}

func Must(t *Template, err error) *Template {
	x := template.Must(t.t, err)
	return &Template{t: x}
}

func ParseFiles(filenames ...string) (*Template, error) {
	x, e := template.ParseFiles(filenames...)
	return &Template{t: x}, e
}
