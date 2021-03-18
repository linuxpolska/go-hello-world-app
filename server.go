package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"os"
)

const html_content = `<!DOCTYPE html>
<html>
<head>
<style>
h1 {text-align: center;}
p {text-align: center;}
.center {
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 51%;
}
body {background-color: green;}
</style>
<title>Hello World</title>
</head>
<body>
<p>The date today is {{.Date}}</p>
<p>And the time is {{.Time}}</p>
<img
src="https://images.pexels.com/photos/1114900/pexels-photo-1114900.jpeg?cs=srgb&dl=pexels-johannes-plenio-1114900.jpg&fm=jpg"
alt="Girl in a jacket"
style="width:700px;height:600px;"
class="center">
<p>Environment is {{.Environment}}</p>
<p>App version is {{.Version}}</p>
</body>
</html>`

type PageVariables struct {
	Date         string
	Time         string
	Version 	 string
	Environment  string
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Print("Sever is starting at port 8080")
	err := http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}

}

func HomePage(w http.ResponseWriter, r *http.Request){

    now := time.Now()

    HomePageVars := PageVariables{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
		Version: getVersion(),
		Environment: getEnvironment(),
    }

	t, err := template.New("HomePageVars").Parse(html_content)
    if err != nil {
  	  log.Print("template parsing error: ", err)
  	}
    err = t.Execute(w, HomePageVars)
    if err != nil {
  	  log.Print("template executing error: ", err)
  	}
}

func getVersion() string {
	version := os.Getenv("APP_VERSION")
	if (len(version) == 0) {
		version = "not specified"
	}
	return version
}

func getEnvironment() string {
	environment := os.Getenv("APP_ENVIRONMENT")
	if (len(environment) == 0) {
		environment = "not specified"
	}
	return environment
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.RemoteAddr, r.Method)
		handler.ServeHTTP(w, r)
	})
}
