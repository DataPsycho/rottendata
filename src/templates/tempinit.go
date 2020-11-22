package templates

import "text/template"

// TPL Load All the templates
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
