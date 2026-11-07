package service

import (
	"fmt"
	"math/rand"
	"regexp"
	"testing"
)

func Test1(t *testing.T) {
	t.Skip("integration test: configure SMTP in config.yaml before enabling")
	SendRegisterEamil("5432", "test@example.com")
}

func Test2(t *testing.T) {
	s := rand.Intn(10000)
	fmt.Println(s)
}

func Test3(t *testing.T) {
	regex := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	fmt.Println(regex.MatchString("test@example.com"))
}
