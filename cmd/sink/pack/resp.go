package pack

import (
	"errors"
	"time"
	"word-count/cmd/sink/kitex_gen/sinkdemo"
	"word-count/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *sinkdemo.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *sinkdemo.BaseResp {
	return &sinkdemo.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
