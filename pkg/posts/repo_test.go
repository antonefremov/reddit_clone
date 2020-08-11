package posts

import (
	"context"
	"errors"
	"golang-stepik-2020q2/6/99_hw/redditclone/pkg/user"
	"time"

	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

// go test -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockCollection := NewMockIMongoCollection(ctrl)
	mockSingleResult := NewMockIMongoSingleResult(ctrl)

	repo := &Repo{
		Collection: mockCollection,
	}

	postID := "12345" // primitive.NewObjectID()

	expectedPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author: user.User{
			Username: "userlogin",
			ID:       "userid",
		},
	}

	// positive outcome
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *expectedPost).
		Return(nil)

	res, err := repo.Get(postID)

	if !reflect.DeepEqual(res, expectedPost) {
		t.Errorf("bad result, expected %v, got %v", expectedPost, res)
	}
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}

	// search error
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.Any()).
		Return(errors.New("mocked-error"))

	_, err = repo.Get(postID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetByAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockCollection := NewMockIMongoCollection(ctrl)
	mockCursor := NewMockIMongoCursor(ctrl)

	repo := &Repo{
		Collection: mockCollection,
	}

	postID := "12345" // primitive.NewObjectID()
	username := "userlogin"

	expectedPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author: user.User{
			Username: username,
			ID:       "userid",
		},
	}

	// test the positive outcome
	mockCollection.EXPECT().
		Find(ctx, gomock.Any()).
		Return(mockCursor, nil)
	mockCursor.EXPECT().
		Next(ctx).MaxTimes(1).
		Return(true)
	mockCursor.EXPECT().
		Next(ctx).MaxTimes(1).
		Return(false)
	mockCursor.EXPECT().
		Close(ctx).
		Return(nil)
	mockCursor.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *expectedPost).
		Return(nil)

	res, err := repo.GetByAuthor(username)

	if !reflect.DeepEqual(res[0], expectedPost) {
		t.Errorf("bad result, expected %v, got %v", expectedPost, res)
	}
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}

	// cover the Collection.Find error
	mockCollection.EXPECT().
		Find(ctx, gomock.Any()).
		Return(nil, errors.New("mocked-error"))

	_, err = repo.GetByAuthor(username)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// cover the Decode error
	mockCollection.EXPECT().
		Find(ctx, gomock.Any()).
		Return(mockCursor, nil)
	mockCursor.EXPECT().
		Next(ctx).MaxTimes(1).
		Return(true)
	mockCursor.EXPECT().
		Next(ctx).MaxTimes(1).
		Return(false)
	mockCursor.EXPECT().
		Close(ctx).
		Return(nil)
	mockCursor.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *expectedPost).
		Return(errors.New("mocked-error"))

	_, err = repo.GetByAuthor(username)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockCollection := NewMockIMongoCollection(ctrl)
	mockCursor := NewMockIMongoCursor(ctrl)

	repo := &Repo{
		Collection: mockCollection,
	}

	postID := "12345" // primitive.NewObjectID()

	expectedPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author: user.User{
			Username: "userlogin",
			ID:       "userid",
		},
	}

	// test the positive outcome
	mockCollection.EXPECT().
		Find(ctx, gomock.Any()).
		Return(mockCursor, nil)
	mockCursor.EXPECT().
		Next(ctx).MaxTimes(1).
		Return(true)
	mockCursor.EXPECT().
		Next(ctx).MaxTimes(1).
		Return(false)
	mockCursor.EXPECT().
		Close(ctx).
		Return(nil)
	mockCursor.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *expectedPost).
		Return(nil)

	res, err := repo.All()

	if len(res) != 1 {
		t.Errorf("bad result, expected length 1, got %d", len(res))
	}
	if !reflect.DeepEqual(res[0], expectedPost) {
		t.Errorf("bad result, expected %v, got %v", expectedPost, res)
	}
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}
}

func TestListByCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockCollection := NewMockIMongoCollection(ctrl)
	mockCursor := NewMockIMongoCursor(ctrl)

	repo := &Repo{
		Collection: mockCollection,
	}

	postID := "12345" // primitive.NewObjectID()
	category := "programming"

	expectedPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: category,
		Author: user.User{
			Username: "userlogin",
			ID:       "userid",
		},
	}

	// test the positive outcome
	mockCollection.EXPECT().
		Find(ctx, gomock.Any()).
		Return(mockCursor, nil)
	mockCursor.EXPECT().
		Next(ctx).MaxTimes(1).
		Return(true)
	mockCursor.EXPECT().
		Next(ctx).MaxTimes(1).
		Return(false)
	mockCursor.EXPECT().
		Close(ctx).
		Return(nil)
	mockCursor.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *expectedPost).
		Return(nil)

	res, err := repo.ListByCategory(category)

	if !reflect.DeepEqual(res[0], expectedPost) {
		t.Errorf("bad result, expected %v, got %v", expectedPost, res)
	}
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}
}

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockCollection := NewMockIMongoCollection(ctrl)
	mockInsertResult := NewMockIMongoInsertOneResult(ctrl)

	repo := &Repo{
		Collection: mockCollection,
	}

	postID := "12345" // primitive.NewObjectID()

	expectedPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author: user.User{
			Username: "userlogin",
			ID:       "userid",
		},
	}

	// positive outcome
	mockCollection.EXPECT().
		InsertOne(ctx, gomock.Any()).
		Return(mockInsertResult, nil)

	res, err := repo.Add(expectedPost)

	if !reflect.DeepEqual(res, expectedPost) {
		t.Errorf("bad result, expected %v, got %v", expectedPost, res)
	}
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}

	// insert error
	mockCollection.EXPECT().
		InsertOne(ctx, gomock.Any()).
		Return(nil, errors.New("mocked-error"))

	_, err = repo.Add(expectedPost)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockCollection := NewMockIMongoCollection(ctrl)
	mockDeleteResult := NewMockIMongoDeleteResult(ctrl)

	repo := &Repo{
		Collection: mockCollection,
	}

	postID := "12345" // primitive.NewObjectID()

	expectedPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author: user.User{
			Username: "userlogin",
			ID:       "userid",
		},
	}

	// positive outcome (no err on delete operation)
	mockCollection.EXPECT().
		DeleteOne(ctx, gomock.Any()).
		Return(mockDeleteResult, nil)

	err := repo.Delete(postID, expectedPost.Author.ID)

	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}

	// delete error
	mockCollection.EXPECT().
		DeleteOne(ctx, gomock.Any()).
		Return(nil, errors.New("mocked-error"))

	err = repo.Delete(postID, expectedPost.Author.ID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestVote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockCollection := NewMockIMongoCollection(ctrl)
	mockSingleResult := NewMockIMongoSingleResult(ctrl)
	mockUpdateResult := NewMockIMongoUpdateResult(ctrl)

	repo := &Repo{
		Collection: mockCollection,
	}

	postID := "12345" // primitive.NewObjectID()
	username := "userlogin"
	vote := Vote{
		User: username,
		Vote: 1,
	}

	expectedPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author: user.User{
			Username: username,
			ID:       "userid",
		},
		Votes: []Vote{
			vote,
		},
		Score:            1,
		UpvotePercentage: 100,
	}

	// positive outcome
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *expectedPost).
		Return(nil)
	mockCollection.EXPECT().
		ReplaceOne(ctx, gomock.Any(), gomock.Any()).
		Return(mockUpdateResult, nil)

	res, err := repo.Vote(postID, vote)

	if !reflect.DeepEqual(res, expectedPost) {
		t.Errorf("bad result, expected %v, got %v", expectedPost, res)
	}
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}

	// another positive outcome when appending a vote (check the if condition inside the method)
	vote2 := Vote{
		User: "userlogin2",
		Vote: 1,
	}

	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *expectedPost).
		Return(nil)
	mockCollection.EXPECT().
		ReplaceOne(ctx, gomock.Any(), gomock.Any()).
		Return(mockUpdateResult, nil)

	res, err = repo.Vote(postID, vote2)

	expectedPost.Votes = append(expectedPost.Votes, vote2)
	expectedPost.Score = 2

	if !reflect.DeepEqual(res, expectedPost) {
		t.Errorf("bad result, expected %v, got %v", expectedPost, res)
	}
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}

	// repo.Get error inside the method
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.Any()).
		Return(errors.New("mocked-error"))

	_, err = repo.Vote(postID, vote)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// repo.ReplaceOne error inside the method
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *expectedPost).
		Return(nil)
	mockCollection.EXPECT().
		ReplaceOne(ctx, gomock.Any(), gomock.Any()).
		Return(nil, errors.New("mocked-error"))

	_, err = repo.Vote(postID, vote)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestUnvote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockCollection := NewMockIMongoCollection(ctrl)
	mockSingleResult := NewMockIMongoSingleResult(ctrl)
	mockUpdateResult := NewMockIMongoUpdateResult(ctrl)

	repo := &Repo{
		Collection: mockCollection,
	}

	postID := "12345" // primitive.NewObjectID()
	username := "userlogin"
	username2 := "userlogin2"
	vote := Vote{
		User: username,
		Vote: 1,
	}

	vote2Positive := Vote{
		User: username2,
		Vote: 1,
	}

	vote2Unvote := Vote{
		User: username2,
		Vote: 0,
	}

	expectedPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author: user.User{
			Username: username,
			ID:       "userid",
		},
		Votes: []Vote{
			vote,
			vote2Unvote,
		},
		Score:            1,
		UpvotePercentage: 50,
	}

	inputPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author: user.User{
			Username: username,
			ID:       "userid",
		},
		Votes: []Vote{
			vote,
			vote2Positive,
		},
		Score:            2,
		UpvotePercentage: 100,
	}

	// positive outcome
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *inputPost).
		Return(nil)
	mockCollection.EXPECT().
		ReplaceOne(ctx, gomock.Any(), gomock.Any()).
		Return(mockUpdateResult, nil)

	res, err := repo.Unvote(postID, username2)

	if !reflect.DeepEqual(res, expectedPost) {
		t.Errorf("bad result, expected %v, got %v", expectedPost, res)
	}
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}
}

func TestAddComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockCollection := NewMockIMongoCollection(ctrl)
	mockSingleResult := NewMockIMongoSingleResult(ctrl)
	mockUpdateResult := NewMockIMongoUpdateResult(ctrl)

	repo := &Repo{
		Collection: mockCollection,
	}

	postID := "12345" // primitive.NewObjectID()
	commentID := "1234"
	username := "userlogin"

	user := user.User{
		Username: username,
		ID:       "userid",
	}

	comment := Comment{
		ID:      commentID,
		Author:  user,
		Body:    "comment",
		Created: time.Now(),
	}

	inputPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author:   user,
		Comments: []Comment{},
	}

	expectedPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author:   user,
		Comments: []Comment{
			comment,
		},
	}

	// positive outcome
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.AssignableToTypeOf(inputPost)).
		SetArg(0, *inputPost).
		Return(nil)
	mockCollection.EXPECT().
		ReplaceOne(ctx, gomock.Any(), gomock.Any()).
		Return(mockUpdateResult, nil)

	res, err := repo.AddComment(postID, &comment)

	if len(res.Comments) != 1 ||
		!reflect.DeepEqual(res.Comments[0].Author, expectedPost.Comments[0].Author) ||
		res.Comments[0].Body != expectedPost.Comments[0].Body ||
		res.Comments[0].Created != expectedPost.Comments[0].Created {
		t.Errorf("bad result, expected %v, got %v", expectedPost, res)
	}
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}

	// repo.Get error inside the method
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.Any()).
		Return(errors.New("mocked-error"))

	_, err = repo.AddComment(postID, &comment)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// repo.ReplaceOne error inside the method
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *expectedPost).
		Return(nil)
	mockCollection.EXPECT().
		ReplaceOne(ctx, gomock.Any(), gomock.Any()).
		Return(nil, errors.New("mocked-error"))

	_, err = repo.AddComment(postID, &comment)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestDeleteComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockCollection := NewMockIMongoCollection(ctrl)
	mockSingleResult := NewMockIMongoSingleResult(ctrl)
	mockUpdateResult := NewMockIMongoUpdateResult(ctrl)

	repo := &Repo{
		Collection: mockCollection,
	}

	postID := "12345" // primitive.NewObjectID()
	commentID := "1234"
	username := "userlogin"

	user := user.User{
		Username: username,
		ID:       "userid",
	}

	comment := Comment{
		ID:      commentID,
		Author:  user,
		Body:    "comment",
		Created: time.Now(),
	}

	expectedPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author:   user,
		Comments: []Comment{},
	}

	inputPost := &Post{
		ID:       postID,
		Type:     "text",
		Category: "programming",
		Author:   user,
		Comments: []Comment{
			comment,
		},
	}

	// positive outcome
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.AssignableToTypeOf(inputPost)).
		SetArg(0, *inputPost).
		Return(nil)
	mockCollection.EXPECT().
		ReplaceOne(ctx, gomock.Any(), gomock.Any()).
		Return(mockUpdateResult, nil)

	res, err := repo.DeleteComment(postID, commentID)

	if !reflect.DeepEqual(res, expectedPost) {
		t.Errorf("bad result, expected %v, got %v", expectedPost, res)
	}
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}

	// repo.Get error inside the method
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.Any()).
		Return(errors.New("mocked-error"))

	_, err = repo.DeleteComment(postID, commentID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// repo.ReplaceOne error inside the method
	mockCollection.EXPECT().
		FindOne(ctx, gomock.Any()).
		Return(mockSingleResult)
	mockSingleResult.EXPECT().
		Decode(gomock.AssignableToTypeOf(expectedPost)).
		SetArg(0, *expectedPost).
		Return(nil)
	mockCollection.EXPECT().
		ReplaceOne(ctx, gomock.Any(), gomock.Any()).
		Return(nil, errors.New("mocked-error"))

	_, err = repo.DeleteComment(postID, commentID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
