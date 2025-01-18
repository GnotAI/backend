package models

type Powerups struct {
  ID        int     `json:"id"`
  Name      string  `json:"name"`
  Duration  int     `json:"duration"`
  Active    bool    `json:"active"`
}
