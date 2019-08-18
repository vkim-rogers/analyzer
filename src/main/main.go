package main

import (
	"connectors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"models"
	"net/http"
	"strconv"
)

// go "github.com/gorilla/mux"

var mc = connectors.InitMongo()
var sc = connectors.InitSqs()

func main(){

	router := mux.NewRouter()
	router.HandleFunc("/quizSubmit", QuizSubmit).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}

func QuizSubmit(w http.ResponseWriter, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		_, _ = w.Write([]byte("Error parsing request body"))
		log.Fatal(err)
	}

	quizAnswers := mc.GetEntityById("quizAnswers", string(bodyBytes))


	var res int
	for _, qa := range quizAnswers.(models.QuizAnswers).Questions{
		value, _ := strconv.Atoi(qa.Answer)

		res += value
	}

	_, _ = w.Write([]byte(fmt.Sprintf("%d", res)))

	err = mc.UpdateEntityByIdAndField("personalityMap", quizAnswers.(models.QuizAnswers).PersonId, "total", strconv.Itoa(res))
	if err != nil {
		_, _ = w.Write([]byte("Error updating personality map with the quiz results"))
		log.Fatal(err)
	}
}