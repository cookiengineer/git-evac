package actions

import "gooey/fetch"
import "git-evac/structs"
import "git-evac/server/schemas"
import "bytes"
import "encoding/json"

func SaveSettings(settings *structs.Settings) (*schemas.Settings, error) {

	var result_schema *schemas.Settings = nil
	var result_error error = nil

	payload, err0 := json.MarshalIndent(settings, "", "\t")

	if err0 == nil {

		response, err1 := fetch.Fetch("/api/settings", &fetch.Request{
			Method:   fetch.MethodPOST,
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
