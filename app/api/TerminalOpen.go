package api

import "gooey/xhr"
import "git-evac/server/schemas"
import "encoding/json"
import "errors"
import "strings"

func TerminalOpen(owner string, repository string) (*schemas.Repository, error) {

	var result_schema *schemas.Repository = nil
	var result_error error = nil

	if !strings.Contains(owner, "/") && !strings.Contains(repository, "/") {

		request := xhr.NewXMLHttpRequest()
		request.OpenSync(xhr.MethodGET, "/api/terminal/open/" + owner + "/" + repository)
		request.OnLoad = func(status int, response []byte) {

			if status == 200 || status == 304 {

				var schema schemas.Repository

				err := json.Unmarshal(response, &schema)

				if err == nil {
					result_schema = &schema
				} else {
					result_error = err
				}

			} else {
				result_error = errors.New("XHR: Invalid HTTP Status Code")
			}

		}

		request.OnError = func() {
			result_error = errors.New("XHR: Unknown Error")
		}

		request.OnTimeout = func() {
			result_error = errors.New("XHR: Timeout")
		}

		request.SendSync()

	}

	return result_schema, result_error

}
