package api

import (
  "buff/internal/buff"
  "buff/internal/buff/model"
  "buff/internal/io/api/util"
  "log"
  "net/http"
)

type Api struct {
  ApiUrl string
  Method string
  Buffer *buff.Buffer
}

func NewApi(apiUrl string, method string, buffer *buff.Buffer) Api {
  return Api{
    ApiUrl: apiUrl,
    Method: method,
    Buffer: buffer,
  }
}

func (a *Api) Listen() {
  go func() {
    for {
      select {
      case buffer := <-a.Buffer.BuffData.Buffer:
        data, ct, err := util.CreateMultipartData(buffer.FactData)
        if err != nil {
          log.Printf("Error creating multipart data: %v\n", err)
          continue
        }
        a.Buffer.BuffData.Workers <- model.Worker{
          Url:    a.ApiUrl,
          Method: ct,
          Data:   data,
          Token:  buffer.Token,
        }
      default:
      }
    }
  }()

  go func() {
    client := &http.Client{}
    for worker := range a.Buffer.BuffData.Workers {
      req, err := http.NewRequest("POST", worker.Url, worker.Data)
      if err != nil {
        log.Printf("Error creating request: %v\n", err)
        continue
      }
      req.Header.Set("Authorization", worker.Token)
      req.Header.Set("Content-Type", worker.Method)

      worker.Resp, worker.Err = client.Do(req)
      if worker.Err != nil {
        log.Printf("Error sending request: %v\n", worker.Err)
        worker.Err = nil
        continue
      }
      log.Printf("Worker done Response status: %s\n", worker.Resp.Status)
      worker.Resp.Body.Close()
    }
  }()

}
