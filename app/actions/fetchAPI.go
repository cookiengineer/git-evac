package actions

import "github.com/cookiengineer/gooey/bindings/fetch"
import "git-evac/schemas"
import "encoding/json"
import "errors"
import "strconv"
import "strings"

func fetchAPI(method string, path string, owner string, repository string) (*schemas.Repository, error) {

	var result_schema *schemas.Repository = nil
	var result_error error = nil

	if strings.HasPrefix(path, "/api/") && !strings.Contains(owner, "/") && !strings.Contains(repository, "/") {

		response, err1 := fetch.Fetch(path + "/" + owner + "/" + repository, &fetch.Request{
			Method:   fetch.Method(method),
			Mode:     fetch.ModeSameOrigin,
			Cache:    fetch.CacheDefault,
			Redirect: fetch.RedirectError,
			Headers:  map[string]string{
				"Accept": "application/json",
			},
		})

		if err1 == nil {

			if response.Status == 200 || response.Status == 304 {

				schema := schemas.Repository{}
				err2   := json.Unmarshal(response.Body, &schema)

				if err2 == nil {
					result_schema = &schema
					result_error  = nil
				} else {
					result_error = err2
				}

			} else {
				result_error = errors.New(strconv.Itoa(response.Status) + ": " + response.StatusText)
			}

		} else {
			result_error = err1
		}

	}

	return result_schema, result_error

}
