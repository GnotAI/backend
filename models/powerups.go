package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Powerups struct {
  ID        primitive.ObjectID     `json:"id" bson:"id"`
  Name      string                 `json:"name" bson:"name"`
  Duration  int                    `json:"duration" bson:"duration"`
  Active    bool                   `json:"active" bson:"active"`
}
