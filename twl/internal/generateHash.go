package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type AuthorizationData struct {
	FirstName string `json:"first_name,omitempty"`
	Hash      string `json:"hash,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	PhotoURL  string `json:"photo_url,omitempty"`
	Username  string `json:"username,omitempty"`
	AuthDate  int64  `json:"auth_date,omitempty"`
	ID        int64  `json:"id,omitempty"`
}

func New() *AuthorizationData {
	return &AuthorizationData{}
}

func (x *AuthorizationData) String() string {
	vs := make([]string, 0, 6)
	addNotEmpty := func(key, value string) {
		if value != "" {
			vs = append(vs, fmt.Sprintf("%s=%s", key, value))
		}
	}

	addNotEmpty("auth_date", strconv.FormatInt(x.AuthDate, 10))
	addNotEmpty("first_name", x.FirstName)
	addNotEmpty("id", strconv.FormatInt(x.ID, 10))
	addNotEmpty("last_name", x.LastName)
	addNotEmpty("photo_url", x.PhotoURL)
	addNotEmpty("username", x.Username)

	return strings.Join(vs, "\n")
}

func (x *AuthorizationData) Sum(token string) string {
	return hex.EncodeToString(HashHMAC([]byte(x.String()), HashSHA256([]byte(token)), sha256.New))
}
