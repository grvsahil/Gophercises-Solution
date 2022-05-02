package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"gopkg.in/yaml.v2"
)

type routes struct {
	Path string
	Url  string
}

func welHandler(myMap map[string]string) func(rp http.ResponseWriter, rq *http.Request) {
	return func(rp http.ResponseWriter, rq *http.Request) {
		rp.Header().Set("Content-Type", "text/html; charset=utf-8")

		if url,ok := myMap[rq.URL.Path]; ok {
			http.Redirect(rp,rq,url,301)
			return
		}
	
		fmt.Fprintln(rp, "<h2> /name to redirect to www.name.com </h2>")
	}
}

func ymlHandler(ymlFile string) []routes {
	fileBytes, err := os.ReadFile(ymlFile)
	if err != nil {
		log.Fatal(err)
	}

	var ymlPaths []routes

	err = yaml.Unmarshal(fileBytes, &ymlPaths)
	if err != nil {
		log.Fatal(err)
	}

	return ymlPaths
}

func jsonHandler(jsonFile string) []routes {
	fileBytes, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var jsonPaths []routes

	err = yaml.Unmarshal(fileBytes, &jsonPaths)
	if err != nil {
		log.Fatal(err)
	}

	return jsonPaths
}

func mapHandler(paths []routes) map[string]string {
	myMap := make(map[string]string)
	for _, route := range paths {
		myMap[route.Path] = route.Url
	}
	return myMap
}

func main() {
	ymlFlag := flag.String("yml", "urls.yml", "set yml file to use for routes")
	flag.Parse()

	myYmlPaths := ymlHandler(*ymlFlag)

	myYmlMap := mapHandler(myYmlPaths)

	// jsonFlag := flag.String("json","urls.json","set json file to use for routes")
	// flag.Parse()

	// myJsonPaths := jsonHandler(*jsonFlag)

	// myJsonMap := mapHandler(myJsonPaths)

	http.ListenAndServe(":9091", http.HandlerFunc(welHandler(myYmlMap)))

}
