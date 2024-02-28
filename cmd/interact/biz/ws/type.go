package ws

const (
	MatchAction = iota
	MoveAction
)

type MatchReq struct {
	UserId int64 `json:"userId"`
}

type MatchResp struct {
	MatchedUserId int64 `json:"matchedUserId"`
}

type MoveData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
