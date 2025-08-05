package fun

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Uint64PtrToTime(ts *uint64) time.Time {
	if ts == nil {
		return time.Time{} // 返回零值（0001-01-01 00:00:00）
	}
	return time.UnixMilli(int64(*ts)) // 转换为 time.Time
}
func Int64ToUint64Ptr(v int64) *uint64 {
	u := uint64(v)
	return &u
}

func GetStringLocal(s string) *string {
	return &s
}

// int 转成 bool
func IntToBoolPtr(val int64) *bool {
	b := val == 1
	return &b
}

// 专门处理字符串指针转换
func NullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
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
