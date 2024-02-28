package ws

const (
	MatchAction = iota
	MoveAction
)

type MatchReq struct {
	UserID int64 `json:"user_id"`
}

type MatchResp struct {
	MatchedUserID int64 `json:"matched_user_id"`
}

type MoveData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
