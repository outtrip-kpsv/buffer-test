package handlers

import (
  "buff/internal/buff"
  "buff/internal/buff/model"
  "encoding/json"
  "log"
  "net/http"
)

type Handler struct {
  buffer *buff.Buffer
}

func NewHandler(buffer *buff.Buffer) *Handler {
  return &Handler{buffer: buffer}
}

func (h *Handler) AddBuff(w http.ResponseWriter, req *http.Request) {
  err := req.ParseMultipartForm(10 << 20)
  if err != nil {
    http.Error(w, "Ошибка разбора формы", http.StatusBadRequest)
    return
  }

  authHeader := req.Header.Get("Authorization")
  if authHeader == "" {
    http.Error(w, "Missing Authorization header", http.StatusBadRequest)
    return
  }

  fact := model.FactData{
    PeriodStart:         req.FormValue("period_start"),
    PeriodEnd:           req.FormValue("period_end"),
    PeriodKey:           req.FormValue("period_key"),
    IndicatorToMoId:     req.FormValue("indicator_to_mo_id"),
    IndicatorToMoFactId: req.FormValue("indicator_to_mo_fact_id"),
    Value:               req.FormValue("value"),
    FactTime:            req.FormValue("fact_time"),
    IsPlan:              req.FormValue("is_plan"),
    AuthUserId:          req.FormValue("auth_user_id"),
    Comment:             req.FormValue("comment"),
  }
  h.buffer.BuffData.Buffer <- model.Req{
    FactData: fact,
    Token:    authHeader,
  }
  log.Println("Received buffer (comment)", fact.Comment)

  response := map[string]string{"status": "add to queue"}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}
