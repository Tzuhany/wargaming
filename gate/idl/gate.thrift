namespace go gate

struct User {
    1: required string user_id,
    2: required string username,
}

struct RegisterReq {
    1: required string username,
    2: required string password,
}

struct RegisterResp {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required string token,
    4: required string connector_addr,
}

struct LoginReq {
    1: required string username,
    2: required string password,
}

struct LoginResp {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required string user_id
    4: required string token,
    5: required string connector_addr,
}

struct UserInfoReq {
    1: required i64 user_id,
}

struct UserInfoResp {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required User user,
}

service UserService {
    RegisterResp Register(1: RegisterReq req) (api.post="/user/register"),
    LoginResp Login(1: LoginReq req) (api.post="/user/login"),
    UserInfoResp UserInfo(1: UserInfoReq req) (api.post="/user/info"),
}