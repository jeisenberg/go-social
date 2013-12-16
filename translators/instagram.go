package translators

import "github.com/jeisenberg/go-social/normalized"

type IGResults struct {
  Data []IGMedia
}

type IGMedia struct {
  Comments []IGComments
  Images IGImages
  Caption IGCaption
  User IGUser
}

type IGComment struct {

}

type IGImages struct {
  LowResolution IGImage `json:"low_resolution"`
  Thumbnail IGImage `json:"thumbnail"`
  StandardResolution IGImage `json:"standard_resolution"`
}

type IGImage struct {
  Url string `json:"url"`
  Width int `json:"width"`
  Height int `json:"height"`
}

type IGCaption struct {
  Text byte
}

type IGUser struct {
  Username string `json:"username"`
  ProfilePicture string `json:"profile_picture"`
  FullName string `json:"full_name"`
}