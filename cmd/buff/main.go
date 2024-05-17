package main

import (
  "buff/internal/buff"
  "buff/internal/io/api"
  "buff/internal/io/http"
)

func main() {

  buffFact := buff.NewBuffer(1000, 1)

  fin := make(chan bool)

  app := api.NewApi("https://development.kpi-drive.ru/_api/facts/save_fact", "multipart/form-data", buffFact)
  app.Listen()

  srv := http.NewHTTPServer(buffFact, fin)
  srv.Run("localhost:3000")

  <-fin
}
