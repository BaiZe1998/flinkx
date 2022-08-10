namespace go keybydemo

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct CreateKeybyRequest {
    1:list<string> content
    2:i64 value
    3:string time_stamp
}

struct CreateKeybyResponse {
    1:BaseResp base_resp
}

service KeybyService {
    CreateKeybyResponse CreateKeyby(1:CreateKeybyRequest req)
}