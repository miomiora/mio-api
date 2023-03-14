package routes

import "mio-api-interface/api"

func addSentenceRouter() {
	apiGroup.GET("/sentence", api.RootApi.GetSentence)
}
