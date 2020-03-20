package app

import (
	"acfeedback/pkg/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const contentType = "Content-Type"
const value = "application/json; charset=utf-8"

func (server *MainServer) SaveNewFeedbackHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	var requestBody models.FeedBack
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		//log.Printf("req = %v\n", requestBody)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
//	log.Printf("req = %v\n", requestBody)
	err = server.feedbackSvc.Save(requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}// TODO: normal response
	return
}
func (server *MainServer) GetFeedbackByIdHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	id := pr.ByName("id")
	//log.Printf("%s\n", id)
	feedback, err := server.feedbackSvc.GetFeedbackById(id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(feedback)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}
func (server *MainServer) RemoveFeedbackByIdHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	name := pr.ByName("id")
	//log.Printf("I am here %s\n", len(name))
	err := server.feedbackSvc.RemoveFeedbackByID(name)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(err)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}
func (server *MainServer) GetFeedbackListHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	feedbacks, err := server.feedbackSvc.GetFeedbackList()
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set(contentType, value)
	err = json.NewEncoder(writer).Encode(feedbacks)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	return
}