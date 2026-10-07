package routes

import "git-evac/schemas"
import "git-evac/structs"
import "encoding/json"
import "io"
import "net/http"
import "os"
import "path/filepath"

func Settings(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		payload, _ := json.MarshalIndent(schemas.Settings{
			Settings: profile.Settings,
		}, "", "\t")

		profile.Console.Log("> " + request.Method + " /api/settings: " + http.StatusText(http.StatusOK))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(payload)

	} else if request.Method == http.MethodPost {

		bytes0, err0 := io.ReadAll(request.Body)

		if err0 != nil {

			profile.Console.Error("> " + request.Method + " /api/settings: " + http.StatusText(http.StatusBadRequest))

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("{}"))

			return

		}

		var schema schemas.Settings

		err1 := json.Unmarshal(bytes0, &schema)

		if err1 != nil || schema.IsValid() == false {

			profile.Console.Error("> " + request.Method + " /api/settings: " + http.StatusText(http.StatusBadRequest))

			if err1 != nil {
				profile.Console.Error("> " + err1.Error())
			}

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("{}"))

			return

		}

		profile.Update(schema.Settings)

		config := profile.Settings.GetConfig()

		if config == "" {

			profile.Console.Error("> " + request.Method + " /api/settings: " + http.StatusText(http.StatusInternalServerError))
			profile.Console.Error("> No config file path configured")

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("{}"))

			return

		}

		err2 := os.MkdirAll(filepath.Dir(config), 0755)

		if err2 == nil {

			payload_file, _ := json.MarshalIndent(profile.Settings, "", "\t")

			err3 := os.WriteFile(config, payload_file, 0666)

			if err3 == nil {

				profile.Refresh()

				payload, _ := json.MarshalIndent(schemas.Settings{
					Settings: profile.Settings,
				}, "", "\t")

				profile.Console.Log("> " + request.Method + " /api/settings: " + http.StatusText(http.StatusOK))

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {

				profile.Console.Error("> " + request.Method + " /api/settings: " + http.StatusText(http.StatusInternalServerError))
				profile.Console.Error("> " + err3.Error())

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte("{}"))

			}

		} else {

			profile.Console.Error("> " + request.Method + " /api/settings: " + http.StatusText(http.StatusInternalServerError))
			profile.Console.Error("> " + err2.Error())

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("{}"))

		}

	} else {

		profile.Console.Error("> " + request.Method + " /api/settings: " + http.StatusText(http.StatusMethodNotAllowed))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("[]"))

	}

}
