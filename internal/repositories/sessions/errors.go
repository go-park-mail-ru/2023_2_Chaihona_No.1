package sessions

import (
	"errors"
)

var (
	ErrNoSuchSession              = errors.New("no such session")
	ErrResultNotOk                = errors.New("result not ok")
	ErrRedis                      = errors.New("internal redis error")
	ErrRedisCantGetData           = errors.New("can't get data")
	ErrRedisCantUnpackSessionData = errors.New("can't unpack session data")
)
