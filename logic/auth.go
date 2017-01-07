package main

import (
	"errors"
	"goim/libs/define"
	"strconv"
	"strings"
)

// developer could implement "Auth" interface for decide how get userId, or roomId
type Auther interface {
	Auth(token string) (userId int64, roomId int32, err error)
}

type DefaultAuther struct {
}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}

func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32, err error) {
	user_array := strings.Split(token, ",")
	if user_array == nil {
		err = errors.New("token is empty")
		return
	}
	if userId, err = strconv.ParseInt(user_array[0], 10, 64); err != nil {
		userId = 0
		roomId = define.NoRoom
		return
	}
	if len(user_array) == 2 {
		roomId, err = strconv.ParseInt(user_array[1], 10, 32)
	} else {
		roomId = define.NoRoom
	}
	//if userId, err = strconv.ParseInt(token, 10, 64); err != nil {
	//	userId = 0
	//	roomId = define.NoRoom
	//} else {
	//	roomId = 1 // only for debug
	//}
	return
}
