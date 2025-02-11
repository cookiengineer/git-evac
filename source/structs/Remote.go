package structs

type Remote struct {
	Name string `json:"name"`
	URL  string `json:"url"` // TODO: should be net/url.URL pointer
}
