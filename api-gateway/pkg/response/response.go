package response

import "github.com/NeverlnadMJ/GRUD/api-gateway/protos/grudpb"

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func ToDataResponse(protoPost *grudpb.Post) Post {
	return Post{
		ID:     int(protoPost.Id),
		UserID: int(protoPost.UserId),
		Title:  protoPost.Title,
		Body:   protoPost.Body,
	}
}
