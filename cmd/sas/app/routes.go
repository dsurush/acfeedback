package app


const saveFeedBack = `/feedbacks/newFeedback/save`
const getFeedBackByID = `/feedbacks/:id`
const removeFeedbackById = `/feedbacks/:id/remove`
const getFeedbackList = `/feedbacks`

func (server *MainServer) InitRouts() {
	server.router.POST(saveFeedBack, server.SaveNewFeedbackHandler)
	server.router.GET(getFeedBackByID, server.GetFeedbackByIdHandler)
	server.router.DELETE(removeFeedbackById, server.RemoveFeedbackByIdHandler)
	server.router.GET(getFeedbackList, server.GetFeedbackListHandler)
}
