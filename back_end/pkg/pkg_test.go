package pkg

import (
	"fmt"
	"shoollove/config"
	jwtutil "shoollove/pkg/jwt_util"
	"shoollove/pkg/tool"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Test1(t *testing.T) {
	message := make(jwt.MapClaims)
	message["phone"] = "test@test.com"
	token, _ := jwtutil.GenerateJwt([]byte(config.Config.Jwt.Key), jwt.SigningMethodHS256, message)
	fmt.Println(token)
}

func Test2(t *testing.T) {
	parseJwt, err := jwtutil.ParseJwt([]byte(config.Config.Jwt.Key), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20ifQ.Jm4DvW4dr16TaaJx29IcvaLJxoOCCQ2RHdSwI95LqRE")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parseJwt)
}

func Test3(t *testing.T) {
	claims, err := jwtutil.ParseJwtWithClaims([]byte(config.Config.Jwt.Key), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20ifQ.Jm4DvW4dr16TaaJx29IcvaLJxoOCCQ2RHdSwI95LqRE")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(claims)
}

func Test4(t *testing.T) {
	str := "[1,2,3,4,5,6,7,8]"
	//slice, err := tool.StringToIntSlice(str)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(slice)
	deduplicate, err := tool.AppendAndDeduplicate(&str, 1, 2, 3, 9)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(deduplicate)
}

func Test5(t *testing.T) {
	t1 := time.Now()
	t2 := time.Now().Add(61 * time.Minute)
	str := tool.FormatTimeInterval(t1, t2)
	fmt.Println(str)
}

func Test6(t *testing.T) {
	str := strings.ToUpper(tool.Uid())
	fmt.Println(str)
	fmt.Println(len(str))
}
