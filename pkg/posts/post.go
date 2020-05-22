package posts

import (
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/users"
	"time"
)

// Post struct which contains full information about posts
type Post struct {
	ID               string     `json:"id"` // uint32
	Title            string     `json:"title"`
	Url              string     `json:"url"`
	Author           users.User `json:"author"`
	Category         string     `json:"category"`
	Score            int        `json:"score"`
	Votes            []Vote     `json:"votes"`
	Comments         []Comment  `json:"comments"`
	Created          time.Time  `json:"created"`
	Views            int        `json:"views"`
	Type             string     `json:"type"`
	Text             string     `json:"text"`
	UpvotePercentage uint8      `json:"upvotePercentage"`
}

// Vote counts votes from users
type Vote struct {
	User string `json:"user"`
	Vote int    `json:"vote"`
}

// Comment object
type Comment struct {
	ID      string     `json:"id"`
	Author  users.User `json:"author"`
	Body    string     `json:"body"`
	Created time.Time  `json:"created"`
}

// NetworkComment represents json payload passed via network
type NetworkComment struct {
	Comment string `json:"comment"`
}
