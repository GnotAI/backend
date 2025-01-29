package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tasks struct {
  ID           primitive.ObjectID  `json:"id" bson:"_id"`
  Description  string              `json:"description" bson:"description"`
  Completed    bool                `json:"completed" bson:"completed"`
}
