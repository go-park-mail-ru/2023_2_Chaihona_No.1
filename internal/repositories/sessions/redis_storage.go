package sessions

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

const longForm = "Jan 2, 2006 at 3:04pm (MST)"

type SessionRedisStorage struct {
	// redisConn redis.Conn
	UnimplementedAuthCheckerServer
	redisPool *redis.Pool
}

type RedisManager struct {
	storage *SessionRedisStorage
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

func CreateRedisManager(pool *redis.Pool) *RedisManager {
	return &RedisManager{
		storage: CreateRedisSessionStorage(pool),
	}
}

func (storage *SessionRedisStorage) RegisterNewSession(session SessionOld) error {
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

func (storage *SessionRedisStorage) RegisterNewSessionCtx(ctx context.Context, session *Session) (*Nothing, error) {
	sessionSerialized, _ := json.Marshal(session)
	conn := storage.redisPool.Get()
	defer conn.Close()

	// TTL := session.TTL.Hour()*int(time.Hour) + session.TTL.Minute()*int(time.Minute) + session.TTL.Second()*int(time.Second)
	result, err := redis.String(conn.Do("SET", session.SessionId, sessionSerialized, "EX", session.Ttl))
	if err != nil {
		log.Println(err)
		return &Nothing{}, err
	}

	if result != "OK" {
		log.Println(ErrNoSuchSession.Error())
		return &Nothing{}, ErrResultNotOk
	}

	return &Nothing{}, nil
}

func (manager *RedisManager) RegisterNewSessionCtxWrapper(ctx context.Context, session SessionOld) error {
	_, err := manager.storage.RegisterNewSessionCtx(ctx, &Session{
		SessionId: session.SessionID,
		UserId: session.UserID,
		Ttl: session.TTL.String(),
	})

	return err
}


func (storage *SessionRedisStorage) CheckSession(sessionId string) (*SessionOld, bool) {
	conn := storage.redisPool.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", sessionId))

	if err != nil {
		// err = errors.Join(ErrRedisCantGetData, err)
		log.Println(err.Error())

		return nil, false
	}

	sess := &SessionOld{}

	err = json.Unmarshal(data, sess)
	if err != nil {
		// err = errors.Join(ErrRedisCantUnpackSessionData, err)
		log.Println(err.Error())

		return nil, false
	}

	return sess, true
}

func (storage *SessionRedisStorage) CheckSessionCtx(ctx context.Context, sessionId *SessionID) (*CheckSession, error) {
	conn := storage.redisPool.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", sessionId.Id))

	if err != nil {
		// err = errors.Join(ErrRedisCantGetData, err)
		log.Println(err.Error())

		return &CheckSession{
			Exists: false,
		}, err
	}

	sess := &SessionOld{}

	err = json.Unmarshal(data, sess)
	if err != nil {
		// err = errors.Join(ErrRedisCantUnpackSessionData, err)
		log.Println(err.Error())

		return &CheckSession{
			Exists: false,
		}, err
	}

	return &CheckSession{
		SessionId: sess.SessionID,
		UserId: sess.UserID,
		Ttl: sess.TTL.String(),
	}, nil
}

func (manager *RedisManager) CheckSessionCtxWrapper(ctx context.Context, sessionID string) (*SessionOld, bool) {
	session, err := manager.storage.CheckSessionCtx(ctx, &SessionID{
		Id: sessionID,
	})

	if err != nil {
		return nil, false
	}

	ttl, err := time.Parse(longForm, session.Ttl) 
	return &SessionOld{
		SessionID: session.SessionId,
		UserID: session.UserId,
		TTL: ttl,
	}, true
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

func (storage *SessionRedisStorage) DeleteSessionCtx(ctx context.Context, sessionId *SessionID) (*Nothing, error) {
	conn := storage.redisPool.Get()
	defer conn.Close()

	_, err := redis.Int(conn.Do("DEL", sessionId.Id))

	if err != nil {
		// err = errors.Join(ErrRedis, err)
		log.Println(err.Error())

		return &Nothing{}, err
	}

	return &Nothing{}, nil
}

func (manager *RedisManager) DeleteSessionCtxWrapper(ctx context.Context, sessionId string) error {
	_, err := manager.storage.DeleteSessionCtx(ctx, &SessionID{
		Id: sessionId,
	})

	return err
}
