package pack

import (
	"errors"
	"time"
	"word-count/kitex_gen/mapdemo"
	"word-count/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *mapdemo.BaseResp {
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

func baseResp(err errno.ErrNo) *mapdemo.BaseResp {
	return &mapdemo.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
