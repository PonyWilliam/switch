package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "1.15.143.79:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	_ = initClient()
}

func SetUser(phonenum string, ip string, code string) {
	rdb.Set(phonenum, code, time.Minute*5) //验证码有效期5分钟
	rdb.Set(ip, "1", time.Minute*1)        //同一ip一分钟内只能申请一个验证码
}

func CheckIp(ip string) bool {
	res, err := rdb.Get(ip).Result()
	if err != nil {
		return true //可以
	}
	if res == "1" {
		return false
	}
	return true
}

func GetUser(phonenum string) (string, error) {
	res, err := rdb.Get(phonenum).Result()
	if err != nil {
		return "", errors.New("已过期，请重新获取验证码")
	}
	return res, nil
}

func DelUser(phonenum string) {
	//不需要管，删掉就行了
	rdb.Del(phonenum)
}
