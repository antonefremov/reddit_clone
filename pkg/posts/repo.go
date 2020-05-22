package posts

import (
	"encoding/json"
	"fmt"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/ids"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

// Repo represents a repository with all the Posts in memory
type Repo struct {
	posts []*Post
	s     sync.RWMutex
}

// NewRepo creates a new Repository for Posts
func NewRepo() *Repo {
	return &Repo{
		posts: loadDataFromFile("./../../assets/posts.json"),
	}
}

// All returns all the existing posts from the Repository
func (repo *Repo) All() ([]*Post, error) {
	repo.s.RLock()
	defer repo.s.RUnlock()

	return repo.posts, nil
}

// ListByCategory returns posts filtered by a respective category
func (repo *Repo) ListByCategory(category string) ([]*Post, error) {
	repo.s.RLock()
	defer repo.s.RUnlock()

	result := make([]*Post, 0, len(repo.posts))

	for _, item := range repo.posts {
		if item.Category == category {
			result = append(result, item)
		}
	}

	return result, nil
}

// Get returns a Post from the Repository by Id
func (repo *Repo) Get(id string) (*Post, error) {
	repo.s.RLock()
	defer repo.s.RUnlock()

	for _, item := range repo.posts {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, nil
}

// Add adds a new Post item into the Repository
func (repo *Repo) Add(post *Post) (*Post, error) {
	repo.s.Lock()
	defer repo.s.Unlock()

	vote := Vote{User: post.Author.ID, Vote: 1}
	votes := make([]Vote, 1)
	post.Votes = append(votes, vote)

	post.Score = 1
	post.Views = 0

	post.Comments = make([]Comment, 0)
	post.Created = time.Now()
	post.UpvotePercentage = 100
	post.ID = ids.GenerateID()

	repo.posts = append(repo.posts, post)
	return post, nil
}

// Delete removes an existing Post item from the Repository
func (repo *Repo) Delete(id string, userID string) error {
	repo.s.Lock()
	defer repo.s.Unlock()

	length := len(repo.posts)
	for i, item := range repo.posts {
		if item.ID == id {
			if item.Author.ID == userID {
				repo.posts[i] = repo.posts[length-1]
				repo.posts = repo.posts[:length-1]
				return nil
			}

			return fmt.Errorf("Access to the item is not allowed")
		}
	}

	return fmt.Errorf("Item not found")
}

// Vote adds user's vote with either positive or negative value to a Post by Id
func (repo *Repo) Vote(id string, vote Vote) (*Post, error) {
	repo.s.Lock()
	defer repo.s.Unlock()

	for _, item := range repo.posts {
		if item.ID == id {
			var counter int
			var isUserVoted bool
			for _, itemVote := range item.Votes {
				counter += itemVote.Vote
				if itemVote.User == vote.User {
					isUserVoted = true
				}
			}
			if !isUserVoted {
				item.Votes = append(item.Votes, vote)
				counter += vote.Vote
			}
			item.Score = counter
			return item, nil
		}
	}
	return nil, nil
}

// Unvote removes user's vote from a Post by Id
func (repo *Repo) Unvote(id string, userID string) (*Post, error) {
	repo.s.Lock()
	defer repo.s.Unlock()

	for _, item := range repo.posts {
		if item.ID == id {
			var counter int
			length := len(item.Votes)
			for i, vote := range item.Votes {
				if vote.User == userID {
					// remove vote item created by this user
					item.Votes[i] = item.Votes[length-1]
					item.Votes = item.Votes[:length-1]
				} else {
					counter += vote.Vote
				}
			}
			item.Score = counter
			return item, nil
		}
	}
	return nil, nil
}

// AddComment adds a new comment to a Post
func (repo *Repo) AddComment(postID string, comment *Comment) (*Post, error) {
	repo.s.Lock()
	defer repo.s.Unlock()

	comment.ID = ids.GenerateID()

	for _, item := range repo.posts {
		if item.ID == postID {
			item.Comments = append(item.Comments, *comment)
			return item, nil
		}
	}
	return nil, nil
}

// DeleteComment removes an existing comment from a Post
func (repo *Repo) DeleteComment(postID string, commentID string) (*Post, error) {
	repo.s.Lock()
	defer repo.s.Unlock()

	for _, item := range repo.posts {
		if item.ID == postID {
			length := len(item.Comments)
			for i, comment := range item.Comments {
				if comment.ID == commentID {
					item.Comments[i] = item.Comments[length-1]
					item.Comments = item.Comments[:length-1]
					return item, nil
				}
			}
		}
	}
	return nil, nil
}

func loadDataFromFile(filePath string) []*Post {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var posts []*Post

	log.Println("Reading the file, trying to serialise the object")

	err = json.Unmarshal([]byte(byteValue), &posts)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return posts
}
