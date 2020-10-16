package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// DigestHeader структура заголовка digest auth
type digestHeader struct {
	realm     string
	domain    string
	nonce     string
	opaque    string
	stale     string
	algorithm string
}

// GetDigestString возвращает строку ответа неавторизированному пользователю
func GetDigestString(realm string, ip string) (string, string, error) {
	dh := new(digestHeader)

	ipInts := make([]int16, 0)
	for _, numStr := range strings.Split(ip, ".") {
		num, err := strconv.ParseInt(numStr, 10, 16)
		if err != nil {
			return "", "", err
		}
		ipInts = append(ipInts, int16(num))
	}

	dh.realm = realm
	dh.domain = ""
	dh.nonce = generateNonce(ipInts)
	dh.opaque = dh.nonce
	dh.stale = "false"

	return dh.nonce, fmt.Sprintf("Digest %s", dh.digestChallenges()), nil
}

func (dh *digestHeader) digestChallenges() string {
	return fmt.Sprintf("realm=%s, domain='%s', nonce=%s, opaque=%s, stale=%s", dh.realm, dh.domain, dh.nonce, dh.opaque, dh.stale)
}

func GetDigestHeaders(response string) (username, realm, nonce, uri, responseVal string) {
	if response != "" {
		resp := strings.Replace(response, "Digest ", "", 1)
		respStrs := strings.Split(resp, ", ")

		respKeysVals := make(map[string]string, 0)
		for _, str := range respStrs {
			keyVal := strings.Split(str, "=")
			respKeysVals[keyVal[0]] = strings.Trim(keyVal[1], "'\"")
		}

		username = respKeysVals["username"]
		realm = respKeysVals["realm"]
		nonce = respKeysVals["nonce"]
		uri = respKeysVals["uri"]
		responseVal = respKeysVals["response"]
	}

	return
}

func generateNonce(ip []int16) string {
	ipBytes := "" // ip в 16-ом виде
	for _, num := range ip {
		ipBytes += fmt.Sprintf("%x", num)
	}
	timeBytes := fmt.Sprintf("%x", time.Now().Unix()) // время в 16-ом виде
	keyBytes := fmt.Sprintf("%x", "token")            // ключ в 16-ом виде

	return ipBytes + timeBytes + keyBytes
}
