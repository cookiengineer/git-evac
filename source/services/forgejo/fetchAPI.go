package forgejo

import "git-evac/services/gitea/schemas"
import "encoding/json"
import "fmt"
import "net/http"
import "time"

func fetchAPI(url string, token string) ([]schemas.Repository, int, error) {

	request, err1 := http.NewRequest(http.MethodGet, url, nil)

	if err1 == nil {

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		response, err2 := client.Do(request)

		if err2 == nil {

			defer response.Body.Close()

			if response.StatusCode == http.StatusOK {

				var schema []schemas.Repository

				err3 := json.NewDecoder(response.Body).Decode(&schema)

				if err3 == nil {
					return schema, response.StatusCode, nil
				} else {
					return nil, response.StatusCode, err3
				}

			} else {
				return nil, response.StatusCode, nil
			}

		} else {
			return nil, 500, err2
		}

	} else {
		return nil, 500, err1
	}

}

