package mojoauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/dchest/uniuri"
	"strconv"
	"strings"
	"time"
)

const dayInSeconds = 86400

func CreateSecret() string {
	return uniuri.NewLen(93)
}

func CreateCredentials(id string, secret string, ttl_seconds ...int) (username string, password string) {
	ttl := dayInSeconds
	if len(ttl_seconds) > 0 {
		ttl = ttl_seconds[0]
	}

	total_ttl := nowTimeInt() + ttl

	if id == "" {
		username = fmt.Sprintf("%d", total_ttl)
	} else {
		username = fmt.Sprintf("%d:%s", total_ttl, id)
	}

	password = SignMessage(username, secret)
	return
}

func SignMessage(message string, secret string) (signed string) {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	signed = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}

func TestCredentials(username string, password string, secret string, ttl_timestamp ...int) (result bool, id string) {
	result = true
	id = ""
	max_ttl := nowTimeInt()
	split_parts := strings.Split(username, ":")
	if len(split_parts) > 1 {
		id = split_parts[1]
	}

	if len(ttl_timestamp) > 0 {
		max_ttl = ttl_timestamp[0]
	}

	expiry_timestamp, _ := strconv.Atoi(split_parts[0])
	if expiry_timestamp > max_ttl {
		result = false
	}
	if SignMessage(username, secret) != password {
		result = false
	}
	return
}

func nowTimeInt() int {
	return int(time.Now().UTC().Unix())
}
