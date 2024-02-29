namespace go game

// model

struct BaseResp {
    1: required i64 code,
    2: required string msg,
}

// base

struct MatchReq {
    1: required i64 user_id,
}

struct MatchResp {
    1: required BaseResp base,
    2: required i64 matched_user_id,
}

struct FindPathReq {
    1: required list<i64> cells,
    2: required list<i64> obstacle,
    3: required i64 target_cell,
    4: required i64 origin_cell,
}

struct FindPathResp {
    1: required BaseResp base,
    2: required list<i64> path,
}

struct MoveReq {
    1: required i64 user_id,
    2: required i64 operator_id,
    3: required i64 new_cell,
}

struct MoveResp {
    1: required BaseResp base,
    2: required map<i64, i64> in_view,
}

service GameService {
    MatchResp Match(1: MatchReq req),
    FindPathResp Move(1: FindPathReq req),
}