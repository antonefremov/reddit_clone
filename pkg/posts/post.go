package posts

import (
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/user"
	"time"
)

// Post struct which contains full information about posts
type Post struct {
	ID               string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title            string    `json:"title" bson:"title"`
	Url              string    `json:"url" bson:"url"`
	Author           user.User `json:"author" bson:"author"`
	Category         string    `json:"category" bson:"category"`
	Score            int       `json:"score" bson:"score"`
	Votes            []Vote    `json:"votes" bson:"votes"`
	Comments         []Comment `json:"comments" bson:"comments"`
	Created          time.Time `json:"created" bson:"created"`
	Views            int       `json:"views" bson:"views"`
	Type             string    `json:"type" bson:"type"`
	Text             string    `json:"text" bson:"text"`
	UpvotePercentage uint8     `json:"upvotePercentage" bson:"upvotePercentage"`
}

// Vote counts votes from users
type Vote struct {
	User string `json:"user"`
	Vote int    `json:"vote"`
}

// Comment object
type Comment struct {
	ID      string    `json:"id"`
	Author  user.User `json:"author"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

// NetworkComment represents json payload passed via network
type NetworkComment struct {
	Comment string `json:"comment"`
}
