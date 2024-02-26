namespace go interact

// websocket

struct InteractReq{}

struct InteractResp{}

service InteractService{
    InteractResp Interact(1: InteractReq req)(api.get="/interact"),
}