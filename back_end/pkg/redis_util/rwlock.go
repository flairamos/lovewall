package redisutil

import (
	"errors"
	"fmt"
	"shoollove/pkg/tool"
)

func Lock(name string) (error, *string) {
	apply, err := Client.Cmd("GET", name)
	if apply != nil || err != nil {
		return errors.New("locking"), nil
	}
	uid := tool.Uuid()
	fmt.Printf("name: %v, uid: %v \n", name, uid)
	_, err = Client.Cmd("SET", name, uid, "EX", 15)
	if err != nil {
		Client.Cmd("DEL", name)
		return errors.New("redis set err"), nil
	}
	return nil, &uid
}

func LockEX(name string, ex int) (error, *string) {
	apply, err := Client.Cmd("GET", name)
	if apply != nil || err != nil {
		return errors.New("locking"), nil
	}
	uid := tool.Uuid()
	// fmt.Printf("name: %v, uid: %v \n", name, uid)
	_, err = Client.Cmd("SET", name, uid, "EX", ex)
	if err != nil {
		Client.Cmd("DEL", name)
		return errors.New("redis set err"), nil
	}
	return nil, &uid
}

func Unlock(name string, uid string) error {
	reply, err := Client.Cmd("GET", name)
	if reply == nil && err == nil {
		return nil
	}
	s := string(reply.([]byte))
	if s != uid {
		Client.Cmd("DEL", name)
		return errors.New("can not unlock")
	}
	_, err = Client.Cmd("DEL", name)
	if err != nil {
		Client.Cmd("DEL", name)
		return err
	}
	return nil
}
