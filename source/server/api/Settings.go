package api

import "git-evac/console"
import "git-evac/server/schemas"
import "git-evac/structs"
import "encoding/json"
import "io"
import "net/http"
import "os"

func Settings(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		payload, err := json.MarshalIndent(schemas.Settings{
			Settings: profile.Settings,
		}, "", "\t")

		if err == nil {

			console.Log("> api.Settings()")

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)
			response.Write(payload)

		} else {

			console.Error("> api.Settings()")

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("[]"))

		}

	} else if request.Method == http.MethodPost {

		bytes, err0 := io.ReadAll(request.Body)

		if err0 == nil {

			var schema schemas.Settings

			err1 := json.Unmarshal(bytes, &schema)

			if err1 == nil && schema.IsValid() {

				profile.Settings.Backup        = schema.Settings.Backup
				profile.Settings.Folder        = schema.Settings.Folder
				profile.Settings.Port          = schema.Settings.Port
				profile.Settings.Organizations = schema.Settings.Organizations

				stat, err2 := os.Stat(profile.Settings.Folder)

				if err2 == nil && stat.IsDir() {

					payload, _ := json.MarshalIndent(schemas.Settings{
						Settings: profile.Settings,
					}, "", "\t")

					err4 := os.WriteFile(profile.Settings.Folder + "/git-evac.json", payload, 0666)

					if err4 == nil {

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusOK)
						response.Write(payload)

					} else {

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusInternalServerError)
						response.Write([]byte("{}"))

					}

				} else {

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusConflict)
					response.Write([]byte("{}"))

				}

			} else {

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte("{}"))

			}

		} else {

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("{}"))

		}

	} else {

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("[]"))

	}

}
