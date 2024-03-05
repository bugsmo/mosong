package web

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestEmailPattern 用来验证我们的邮箱正则表达式对不对
func TestEmailPattern(t *testing.T) {
	testCase := []struct {
		name  string
		email string
		match bool
	}{
		{
			name:  "不带@",
			email: "123456",
			match: false,
		},
		{
			name:  "带@, 但是没后缀",
			email: "123456@",
			match: false,
		},
		{
			name:  "合法邮箱",
			email: "123456@qq.com",
			match: true,
		},
	}

	h := NewUserHandler()

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			match, err := h.emailRegexExp.MatchString(tc.email)
			require.NoError(t, err)
			assert.Equal(t, tc.match, match)
		})
	}
}

// TestPasswordPattern
func TestPasswordPattern(t *testing.T) {
	testCase := []struct {
		name     string
		password string
		match    bool
	}{
		{
			name:     "没有数字",
			password: "Hello#world",
			match:    false,
		},
		{
			name:     "没有特殊字符",
			password: "Hello123world",
			match:    false,
		},
		{
			name:     "没有字母",
			password: "1234#5678",
			match:    false,
		},
		{
			name:     "长度不足",
			password: "He#123",
			match:    false,
		},
		{
			name:     "合法密码",
			password: "hello123#456world",
			match:    true,
		},
	}

	h := NewUserHandler()

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			match, err := h.passwordRegexExp.MatchString(tc.password)
			require.NoError(t, err)
			assert.Equal(t, match, tc.match)
		})
	}
}
