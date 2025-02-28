package api

import "gooey/fetch"
import "git-evac/server/schemas"
import "encoding/json"
import "strings"

func GitPull(owner string, repository string) (*schemas.Repository, error) {

	var result_schema *schemas.Repository = nil
	var result_error error = nil

	if !strings.Contains(owner, "/") && !strings.Contains(repository, "/") {

		response, err1 := fetch.Fetch("/api/git/pull/" + owner + "/" + repository, &fetch.Request{
			Method:   fetch.MethodGET,
			Mode:     fetch.ModeSameOrigin,
			Cache:    fetch.CacheDefault,
			Redirect: fetch.RedirectError,
			Headers:  map[string]string{
				"Accept": "application/json",
			},
		})

		if err1 == nil {

			schema := schemas.Repository{}
			err2   := json.Unmarshal(response.Body, &schema)

			if err2 == nil {
				result_schema = &schema
				result_error  = nil
			} else {
				result_error = err2
			}

		} else {
			result_error = err1
		}

	}

	return result_schema, result_error

}
