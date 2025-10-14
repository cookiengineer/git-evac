package actions

import "github.com/cookiengineer/gooey/bindings/fetch"
import "git-evac/schemas"
import "bytes"
import "encoding/json"

func SaveSettings(settings schemas.Settings) (*schemas.Settings, error) {

	var result_schema *schemas.Settings = nil
	var result_error error = nil

	payload, err0 := json.MarshalIndent(settings, "", "\t")

	if err0 == nil {

		response, err1 := fetch.Fetch("/api/settings", &fetch.Request{
			Method:   fetch.MethodPost,
			Mode:     fetch.ModeSameOrigin,
			Cache:    fetch.CacheDefault,
			Redirect: fetch.RedirectError,
			Headers:  map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
			Body: bytes.NewReader(payload),
		})

		if err1 == nil {

			schema := schemas.Settings{}
			err2   := json.Unmarshal(response.Body, &schema)

			if err2 == nil {
				result_schema = &schema
				result_error  = nil
			} else {
				result_error = err2
			}

		}

	} else {
		result_error = err0
	}

	return result_schema, result_error

}
