package service

import (
	"github.com/gomodule/redigo/redis"
	"github.com/astaxie/beego/logs"
	"doc-lixiang/tool"
)

//获取redis数据 某个key
func  RedisGet(key string) (value string, err error){
	redisInit := tool.Redisinit()

	// 用完后将连接放回连接池
	defer redisInit.Close()
	conn := redisInit.Get()
	value, err = redis.String(conn.Do("GET", key))
	if err != nil{
		return "", err
	}
	return value, nil
}

func RedisLpush(key string, value string) (flag bool, err error){

	redisInit := tool.Redisinit()

	// 用完后将连接放回连接池
	defer redisInit.Close()
	conn := redisInit.Get()
	_, err = conn.Do("lpush", key,value)
	logs.Info("type-RedisLpush-key-"+key+"-value-"+value)

	if err != nil{
		return false, err
	}
	return true, nil
}

func RedisdDecrby(key string, value string) (flag bool, err error){

	redisInit := tool.Redisinit()

	// 用完后将连接放回连接池
	defer redisInit.Close()
	conn := redisInit.Get()
	_, err = conn.Do("decrby", key,value)
	if err != nil{
		return false, err
	}
	return true, nil
}


func RedisIncrby(key string, value string) (flag bool, err error) {

	redisInit := tool.Redisinit()

	// 用完后将连接放回连接池
	defer redisInit.Close()
	conn := redisInit.Get()
	_, err = conn.Do("incrby", key, value)
	if err != nil {
		return false, err
	}
	return true, nil

}