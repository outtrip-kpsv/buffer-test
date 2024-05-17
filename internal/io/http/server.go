package http

import (
  "buff/internal/buff"
  "log"
  "net/http"
)

type HttpServer struct {
  buffer   *buff.Buffer // Бизнес логика
  finished chan bool    // Канал о завершении работы сервера
}

func NewHTTPServer(buffer *buff.Buffer, finished chan bool) *HttpServer {
  return &HttpServer{
    buffer:   buffer,
    finished: finished,
  }
}

func (s HttpServer) Run(address string) {
  httpServer := &http.Server{
    Addr:    address,
    Handler: InitRoutes(s.buffer),
  }
  go func() {
    if err := httpServer.ListenAndServe(); err != nil {
      log.Fatal(err.Error())
    }
    s.finished <- true
  }()

  log.Printf("server start address: %s", address)

}
