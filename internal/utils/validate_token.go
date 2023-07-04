package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/golang-module/carbon/v2"
)

func Md5(rawText string) string {
	hasher := md5.New()
	hasher.Write([]byte(rawText))
	return hex.EncodeToString(hasher.Sum(nil))
}
func GenTenantToken(tenantId, sk string) string {
	now := carbon.Now().Format("Y-m-d")
	hashStr := fmt.Sprintf("%s::%s::%s", tenantId, sk, now)
	return Md5(hashStr)
}
