package response

import "time"

type ResponseUser struct {
	Name string    `json:"name"`
	Age  int       `json:"age"`
	Now  time.Time `json:"datetime"`
}
