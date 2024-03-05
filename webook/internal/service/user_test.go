package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordencrypt(t *testing.T) {
	pwd := []byte("Hello#world123789")
	encrypted, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	println(len(encrypted))
	err := bcrypt.CompareHashAndPassword(encrypted, pwd)
	require.NoError(t, err)
}
