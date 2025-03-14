package actions

import "gooey/fetch"
import "git-evac/server/schemas"
import "encoding/json"

func ReadSettings() (*schemas.Settings, error) {

	var result_schema *schemas.Settings = nil
	var result_error error = nil

	response, err1 := fetch.Fetch("/api/settings", &fetch.Request{
		Method:   fetch.MethodGet,
		Mode:     fetch.ModeSameOrigin,
		Cache:    fetch.CacheDefault,
		Redirect: fetch.RedirectError,
		Headers:  map[string]string{
			"Accept": "application/json",
		},
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

	return result_schema, result_error

}
