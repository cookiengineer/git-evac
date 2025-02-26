package api

import "gooey/xhr"
import "git-evac/server/schemas"
import "encoding/json"
import "errors"

func Repositories() (*schemas.Repositories, error) {

	var result_schema *schemas.Repositories = nil
	var result_error error = nil

	request := xhr.NewXMLHttpRequest()
	request.OpenSync(xhr.MethodGET, "/api/repositories")
	request.OnLoad = func(status int, response []byte) {

		if status == 200 || status == 304 {

			var schema schemas.Repositories

			err := json.Unmarshal(response, &schema)

			if err == nil {
				result_schema = &schema
			} else {
				result_error = err
			}

		} else {
			result_error = errors.New("TODO: Error Description")
		}

	}

	request.OnError = func() {
		result_error = errors.New("TODO: Error Description")
	}

	request.OnTimeout = func() {
		result_error = errors.New("TODO: Error Description for Timeout")
	}

	request.SendSync()

	return result_schema, result_error

}
