namespace go mapdemo

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct CreateMapRequest {
    1:string content
}

struct CreateMapResponse {
    1:BaseResp base_resp
}

service MapService {
    CreateMapResponse CreateMap(1:CreateMapRequest req)
}