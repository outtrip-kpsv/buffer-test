package buff

import (
  "buff/internal/buff/model"
)

type Buffer struct {
  BuffData model.BufChan
}

func NewBuffer(sizeBuff int, sizeWorker int) *Buffer {
  return &Buffer{
    BuffData: NewBuf(sizeBuff, sizeWorker),
  }
}

func NewBuf(sizeBuff int, sizeWorker int) model.BufChan {
  return model.BufChan{
    Buffer:  make(chan model.Req, sizeBuff),
    Workers: make(chan model.Worker, sizeWorker),
  }
}
