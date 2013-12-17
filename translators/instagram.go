package translators

import "github.com/jeisenberg/go-social/normalized"

type IGResults struct {
  Data []IGMedia
}

type IGMedia struct {
  //Comments []IGComments
  Images IGImages `json:"images"`
  Caption IGCaption `json:"caption"`
  User IGUser `json:"user"`
}

// type IGComment struct {

// }

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
  Text string
}

type IGUser struct {
  Username string `json:"username"`
  ProfilePicture string `json:"profile_picture"`
  FullName string `json:"full_name"`
}

func (r IGResults) Normalize() (posts []normalized.Post) {
  for _, post := range r.Data {
    posts = append(posts, normalized.Post{
      Title: post.Caption.Text,
      AuthorName: post.User.Username,
      Image: post.User.ProfilePicture,
      BodyPhoto: post.Images.StandardResolution.Url,
    })
  }
  return posts
}