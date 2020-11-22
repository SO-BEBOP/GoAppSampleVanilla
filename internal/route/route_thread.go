package route

import (
	"GoAppSampleVanilla/pkg/data"
	"GoAppSampleVanilla/pkg/utils"
	"fmt"
	"net/http"
)

// NewThread ...GET /threads/new
// Show the new thread form page
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		utils.GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// CreateThread ...POST /signup
// Create the user account
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			utils.Danger(err, "Cannot get user from utils.Session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			utils.Danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// ReadThread ...GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		utils.ErrorMsg(writer, request, "Cannot read thread")
	} else {
		_, err := utils.Session(writer, request)
		if err != nil {
			utils.GenerateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			utils.GenerateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// PostThread ...POST /thread/post
// Create the post
func PostThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			utils.Danger(err, "Cannot get user from utils.Session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			utils.ErrorMsg(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			utils.Danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
