namespace go game

// model

struct BaseResp {
    1: required i64 code,
    2: required string msg,
}

struct Point {
    1: required double lat,
    2: required double lng,
}

// base

struct MatchReq {
    1: required i64 user_id,
}

struct MatchResp {
    1: required BaseResp base,
    2: required i64 matched_user_id,
}

struct MoveReq {
    1: required list<i64> cells,
    2: required list<i64> obstacle,
    3: required i64 target_cell,
    4: required i64 origin_cell,
}

struct MoveResp {
    1: required BaseResp base,
    2: required list<i64> path,
}

service GameService {
    MatchResp Match(1: MatchReq req),
    MoveResp Move(1: MoveReq req),
}