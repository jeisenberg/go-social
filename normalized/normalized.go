package normalized

//this is the structure that we will normalize all posts to

//keeping it simple for now
type Post struct {
  Title string
  AuthorName string
  Image string
  Body []byte
}