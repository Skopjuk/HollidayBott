package weather_subscription

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Location struct {
	Longitude float64 `bson:"longitude"`
	Latitude  float64 `bson:"latitude"`
}

type Subscription struct {
	ID        primitive.ObjectID `bson:"_id"`
	ChatId    int64              `bson:"chat_id"`
	Username  string             `bson:"username"`
	SendAt    string             `bson:"send_at"`
	Location  Location           `bson:",inline"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
