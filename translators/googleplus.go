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
  GPObject `json:"object'`
}

type GPResults struct {
  Items []GPItem `json:"items"`
}

type GPObject struct {
  ObjectType string `json:"objectType"`
  Url string  `json:"url"`
  Content string `json:"content"`
}

func (r GPResults) Normalize() (posts []normalized.Post) {
  for _, post := range r.Items {
    //copy the struct value into our own layout type
    posts = append(posts, normalized.Post{
      Title: post.Title,
      AuthorName: post.GPActor.DisplayName,
      Image: post.GPActor.Image.Url,
      Body: post.GPObject.Content,
    })
  }
  return posts
}