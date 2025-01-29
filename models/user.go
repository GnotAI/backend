package models

type Users struct {
  ID        int        `json:"id" bson:"id"`
  Username  string     `json:"username" bson:"username"`
  Password  string     `json:"password" bson:"password"`
}
