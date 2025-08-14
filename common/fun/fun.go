package fun

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var RedisPrefix = "go-zero-fast"

func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func Int64ToUint64Ptr(v int64) *uint64 {
	u := uint64(v)
	return &u
}

func GetStringLocal(s string) *string {
	return &s
}

// int 转成 bool
func IntToBool(n int64) bool {
	return n != 0
}

// HashPassword 使用 bcrypt 对密码进行加密
func HashPassword(password string) (string, error) {
	// GenerateFromPassword 的第二个参数是 cost 值，范围从 4 到 31
	// 值越大，加密过程越安全但也越耗时
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword 验证密码是否匹配哈希值
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormatTimestampToDate(timestamp uint64) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
