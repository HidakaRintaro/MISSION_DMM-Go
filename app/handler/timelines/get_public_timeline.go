package timelines

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) GetPublicTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryParameter := r.URL.Query()
	tl := h.app.Dao.Status()
	timeLine, err := tl.FindByQuery(ctx, queryParameter)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	fmt.Println("aaa=========================")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timeLine); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
