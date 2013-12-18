package service

import (
  "net/http"
  "log"
  "encoding/json"
  "net/url"
  "bytes"
  "encoding/base64"
  "github.com/jeisenberg/go-social/translators"
  "github.com/jeisenberg/go-social/normalized"
)


type Service struct {
  Name string
  FetchUrl string
  AppId string
  AppSecret string
  AppKey string
  AppAuthUrl string
  Oauth bool
  OauthToken string
}

type Response struct {
  AccessToken string `json:"access_token"`
}

func (s Service) Fetch() (posts []normalized.Post){
  if s.Oauth == true && s.OauthToken != "" {
    return s.FetchOauth(s.OauthToken)
  } else if s.Oauth == true && s.OauthToken == "" {
    token := s.GetOauthToken()
    return s.FetchOauth(token)
  }
  resp, err := http.Get(s.FetchUrl)
  if err != nil {
    return
  }
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

func (s Service) GetOauthToken() string {
  data := []byte(s.AppKey+":"+s.AppSecret)
  encoded := base64.StdEncoding.EncodeToString(data)
  client:= &http.Client{}
  payload := url.Values{}
  payload.Set("grant_type", "client_credentials")
  req, _ := http.NewRequest("POST",s.AppAuthUrl, bytes.NewBufferString(payload.Encode()))
  req.Header.Add("Authorization", "Basic "+encoded)
  newresp, err := client.Do(req)
  if err != nil {
    log.Printf("%s", err)
  }
  defer newresp.Body.Close()
  var r Response
  decoder := json.NewDecoder(newresp.Body)
  newerr := decoder.Decode(&r)
  if newerr != nil {
    log.Printf("%s", newerr)
  }
  return r.AccessToken
}

func (s Service) FetchOauth(token string) (posts []normalized.Post){
  client:= &http.Client{}
  req, _ := http.NewRequest("GET", s.FetchUrl, nil)
  log.Printf("token in callback %s", token)
  req.Header.Add("Authorization", "Bearer "+token)
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  switch {
  case s.Name == "twitter":
    x := new(translators.TWResults)
    err = json.NewDecoder(resp.Body).Decode(&x)
    log.Printf("%s decoder", x)
    if err != nil {
      return
    }
    posts = x.Normalize()
  }
  return posts
}


