package dal

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UID        string             `bson:"uid" json:"uid"`
	Username   string             `bson:"username" json:"username"`
	Password   string             `bson:"password" json:"password"`
	CreateTime time.Time          `bson:"create_time" json:"create_time"`
	Avatar     string             `bson:"avatar" json:"avatar"` // 头像
}

func FindUserByUID(ctx context.Context, uid string) (*User, error) {
	db := MongoDB.Collection("user")
	singleResult := db.FindOne(ctx, bson.D{
		{"uid", uid},
	})
	user := new(User)
	err := singleResult.Decode(user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func FindUserByUsername(ctx context.Context, username string) (*User, error) {
	db := MongoDB.Collection("user")
	singleResult := db.FindOne(ctx, bson.D{
		{"username", username},
	})
	user := new(User)
	err := singleResult.Decode(user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func Insert(ctx context.Context, user *User) error {
	db := MongoDB.Collection("user")
	_, err := db.InsertOne(ctx, user)
	return err
}
