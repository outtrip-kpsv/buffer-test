package http

import (
  "buff/internal/buff"
  "buff/internal/io/http/handlers"
  "net/http"
)

type router struct {
  bl     *buff.Buffer
  router *http.ServeMux
}

func InitRoutes(bl *buff.Buffer) http.Handler {
  r := &router{
    bl:     bl,
    router: http.NewServeMux(),
  }

  r.initRoutes(handlers.NewHandler(bl))
  return r.router
}

func (r *router) initRoutes(handlers *handlers.Handler) {
  r.router.HandleFunc("/api/setfact", handlers.AddBuff)
}
