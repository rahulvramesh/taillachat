package api

import (
	"github.com/elcct/taillachat/helpers"
	"html/template"
	"net/http"
)

// Index serves default index page
func Index(w http.ResponseWriter, r *http.Request) {
	t := r.Context().Value("template").(*template.Template)
	widgets := helpers.Parse(t, "home", nil)

	data := map[string]interface{}{}
	data["Content"] = template.HTML(widgets)

	data["Title"] = "Tailla Chat - Best UK Chat"
	data["SocketURL"] = "/chat"

	w.Write([]byte(helpers.Parse(t, "main", data)))
}
