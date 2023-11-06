package sessions

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type SessionRedisStorage struct {
	// redisConn redis.Conn
	redisPool *redis.Pool
}

const (
	maxIdles   = 10
	timeOutSec = 240
)

func NewPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdles,
		IdleTimeout: timeOutSec * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func CreateRedisSessionStorage(pool *redis.Pool) *SessionRedisStorage {
	return &SessionRedisStorage{
		redisPool: pool,
	}
}

func (storage *SessionRedisStorage) RegisterNewSession(session Session) error {
	sessionSerialized, _ := json.Marshal(session)
	conn := storage.redisPool.Get()
	defer conn.Close()

	TTL := session.TTL.Hour()*int(time.Hour) + session.TTL.Minute()*int(time.Minute) + session.TTL.Second()*int(time.Second)
	result, err := redis.String(conn.Do("SET", session.SessionID, sessionSerialized, "EX", TTL))
	if err != nil {
		log.Println(err)
		return err
	}

	if result != "OK" {
		log.Println(ErrNoSuchSession.Error())
		return ErrResultNotOk
	}

	return nil
}

func (storage *SessionRedisStorage) CheckSession(sessionId string) (*Session, bool) {
	conn := storage.redisPool.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", sessionId))

	if err != nil {
		// err = errors.Join(ErrRedisCantGetData, err)
		log.Println(err.Error())

		return nil, false
	}

	sess := &Session{}

	err = json.Unmarshal(data, sess)
	if err != nil {
		// err = errors.Join(ErrRedisCantUnpackSessionData, err)
		log.Println(err.Error())

		return nil, false
	}

	return sess, true
}

func (storage *SessionRedisStorage) DeleteSession(sessionID string) error {
	conn := storage.redisPool.Get()
	defer conn.Close()

	_, err := redis.Int(conn.Do("DEL", sessionID))

	if err != nil {
		// err = errors.Join(ErrRedis, err)
		log.Println(err.Error())

		return err
	}

	return nil
}
