package session

import (
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

type Session struct {
	Token     string    `json:"token"`
	UserID    string    `json:"user_id"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Config struct {
	secretKey  string
	expiration time.Duration
}

func NewSession(secretKey string, expiration time.Duration) *Config {
	return &Config{
		secretKey:  secretKey,
		expiration: expiration,
	}
}

func GetSessionConfig() *Config {
	secretKey := os.Getenv("JWT_SECRET")
	expiration := os.Getenv("JWT_EXPIRATION")
	duration, err := time.ParseDuration(expiration)
	if err != nil {
		duration = constants.DefaultSessionExpiration
	}

	return NewSession(secretKey, duration)
}

func (s *Config) GenerateToken(userID string) (session Session, err error) {
	expiredAt := time.Now().Add(s.expiration)
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expiredAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return session, err
	}

	session = Session{
		Token:     tokenString,
		UserID:    userID,
		ExpiredAt: expiredAt,
	}

	return session, nil
}

func (s *Config) ParseToken(tokenStr string) (*Session, error) {
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	session := castToSessionToken(claims)
	return session, nil
}

func castToSessionToken(claim map[string]interface{}) (session *Session) {
	var (
		userID     string
		expiration time.Time
	)

	if val, ok := claim["user_id"].(string); ok {
		userID = val
	}
	if val, ok := claim["exp"].(float64); ok {
		expiration = time.Unix(int64(val), 0)
	}

	session = &Session{
		UserID:    userID,
		ExpiredAt: expiration,
	}
	return
}
