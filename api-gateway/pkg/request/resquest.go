package request

type UpdatePostRequest struct {
	PostID int    `json:"post_id"`
	New    string `json:"new"`
}
