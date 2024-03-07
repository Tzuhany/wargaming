namespace go user

struct BaseResp {
    1: required string code,
    2: required string msg,
}

struct User {
    1: required string uid,
    2: required string username,
}

struct RegisterReq {
    1: required string username,
    2: required string password,
}

struct RegisterResp {
    1: required BaseResp base,
    2: required string uid,
    3: required string token,
    4: required string connector_addr
}

struct LoginReq {
    1: required string username,
    2: required string password,
}

struct LoginResp {
    1: required BaseResp base,
    2: required string uid,
    3: required string token
    4: required string connector_addr
}

service UserService {
    RegisterResp Register(1: RegisterReq req),
    LoginResp Login(1: LoginReq req),
}



