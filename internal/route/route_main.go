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

	type detail struct {
		Session    data.User
		Data       []data.Thread
		LoginState string
	}

	threads, err := data.Threads()
	if err != nil {
		utils.ErrorMsg(writer, request, "Cannot get Threads")

	} else {
		session, err := utils.Session(writer, request)
		if err != nil {
			details := detail{
				Data:       threads,
				LoginState: "false",
			}
			utils.GenerateHTML(writer, details, "layout", "public.navbar", "index")

		} else {
			user, _ := data.UserByEmail(session.Email)
			details := detail{
				Session:    user,
				Data:       threads,
				LoginState: "ture",
			}
			utils.GenerateHTML(writer, details, "layout", "private.navbar", "index")
		}
	}
}
