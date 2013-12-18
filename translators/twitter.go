package translators

import "github.com/jeisenberg/go-social/normalized"

type TWResults struct {
  Data []TWStatus `json:"statuses"`
}

type TWStatus struct {
  User TWUser `json:"user"`
  Text string `json:"text"`
}

type TWUser struct {
  Name string `json:"name"`
  ProfileImageUrl string `json:"profile_image_url"`
}

func (r TWResults) Normalize() (posts []normalized.Post) {
  for _, post := range r.Data {
    posts = append(posts, normalized.Post{
      Body: post.Text,
      AuthorName: post.User.Name,
      Image: post.User.ProfileImageUrl,
    })
  }
  return posts
}
