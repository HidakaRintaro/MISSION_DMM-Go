package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

func (h *handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	strId := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	s := h.app.Dao.Status()
	status, err := s.FindById(ctx, id)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
