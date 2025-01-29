package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
  ID        primitive.ObjectID      `json:"id" bson:"_id"`
  Username  string                  `json:"username" bson:"username"`
  Password  string                  `json:"password" bson:"password"`
  Status    bool                    `json:"status" bson:"status"`
}
