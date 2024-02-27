package ws

import "wargaming/cmd/interact/biz/model/interact"

const (
	MatchAction = iota
	MoveAction
)

// Base

type Data struct {
	Action int                `json:"action"`
	Data   interface{}        `json:"Data"`
	Base   *interact.BaseResp `json:"base"`
}

// Ws Trans

type MatchData struct {
	UserId int64 `json:"matchedUserId"`
}

type MoveData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
