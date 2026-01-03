package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MetaHandler struct{}

func NewMetaHandler() http.Handler {
	h := &MetaHandler{}
	r := chi.NewRouter()

	r.Get("/transaction-types", h.transactionTypes)

	return r
}

type TxType struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

func (h *MetaHandler) transactionTypes(w http.ResponseWriter, r *http.Request) {
	list := []TxType{
		{Code: "income", Title: "Доход"},
		{Code: "expense_company", Title: "Расход компании"},
		{Code: "expense_people", Title: "Перевод людям"},
		{Code: "transfer_to_owner", Title: "Отдал владельцу"},
	}

	json.NewEncoder(w).Encode(list)
}
