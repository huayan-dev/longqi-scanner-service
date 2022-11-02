package tools

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	InStorageOrderSnPrefix  = "IN"
	OutStorageOrderSnPrefix = "OUT"
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func HexBytesToInt(b []byte) (n int, err error) {
	s := ""
	for _, b := range b {
		s += fmt.Sprintf("%v", b)
	}
	n, err = strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func HexBytesToFloat(b []byte) (val float64, err error) {
	ss := hex.EncodeToString(b)
	newB, err := hex.DecodeString(ss)
	buf := bytes.NewReader(newB)
	var f float32
	err = binary.Read(buf, binary.BigEndian, &f)
	if err != nil {
		return 0, err
	}
	val, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)
	return val, nil
}

func HexBytesToString(b []byte) (str string, err error) {
	ss := hex.EncodeToString(b)
	newB, err := hex.DecodeString(ss)
	return string(newB), err
}

func GeneratePassword(val string) (pwd string) {
	hash := md5.Sum([]byte(val))
	hash2 := md5.Sum([]byte(hex.EncodeToString(hash[:])))
	return hex.EncodeToString(hash2[:])
}

func PwdCorrect(originalStr, encodingPwd string) bool {
	return encodingPwd == GeneratePassword(originalStr)
}

func GenerateAccessToken() (token string) {
	str := time.Now().Format("2006-01-02 15:04:05")
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func SplitTagsFromBytes(tagBytes []byte) ([]string, error) {
	byteCount := 12
	tagArr := []string{}
	for i := 0; i < len(tagBytes); i++ {
		if i > 0 && (i+1)%byteCount == 0 {
			tag := fmt.Sprintf("%x", tagBytes[(i+1)-byteCount:i+1])
			if tag == "000000000000000000000000" {
				continue
			}
			tagArr = append(tagArr, tag)
		}
	}
	return tagArr, nil
}

// GetLngLatFromStr 获取纬度经度 lng纬度， lat经度
func GetLngLatFromStr(gpsStr string) (lng, lat string, err error) {
	if len(gpsStr) < 8 {
		err = errors.New("gps数据不完整 " + gpsStr)
		return
	}
	substr := gpsStr[5:]
	arr := strings.Split(substr, ",")
	if len(arr) < 2 {
		err = errors.New("gps数据不完整 " + gpsStr)
		return
	}
	lng = arr[0][7:]
	lat = arr[1][7:]
	return
}

func GenerateInStorageOrderSn(createdUserId uint) (sn string) {
	min := 0
	max := 9999
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min+1) + min
	//前缀+年月日时+创建者id+随机数
	sn = fmt.Sprintf("%s%s%03d%04d", InStorageOrderSnPrefix, time.Now().Format("2006010215"), createdUserId, randNum)
	return sn
}

func GenerateOutStorageOrderSn(createdUserId uint) (sn string) {
	min := 0
	max := 9999
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min+1) + min
	//前缀+年月日时+创建者id+随机数
	sn = fmt.Sprintf("%s%s%03d%04d", OutStorageOrderSnPrefix, time.Now().Format("2006010215"), createdUserId, randNum)
	return sn
}

func GenerateOutStoragePictureName(orderId uint) (pictureName string) {
	min := 0
	max := 9999
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min+1) + min
	//前缀+年月日时+出库单id+随机数
	pictureName = fmt.Sprintf("%s%s%03d%04d%s", OutStorageOrderSnPrefix, time.Now().Format("2006010215"), orderId, randNum, ".jpg")
	return pictureName
}

func GetFileExt(filename string) (t string, err error) {
	slc := strings.Split(filename, ".")
	if len(slc) <= 1 {
		return t, errors.New("文件类型有误")
	}
	return slc[len(slc)-1], nil
}
