package jwt

import (
	"crypto/rsa"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type Jwt struct {
	// 签名算法 HS256, HS384, HS512
	// 默认 HS256.
	SigningAlgorithm string
	//私钥
	PriKeyFile string
	//公钥
	PubKeyFile string
	// Private key
	priKey *rsa.PrivateKey
	//普通密钥字符串
	Key []byte
	// Public key
	pubKey *rsa.PublicKey
}

func (j *Jwt) privateKey() error {
	keyData, err := os.ReadFile(j.PriKeyFile)
	if err != nil {
		return errors.New("私钥文件不存在")
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return errors.New("私钥文件不存在格式错误")
	}
	j.priKey = key
	return nil
}

func (j *Jwt) publicKey() error {
	keyData, err := os.ReadFile(j.PubKeyFile)
	if err != nil {
		return errors.New("公钥文件不存在")

	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return errors.New("公钥文件不存在格式错误")
	}
	j.pubKey = key
	return nil
}

// usingPublicKeyAlgo 是否使用公钥私钥算法签名
func (j *Jwt) usingPublicKeyAlgo() bool {
	switch j.SigningAlgorithm {
	case "RS256", "RS512", "RS384":
		return true
	}
	return false
}

func (j *Jwt) signedString(token *jwt.Token) (string, error) {
	var tokenString string
	var err error
	if j.usingPublicKeyAlgo() {
		tokenString, err = token.SignedString(j.priKey)
	} else {
		tokenString, err = token.SignedString(j.Key)
	}
	return tokenString, err
}
