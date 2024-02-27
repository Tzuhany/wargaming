namespace go interact

struct BaseResp {
    1: required i64 code,
    2: required string msg,
}

// websocket

struct InteractReq{
    1: required i64 userId,
}

struct InteractResp{}

service InteractService{
    InteractResp Interact(1: InteractReq req)(api.get="/interact"),
}