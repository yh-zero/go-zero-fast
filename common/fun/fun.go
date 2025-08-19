package fun

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"reflect"
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

// 目前用于更新  针对更新数据的时候 可以更新一部分数据
func UpdateFieldsByReflect(src interface{}) map[string]interface{} {
	updateMap := make(map[string]interface{})
	v := reflect.ValueOf(src).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			// 获取原始驼峰命名字段名
			originalName := t.Field(i).Name
			// 转换为蛇形命名
			snakeName := camelToSnake(originalName)
			// 使用转换后的名称作为键
			updateMap[snakeName] = field.Elem().Interface()
		}
	}
	return updateMap
}

// camelToSnake 将驼峰命名转换为蛇形命名
func camelToSnake(s string) string {
	var result []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			// 如果是大写字母且不是第一个字符，在前面加下划线
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, c+32) // 转换为小写
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}

// BuildUpdateRequest 构建动态更新请求  -- 针对api层
// target: 目标结构体指针(如 &pb.RoleUpdateRequest{})
// source: 源结构体(如 req *types.RoleInfo)
func BuildUpdateRequest(target interface{}, source interface{}) error {
	targetVal := reflect.ValueOf(target).Elem()
	sourceVal := reflect.ValueOf(source).Elem()

	if srcID := sourceVal.FieldByName("Id"); srcID.IsValid() && !isZero(srcID) {
		if targetID := targetVal.FieldByName("Id"); targetID.IsValid() && targetID.CanSet() {
			targetID.SetUint(srcID.Uint())
		}
	}

	// 遍历源结构体字段
	for i := 0; i < sourceVal.NumField(); i++ {
		field := sourceVal.Type().Field(i)
		fieldName := field.Name
		fieldValue := sourceVal.Field(i)

		// 检查字段是否为零值
		if !isZero(fieldValue) {
			// 在目标结构体中查找对应字段
			targetField := targetVal.FieldByName(fieldName)
			if targetField.IsValid() && targetField.CanSet() {
				// 创建指针并赋值
				ptr := reflect.New(fieldValue.Type())
				ptr.Elem().Set(fieldValue)
				targetField.Set(ptr)
			}
		}
	}

	return nil
}

// isZero 检查值是否为零值
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return v.IsNil()
	default:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	}
}
