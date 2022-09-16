package strutil

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func Utf82Gbk(text string) []byte {
	r := bytes.NewReader([]byte(text))
	decoder := transform.NewReader(r, simplifiedchinese.GBK.NewEncoder()) //GB18030
	content, _ := ioutil.ReadAll(decoder)
	return content
}

func Gbk2Utf8(b []byte) []byte {
	tfr := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(tfr)
	if e != nil {
		return nil
	}
	return d
}
