package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func HashStr(doc string) string {
	var v interface{}
	if errUnmarshal := json.Unmarshal([]byte(doc), &v); errUnmarshal != nil {
		return ""
	}
	cdoc, _ := json.Marshal(v)
	sum := sha256.Sum256(cdoc)
	return hex.EncodeToString(sum[0:])
}
func HashBytes(b []byte) string {
	var v interface{}
	if errUnmarshal := json.Unmarshal(b, &v); errUnmarshal != nil {
		return ""
	}
	cdoc, _ := json.Marshal(v)
	sum := sha256.Sum256(cdoc)
	return hex.EncodeToString(sum[0:])
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func GetHmacSha256(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
func GetSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
