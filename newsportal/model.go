package newsportal

// NewsPortal - struct
type NewsPortal struct {
	ID    string `json:"id"  bson:"_id"`
	Title string `json:"title" bson:"title"`
	// Date    time.Time `json:"date" bson:"date"`
	Content string `json:"content" bson:"content"`
}
