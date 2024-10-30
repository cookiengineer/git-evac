package handlers

import "net/http"

func RespondWith(request *http.Request, response http.ResponseWriter, status int) {

	response.WriteHeader(status)
	response.Write([]byte{})

}
