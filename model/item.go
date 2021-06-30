package model

type Item struct {
	// omit empty id field to protect against adding an empty ID when updating/inserting
	Id          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
