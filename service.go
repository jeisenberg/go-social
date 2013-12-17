package service

import (
  "net/http"
  "log"
  "encoding/json"
  //"io/ioutil"
  "github.com/jeisenberg/go-social/translators"
  "github.com/jeisenberg/go-social/normalized"
)


type Service struct {
  Name string
  FetchUrl string
  AppId string
  AppSecret string
  AppKey string
}

//type Receiver struct {}

func (s Service) Fetch() (posts []normalized.Post){
  resp, err := http.Get(s.FetchUrl)
  if err != nil {
    return
  }
  log.Printf("TEST")
  defer resp.Body.Close()
  switch {
  case s.Name == "google plus":
    x := new(translators.GPResults)
    err = json.NewDecoder(resp.Body).Decode(&x)
    if err != nil {
      log.Printf("%s somethgin err", err)
      return
    }
    posts = x.Normalize()
  case s.Name == "instagram":
    x := new(translators.IGResults)
    err = json.NewDecoder(resp.Body).Decode(&x)
    if err != nil {
      log.Printf("%s somethgin err", err)
      return
    }
    posts = x.Normalize()
  }
  return posts
}


