package controllers

import (
	"github.com/dpgolang/PetBook/pkg/logger"
	"github.com/dpgolang/PetBook/pkg/models"
	"github.com/dpgolang/PetBook/pkg/view"
	"net/http"
)

func (c *Controller) ViewTopicsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		topics, err := c.TopicStore.GetAllTopics()
		if err != nil {
			logger.Error(err)
		}

		view.GenerateHTML(w, topics, "viewTopics")
	}
}

// TODO: get user_id with the help of email from context
//  or immediately get user_id from context (NOT IMPLEMENTED)
func (c *Controller) NewTopicHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			view.GenerateHTML(w, nil, "newTopic")
		}

		if r.Method == http.MethodPost {
			r.ParseForm()
			title := r.FormValue("title")
			description := r.FormValue("description")


			topic := models.Topic{
				UserID:      2,
				Title:       title,
				Description: description,
			}
			if err := c.TopicStore.CreateNewTopic(&topic); err != nil {
				logger.Error(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/forum/new_topic", http.StatusFound)
		}
	}
}