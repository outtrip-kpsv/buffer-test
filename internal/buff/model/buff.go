package model

type BufChan struct {
  Buffer  chan Req
  Workers chan Worker
}
