package structs

type Remote struct {
	Name string `json:"name"`
	URL  string `json:"url"` // TODO: should be net/url.URL pointer
}

func NewRemote(name string, url string) Remote {

	var remote Remote

	remote.Name = name
	remote.URL  = url

	return remote

}
