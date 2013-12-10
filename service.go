package service

import (
  "net/http"
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
  revel.INFO.Printf("%s", resp.Body)
}
