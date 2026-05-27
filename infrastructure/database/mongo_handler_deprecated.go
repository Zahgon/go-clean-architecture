package database

import (
	"context"

	mongo "gopkg.in/mgo.v2"
)

type mongoHandlerDeprecated struct {
	database *mongo.Database
	session  *mongo.Session
}

func NewMongoHandlerDeprecated(c *config) (*mongoHandlerDeprecated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mgo mongoHandlerDeprecated) Store(_ context.Context, collection string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgo mongoHandlerDeprecated) Update(_ context.Context, collection string, query interface{}, update interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgo mongoHandlerDeprecated) FindAll(_ context.Context, collection string, query interface{}, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgo mongoHandlerDeprecated) FindOne(
	_ context.Context,
	collection string,
	query interface{},
	selector interface{},
	result interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}
