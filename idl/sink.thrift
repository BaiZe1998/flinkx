namespace go sinkdemo

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct CreateSinkRequest {
    1:list<map<string,i64>> tables
    2:string time_stamp
}

struct CreateSinkResponse {
    1:BaseResp base_resp
}

service SinkService {
    CreateSinkResponse CreateSink(1:CreateSinkRequest req)
}