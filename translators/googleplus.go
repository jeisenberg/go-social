package translators

type GPActor struct {
  DisplayName string `json:"displayName"`
  Image map[string]string `json:"image"`
}

type GPItem struct {
  Kind string `json:"kind"`
  Title string `json:"title"`
  GPActor `json:"actor"`
}

type GPResults struct {
  Items []GPItem `json:"items"`
}