package handler

import (
	"net/http"

	"user/api/internal/logic"
	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserInfoLogic(r.Context(), ctx)
		resp, err := l.GetUserInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
