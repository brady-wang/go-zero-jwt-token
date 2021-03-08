package middleware

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
)

type JsonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ExampleMiddleware struct {
}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
}

func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("example middle")

		inToken := r.Header.Get("Authorization")

		logx.Info("收到token:"+inToken)
		if len(inToken) <= 0 {
			msg, _ := json.Marshal(JsonResult{Code: 401, Msg: "token验证失败 缺少token"})
			w.Write(msg)
		} else {
			token, err := jwt.Parse(inToken, func(token *jwt.Token) (interface{}, error) {
				return []byte("ad879037-c7a4-4063-9236-6bfc35d54b7d"), nil
			})

			if token.Valid {
				next(w, r)
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					msg, _ := json.Marshal(JsonResult{Code: 401, Msg: "token验证失败 缺少token"})
					w.Write(msg)
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not active yet
					msg, _ := json.Marshal(JsonResult{Code: 401, Msg: "Token is either expired or not active yet"})
					w.Write(msg)
				} else {
					msg, _ := json.Marshal(JsonResult{Code: 401, Msg: "Couldn't handle this token"})
					w.Write(msg)
				}
			} else {
				msg, _ := json.Marshal(JsonResult{Code: 401, Msg: "Couldn't handle this token:"})
				w.Write(msg)
			}
		}

	}
}
