package model

import (
  "io"
  "net/http"
)

type Worker struct {
  Url    string
  Method string
  Data   io.Reader
  Token  string
  Resp   *http.Response
  Err    error
}
