package route

import (
	"GoAppSampleVanilla/pkg/data"
	"GoAppSampleVanilla/pkg/utils"
	"net/http"
)

// Err ...GET /err?msg=
// shows the error message page
func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := utils.Session(writer, request)
	if err != nil {
		utils.GenerateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.GenerateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

// Index ...front page public or private.
func Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		utils.ErrorMsg(writer, request, "Cannot get threads")
	} else {
		_, err := utils.Session(writer, request)
		if err != nil {
			utils.GenerateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			utils.GenerateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
