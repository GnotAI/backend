package models

type Tasks struct {
  ID           int        `json:"id"`
  Description  string     `json:"description"`
  Completed    bool       `json:"completed"`
}
