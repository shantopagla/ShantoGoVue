package api

import "net/http"

// View -
func (api *API) View(w http.ResponseWriter, req *http.Request) {
	view := api.views.RandomView()
	w.Write([]byte(view.Text))
}

// SecretView -
func (api *API) SecretView(w http.ResponseWriter, req *http.Request) {
	view := api.views.RandomView()
	w.Write([]byte(view.Text))
}