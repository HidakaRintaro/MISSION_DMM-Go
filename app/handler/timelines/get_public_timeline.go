package timelines

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) GetPublicTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var queryParameter = map[string]int{"limit": 40}
	// TODO: ここのエラーチェックはもうちょっとどうにかなりそう(スライスとforを合わせたり、関数にしたり)
	// TODO: max_id <= since_id のチェックしてエラーを返すか、空の配列を返すのか
	if v := r.URL.Query().Get("max_id"); v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			httperror.BadRequest(w, err)
			return
		}
		queryParameter["max_id"] = i
	}
	if v := r.URL.Query().Get("since_id"); v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			httperror.BadRequest(w, err)
			return
		}
		queryParameter["since_id"] = i
	}
	if v := r.URL.Query().Get("limit"); v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			httperror.BadRequest(w, err)
			return
		}
		if i > 80 {
			httperror.BadRequest(w, errors.New("limit is up to 80"))
			return
		}
		queryParameter["limit"] = i
	}

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
