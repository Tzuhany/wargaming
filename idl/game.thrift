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
    1: required i64 userId,
}

struct MatchResp {
    1: required BaseResp base,
    2: required i64 matchedUserId,
}

struct MoveReq {
    1: required Point orginPos,         // 原始位置
    2: required Point targetPos,        // 目标位置
    3: optional list<Point> obstacle,   // 障碍物位置
    4: required list<Point> corner,     // 四角点坐标
}

struct MoveResp {
    1: required BaseResp base,
    2: required list<Point> path,
}

service GameService {
    MatchResp Match(1: MatchReq req),
    MoveResp Move(1: MoveReq req),
}