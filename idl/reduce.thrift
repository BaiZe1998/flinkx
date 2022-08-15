namespace go reducedemo

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct CreateReduceRequest {
    1:list<Tuple> content
    2:string time_stamp
}

struct Tuple {
    1:string key
    2:i64 value
    3:string table
}

struct CreateReduceResponse {
    1:BaseResp base_resp
}

service ReduceService {
    CreateReduceResponse CreateReduce(1:CreateReduceRequest req)
}