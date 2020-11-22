package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"GoAppSampleVanilla/pkg/data"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

// Config ...configuration struct.
var Config Configuration

var logger *log.Logger
var appVersion string = "v0.1"

// P ...Convenience function for printing to stdout.
func P(a ...interface{}) {
	fmt.Println(a)
}

func init() {
	loadConfig()
	file, err := os.OpenFile("docs/log/GoAppSampleVanilla.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("configs/config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// ErrorMsg ...Convenience function to redirect to the error message page.
func ErrorMsg(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// Session ...Checks if the user is logged in and has a session, if not err is not nil.
func Session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// ParseTemplateFiles ...parse HTML templates
// pass in a list of file names, and get a template.
func ParseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("web/templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// GenerateHTML ...generate HTML templates.
func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("web/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// Info ...logging type.
func Info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// Danger ...logging type.
func Danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

// Warning ...logging type.
func Warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// Version ...loging label.
func Version() string {
	return appVersion
}
