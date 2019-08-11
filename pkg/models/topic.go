package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Topic struct {
	TopicID     int       `json:"topic_id" db:"topic_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
}

type TopicStorer interface {
	GetAllTopics() (topics []*Topic, err error)
	CreateNewTopic(topic *Topic) (err error)
}

type TopicStore struct {
	DB *sqlx.DB
}

func (t *TopicStore) GetAllTopics() (topics []*Topic, err error) {
	rows, err := t.DB.Query("select * from topics order by created_time DESC")
	if err != nil {
		err = fmt.Errorf("Can't read topics-rows from db: %v", err)
		return
	}
	defer rows.Close()
	err = sqlx.StructScan(rows, &topics)
	if err != nil {
		err = fmt.Errorf("Can't scan topics-rows from db: %v", err)
	}
	fmt.Println(topics)
	return
}

func (t *TopicStore) CreateNewTopic(topic *Topic) (err error) {
	_, err = t.DB.Exec(
		`insert into topics (user_id, title, description) values ($1, $2, $3)`,
		topic.UserID, topic.Title, topic.Description)
	if err != nil {
		return fmt.Errorf("cannot affect rows in pets in db: %v", err)
	}
	return
}
