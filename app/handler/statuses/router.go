package statuses

import (
	"net/http"
	"yatter-backend-go/app/app"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi"
)

type handler struct {
	app *app.App
}

// Create Handler for `/v1/statuses`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()
	h := &handler{app: app}

	r.With(auth.Middleware(app)).Post("/", h.Create)
	r.Get("/{id}", h.GetStatus)
	r.Delete("/{id}", h.Delete)

	return r
}
