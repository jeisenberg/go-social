package translators

import "github.com/jeisenberg/go-social/normalized"

type GPActor struct {
  DisplayName string `json:"displayName"`
  Image `json:"image"`
}

type Image struct {
  Url string `json:"url"`
}

type GPItem struct {
  Kind string `json:"kind"`
  Title string `json:"title"`
  GPActor `json:"actor"`
}

type GPResults struct {
  Items []GPItem `json:"items"`
}

func (r GPResults) Normalize() (posts []normalized.Post) {
  for _, post := range r.Items {
    //copy the struct value into our own layout type
    posts = append(posts, normalized.Post{
      Title: post.Title,
      AuthorName: post.GPActor.DisplayName,
      Image: post.GPActor.Image.Url,
    })
  }
  return posts
}