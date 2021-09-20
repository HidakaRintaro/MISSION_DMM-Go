package timelines

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) GetPublicTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryParameter := r.URL.Query()
	// TODO: テーブル結合で１発でデータを取る方法がありそう
	tl := h.app.Dao.Status()
	timeLine, err := tl.FindByQuery(ctx, queryParameter)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	a := h.app.Dao.Account()
	for _, r := range timeLine {
		account, err := a.FindById(ctx, r.AccountID)
		if err != nil {
			httperror.InternalServerError(w, err)
			return
		}
		r.Account = *account
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timeLine); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
