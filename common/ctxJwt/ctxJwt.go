package ctxJwt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"reflect"
	"strconv"
)

var CtxKeyJwtData = "jwtData"

type JWTData struct {
	RoleIds      []uint64 `map:"RoleIds"`      // 多个角色id
	RoleId       uint64   `map:"RoleId"`       // 当前角色id
	DepartmentId uint64   `map:"DepartmentId"` // 部门id
	UserId       uint64   `map:"UserId"`       // 当前用户id

	// 生成token过长 暂时屏蔽不需要的
	//ID           uint64   `map:"ID"`           // id
	//Username     string   `map:"Username"`     // 用户登录名
	//NickName     string   `map:"NickName"`     // 昵称
}

// 获取token
func GetJwtToken(secretKey string, iat, seconds uint64, jwtData JWTData) (string, error) {
	fmt.Println("======== getJwtToken jwtData", jwtData)
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[CtxKeyJwtData] = jwtData
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetJwtData(ctx context.Context) JWTData {
	ctxKeyJwtData := ctx.Value(CtxKeyJwtData)
	var jwtData JWTData
	if ctxKeyJwtDataClaim, ok := ctxKeyJwtData.(map[string]interface{}); ok {
		jwtData = mapToJWTData(ctxKeyJwtDataClaim)
	}
	return jwtData
}

func mapToJWTData(data map[string]interface{}) JWTData {
	jwtData := JWTData{}
	v := reflect.ValueOf(&jwtData).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		mapTag := field.Tag.Get("map")

		if mapTag != "" {
			if value, ok := data[mapTag]; ok {
				switch v.Field(i).Kind() {
				case reflect.Int:
					if num, ok := value.(json.Number); ok {
						intValue, err := strconv.Atoi(num.String())
						if err != nil {
							fmt.Println("转换失败:", err)
						} else {
							v.Field(i).SetInt(int64(intValue))
						}
					}
				case reflect.Int64:
					if num, ok := value.(json.Number); ok {
						intValue, err := strconv.ParseInt(num.String(), 10, 64)
						if err != nil {
							fmt.Println("转换失败:", err)
						} else {
							v.Field(i).SetInt(intValue)
						}
					}
				default:
					if reflect.ValueOf(value).Type().AssignableTo(v.Field(i).Type()) {
						v.Field(i).Set(reflect.ValueOf(value))
					}
				}
			}
		}
	}
	fmt.Println("mapToJWTData", jwtData)
	return jwtData
}

func GetJwtDataRoleId(ctx context.Context) uint64 {
	return GetJwtData(ctx).RoleId
}
func GetJwtDataUserId(ctx context.Context) uint64 {
	return GetJwtData(ctx).UserId
}

//func GetJwtDataUsername(ctx context.Context) string {
//	return GetJwtData(ctx).Username
//}
//
//func GetJwtDataNickName(ctx context.Context) string {
//	return GetJwtData(ctx).NickName
//}
//func GetJwtDataID(ctx context.Context) uint64 {
//	return GetJwtData(ctx).ID
//}
