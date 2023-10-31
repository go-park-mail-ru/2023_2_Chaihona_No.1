package sessions

import (
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"
)

func TestRegisterNewSessionRedis(t *testing.T) {
	conn, err := redis.DialURL("redis://@localhost:6379")

	require.Equal(t, nil, err)

	sessionStorage := CreateRedisSessionStorage(conn)

	err = sessionStorage.RegisterNewSession(Session{
		SessionID: "2",
		UserID:    2,
		TTL:       time.Now().Add(10 * time.Hour),
	})

	require.Equal(t, nil, err)
}

func TestCheckSessionRedis(t *testing.T) {
	conn, err := redis.DialURL("redis://@localhost:6379")

	require.Equal(t, nil, err)

	sessionStorage := CreateRedisSessionStorage(conn)

	s, ok := sessionStorage.CheckSession("2")

	require.Equal(t, true, ok)

	require.Equal(t, "2", s.SessionID)
}
