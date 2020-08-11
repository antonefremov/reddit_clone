package posts

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// mockgen command:
// mockgen -source=mongodb_abstract_layer.go -destination=mongodb_abstract_layer_mock.go -package=posts IMongoDatabase

type IMongoDatabase interface {
	Collection(name string) IMongoCollection
}

type IMongoCollection interface {
	Find(ctx context.Context, filter interface{}) (IMongoCursor, error)
	FindOne(ctx context.Context, filter interface{}) IMongoSingleResult
	InsertOne(ctx context.Context, item interface{}) (IMongoInsertOneResult, error)
	DeleteOne(ctx context.Context, filter interface{}) (IMongoDeleteResult, error)
	ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}) (IMongoUpdateResult, error)
}

type IMongoSingleResult interface {
	Decode(v interface{}) error
}

type IMongoCursor interface {
	Close(context.Context) error
	Next(context.Context) bool
	Decode(interface{}) error
}

type IMongoInsertOneResult interface {
	// Decode(v interface{}) error
}

type IMongoDeleteResult interface {
	// Decode(v interface{}) error
}

type IMongoUpdateResult interface {
	// Decode(v interface{}) error
}

type MongoCollection struct {
	Сoll *mongo.Collection
}

type MongoSingleResult struct {
	sr *mongo.SingleResult
}

type MongoCursor struct {
	cur *mongo.Cursor
}

type MongoInsertOneResult struct {
	ir *mongo.InsertOneResult
}

type MongoDeleteResult struct {
	dr *mongo.DeleteResult
}

type MongoUpdateResult struct {
	ur *mongo.UpdateResult
}

// func (mu *MongoUpdateResult) Decode(val interface{}) error {
// 	return mu.ur. //Decode(val)
// }

func (msr *MongoSingleResult) Decode(v interface{}) error {
	return msr.sr.Decode(v)
}

// func (mior *MongoInsertOneResult) Decode(v interface{}) error {
// 	return mior.Decode(v)
// }

func (mc *MongoCursor) Close(ctx context.Context) error {
	return mc.cur.Close(ctx)
}

func (mc *MongoCursor) Next(ctx context.Context) bool {
	return mc.cur.Next(ctx)
}

func (mc *MongoCursor) Decode(val interface{}) error {
	return mc.cur.Decode(val)
}

func (mc *MongoCollection) Find(ctx context.Context, filter interface{}) (IMongoCursor, error) {
	cursorResult, err := mc.Сoll.Find(ctx, filter)
	return &MongoCursor{cur: cursorResult}, err
}

func (mc *MongoCollection) FindOne(ctx context.Context, filter interface{}) IMongoSingleResult {
	singleResult := mc.Сoll.FindOne(ctx, filter)
	return &MongoSingleResult{sr: singleResult}
}

func (mc *MongoCollection) InsertOne(ctx context.Context, item interface{}) (IMongoInsertOneResult, error) {
	singleResult, err := mc.Сoll.InsertOne(ctx, item)
	return &MongoInsertOneResult{ir: singleResult}, err
}

func (mc *MongoCollection) DeleteOne(ctx context.Context, item interface{}) (IMongoDeleteResult, error) {
	deleteResult, err := mc.Сoll.DeleteOne(ctx, item)
	return &MongoDeleteResult{dr: deleteResult}, err
}

func (mc *MongoCollection) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}) (IMongoUpdateResult, error) {
	updateResult, err := mc.Сoll.ReplaceOne(ctx, filter, replacement)
	return &MongoUpdateResult{ur: updateResult}, err
}
