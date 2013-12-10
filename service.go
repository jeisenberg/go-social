package service

import (
  "net/http"
  "log"
)

type Service struct {
  Name string
  FetchUrl string
  AppId string
  AppSecret string
  AppKey string
}

func (s Service) Fetch() {
  resp, err := http.Get(s.FetchUrl)
  if err != nil {
    return nil
  }
  defer resp.Body.Close()
  log.Printf("%s", resp.Body)
}
