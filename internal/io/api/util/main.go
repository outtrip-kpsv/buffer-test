package util

import (
  "buff/internal/buff/model"
  "bytes"
  "io"
  "mime/multipart"
)

func CreateMultipartData(fact model.FactData) (io.Reader, string, error) {
  var buf bytes.Buffer
  writer := multipart.NewWriter(&buf)

  err := writer.WriteField("period_start", fact.PeriodStart)
  if err != nil {
    return nil, "", err
  }
  err = writer.WriteField("period_end", fact.PeriodEnd)
  if err != nil {
    return nil, "", err
  }
  err = writer.WriteField("period_key", fact.PeriodKey)
  if err != nil {
    return nil, "", err
  }
  err = writer.WriteField("indicator_to_mo_id", fact.IndicatorToMoId)
  if err != nil {
    return nil, "", err
  }
  err = writer.WriteField("indicator_to_mo_fact_id", fact.IndicatorToMoFactId)
  if err != nil {
    return nil, "", err
  }
  err = writer.WriteField("value", fact.Value)
  if err != nil {
    return nil, "", err
  }
  err = writer.WriteField("fact_time", fact.FactTime)
  if err != nil {
    return nil, "", err
  }
  err = writer.WriteField("is_plan", fact.IsPlan)
  if err != nil {
    return nil, "", err
  }
  err = writer.WriteField("auth_user_id", fact.AuthUserId)
  if err != nil {
    return nil, "", err
  }
  err = writer.WriteField("comment", fact.Comment)
  if err != nil {
    return nil, "", err
  }

  err = writer.Close()
  if err != nil {
    return nil, "", err
  }

  return &buf, writer.FormDataContentType(), nil
}
