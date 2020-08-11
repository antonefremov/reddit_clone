package posts

import (
	"context"
	"errors"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/ids"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Repo represents a repository with all the Posts in memory
type Repo struct {
	Collection IMongoCollection
}

var (
	// ErrNoPost is used to indicate that a post doesn't exist
	ErrNoPost = errors.New("Post not found")
)

// NewRepo creates a new Repository for Posts
func NewRepo(collection IMongoCollection) *Repo {
	return &Repo{
		Collection: collection,
	}
}

// Get returns a Post item by ID
func (repo *Repo) Get(objID string) (*Post, error) {
	post := &Post{}
	ctx := context.Background()
	err := repo.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// getByFilter returns posts according to the given filter
func (repo *Repo) getByFilter(filter interface{}) ([]*Post, error) {
	posts := []*Post{}
	ctx := context.Background()
	cur, err := repo.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Post
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &result)
	}
	return posts, nil
}

// All returns all the existing posts from the Repository
func (repo *Repo) All() ([]*Post, error) {
	filter := bson.M{}
	return repo.getByFilter(filter)
}

// ListByCategory returns posts filtered by a respective category
func (repo *Repo) ListByCategory(category string) ([]*Post, error) {
	filter := bson.M{"category": category}
	return repo.getByFilter(filter)
}

// GetByAuthor returns posts by their author
func (repo *Repo) GetByAuthor(login string) ([]*Post, error) {
	filter := bson.M{"author.username": login}
	return repo.getByFilter(filter)
}

// Add adds a new Post item into the Repository
func (repo *Repo) Add(post *Post) (*Post, error) {
	post.ID = ids.GenerateID() // primitive.NewObjectID()
	post.Votes = []Vote{
		{
			User: post.Author.ID,
			Vote: 1,
		},
	}

	post.Score = 1
	post.Views = 0

	post.Comments = make([]Comment, 0)
	post.Created = time.Now()
	post.UpvotePercentage = 100

	ctx := context.Background()
	_, err := repo.Collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// Delete removes an existing Post item from the Repository
func (repo *Repo) Delete(id string, userID string) error {
	filter := bson.M{"id": id}

	ctx := context.Background()
	_, err := repo.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return ErrNoPost
	}
	return nil
}

// Vote adds user's vote with either positive or negative value to a Post by Id
func (repo *Repo) Vote(id string, v Vote) (*Post, error) {
	filter := bson.M{"id": id}
	ctx := context.Background()
	post, err := repo.Get(id)
	if err != nil {
		return nil, ErrNoPost
	}
	flag := true
	for i, vote := range post.Votes {
		if vote.User == v.User {
			post.Votes[i].Vote = v.Vote
			recount(post)
			flag = false
			break
		}
	}
	if flag {
		post.Votes = append(post.Votes, v)
		recount(post)
	}

	_, err = repo.Collection.ReplaceOne(ctx, filter, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// Unvote removes user's vote from a Post by Id
func (repo *Repo) Unvote(id string, userID string) (*Post, error) {
	return repo.Vote(id, Vote{
		User: userID,
		Vote: 0,
	})
}

// // AddComment adds a new comment to a Post
func (repo *Repo) AddComment(postID string, comment *Comment) (*Post, error) {
	comment.ID = ids.GenerateID()
	filter := bson.M{"id": postID}
	ctx := context.Background()

	post, err := repo.Get(postID)
	if err != nil {
		return nil, ErrNoPost
	}

	post.Comments = append(post.Comments, *comment)

	_, err = repo.Collection.ReplaceOne(ctx, filter, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// DeleteComment removes an existing comment from a Post
func (repo *Repo) DeleteComment(postID string, commentID string) (*Post, error) {
	filter := bson.M{"id": postID}
	ctx := context.Background()
	post, err := repo.Get(postID)
	if err != nil {
		return nil, ErrNoPost
	}
	for i, c := range post.Comments {
		if c.ID == commentID {
			post.Comments[i] = post.Comments[len(post.Comments)-1]
			post.Comments = post.Comments[:len(post.Comments)-1]
			break
		}
	}

	_, err = repo.Collection.ReplaceOne(ctx, filter, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func recount(post *Post) {
	post.Score = 0
	up := 0
	for _, v := range post.Votes {
		post.Score += v.Vote
		if v.Vote == 1 {
			up++
		}
	}
	post.UpvotePercentage = uint8(up * 100 / len(post.Votes))
}
