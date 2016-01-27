package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

func genToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
