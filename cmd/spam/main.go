package main

import (
  "buff/internal/buff/model"
  "buff/internal/io/api/util"
  "fmt"
  "log"
  "net/http"
)

func main() {

  factData := model.FactData{
    PeriodStart:         "2024-05-01",
    PeriodEnd:           "2024-05-31",
    PeriodKey:           "month",
    IndicatorToMoId:     "227373",
    IndicatorToMoFactId: "0",
    Value:               "1",
    FactTime:            "2024-05-31",
    IsPlan:              "0",
    AuthUserId:          "40",
    Comment:             "buffer outtrip",
  }

  client := &http.Client{}
  for worker := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
    factData.Comment = fmt.Sprintf("buffer outtrip: %d", worker)
    data, contentType, err := util.CreateMultipartData(factData)
    req, err := http.NewRequest("POST", "http://localhost:3000/api/setfact", data)
    if err != nil {
      log.Printf("Error creating request: %v\n", err)
      continue
    }
    req.Header.Set("Authorization", "Bearer "+"48ab34464a5573519725deb5865cc74c")
    req.Header.Set("Content-Type", contentType)
    //
    resp, err := client.Do(req)
    if err != nil {
      log.Printf("Error sending request: %v\n", err)
      err = nil
      continue
    }
    log.Printf("Worker done Response status: %s\n", resp.Status)
    resp.Body.Close()
  }

}
