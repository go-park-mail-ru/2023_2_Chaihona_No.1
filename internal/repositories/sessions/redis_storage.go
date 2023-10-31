package sessions

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type SessionRedisStorage struct {
	redisConn redis.Conn
}

func CreateRedisSessionStorage(conn redis.Conn) *SessionRedisStorage {
	return &SessionRedisStorage{
		redisConn: conn,
	}
}

func (storage *SessionRedisStorage) RegisterNewSession(session Session) error {
	sessionSerialized, _ := json.Marshal(session)

	TTL := session.TTL.Hour()*int(time.Hour) + session.TTL.Minute()*int(time.Minute) + session.TTL.Second()*int(time.Second)
	result, err := redis.String(storage.redisConn.Do("SET", session.SessionID, sessionSerialized, "EX", TTL))
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
	data, err := redis.Bytes(storage.redisConn.Do("GET", sessionId))

	if err != nil {
		err = errors.Join(ErrRedisCantGetData, err)
		log.Println(err.Error())

		return nil, false
	}

	sess := &Session{}

	err = json.Unmarshal(data, sess)
	if err != nil {
		err = errors.Join(ErrRedisCantUnpackSessionData, err)
		log.Println(err.Error())

		return nil, false
	}

	return sess, true
}

func (storage *SessionRedisStorage) DeleteSession(sessionID string) error {
	_, err := redis.Int(storage.redisConn.Do("DEL", sessionID))

	if err != nil {
		err = errors.Join(ErrRedis, err)
		log.Println(err.Error())

		return err
	}

	return nil
}

func (storage *SessionRedisStorage) GetSessions() ([]Session, error) {
	// sessions := make([]Session, 0, len(storage.Sessions))

	// for _, session := range storage.Sessions {
	// 	sessions = append(sessions, session)
	// }

	return nil, nil
}
