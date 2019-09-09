package utils

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"net/http"
	"strings"
)

func HmacSHA1Base64Encrypt(encryptText, seed string) string {
	mac := hmac.New(sha1.New, []byte(seed))
	mac.Write([]byte(encryptText))
	bytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(bytes)
}

func StringJoin(strings ...string) string {
	var buffer bytes.Buffer
	for _, str := range strings {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func Get6RandomNumber() string {
	return strconv.FormatInt(rand.Int63n(899999)+100000, 10)
}

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

func GetFromMap(mp map[string]interface{}, key string) interface{} {
	if val, ok := mp[key]; ok {
		return val
	}
	return nil
}

func ExecutableDir() string {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Printf("util: find executableDir err: %v", err)
		return ""
	}
	return filepath.Dir(pathAbs)
}

func CheckFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func ReadBufio(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)
	buf := make([]byte, 1024)

	for {
		readNum, err := bufReader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == readNum {
			break
		}
	}
}

func ClientPublicIP(r *http.Request) string {
	var ips string
	ips = strings.TrimSpace(r.Header.Get("X-Forwarded-For"))
	if ips == "" || ips == "unknown" {
		ips = strings.TrimSpace(r.Header.Get("Proxy-Client-IP"))
	}
	if ips == "" || ips == "unknown" {
		ips = strings.TrimSpace(r.Header.Get("WL-Proxy-Client-IP"))
	}
	if ips == "" || ips == "unknown" {
		ips = r.RemoteAddr
	}
	for _, ip := range strings.Split(ips, ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && ip != "unknown" {
			return ip
		}
	}
	return ""
}
