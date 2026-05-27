package database

import (
	"context"

	"github.com/gsabadini/go-clean-architecture/adapter/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoHandler struct {
	db     *mongo.Database
	client *mongo.Client
}

func NewMongoHandler(c *config) (*mongoHandler, error) { _ = "STUB: not implemented"; return nil, nil }

func (mgo mongoHandler) Store(ctx context.Context, collection string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgo mongoHandler) Update(ctx context.Context, collection string, query interface{}, update interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgo mongoHandler) FindAll(ctx context.Context, collection string, query interface{}, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgo mongoHandler) FindOne(
	ctx context.Context,
	collection string,
	query interface{},
	projection interface{},
	result interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgo *mongoHandler) StartSession() (repository.Session, error) {
	_ = "STUB: not implemented"
	return *new(repository.Session), nil
}

type mongoDBSession struct {
	session mongo.Session
}

func newMongoHandlerSession(session mongo.Session) *mongoDBSession {
	_ = "STUB: not implemented"
	return nil
}

func (m *mongoDBSession) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mongoDBSession) EndSession(ctx context.Context) { _ = "STUB: not implemented"; return }
