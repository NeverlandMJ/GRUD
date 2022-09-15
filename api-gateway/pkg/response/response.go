package response

import "github.com/NeverlnadMJ/GRUD/api-gateway/protos/grudpb"

type DataResponse struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func ToDataResponse(protoPost *grudpb.Data) DataResponse {
	return DataResponse{
		ID:     int(protoPost.Id),
		UserID: int(protoPost.UserId),
		Title:  protoPost.Title,
		Body:   protoPost.Body,
	}
}
