package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8000"
	}
	api := rest.NewApi()
	router, _ := rest.MakeRouter(
		rest.Get("/api/mercari/seller/getMercariBlackSellerIds", func(res rest.ResponseWriter, req *rest.Request) {
			buff, _ := ioutil.ReadFile("./responses/blackList.json")
			var dat []string
			json.Unmarshal(buff, &dat)
			res.WriteJson(dat)
		}),
	)
	api.SetApp(router)
	log.Print(fmt.Sprintf("Service run on: %v", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), api.MakeHandler()))
}
