package router

import (
	"database/sql"
	"net/http"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {

	// NOTE: ここに新しいエンドポイントを登録する
	mux := http.NewServeMux()
	return mux
}
