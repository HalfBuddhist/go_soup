package ts_encoding

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"testing"
)

// 测试RSA生成与加密的整个流程。
func TestRsa(t *testing.T) {
	// 1. 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("生成密钥对失败:", err)
	}

	// 2. 获取公钥
	publicKey := &privateKey.PublicKey

	// 3. 要加密的消息
	message := []byte("你好，这是一个需要加密的消息！")

	// 4. 使用公钥加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		log.Fatal("加密失败:", err)
	}
	fmt.Printf("加密后的数据: %x\n", ciphertext)

	// 5. 使用私钥解密
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		log.Fatal("解密失败:", err)
	}
	fmt.Printf("解密后的数据: %s\n", string(plaintext))

	// 6. 将私钥转换为PEM格式（用于存储）
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	fmt.Printf("\n私钥PEM格式:\n%s\n", pem.EncodeToMemory(privateKeyPEM))

	// 7. 将公钥转换为PEM格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatal("序列化公钥失败:", err)
	}
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	fmt.Printf("\n公钥PEM格式:\n%s\n", pem.EncodeToMemory(publicKeyPEM))
}

// 测试RSA加密
func TestRsaEncrypt(t *testing.T) {
	// 1. PEM格式的公钥字符串
	publicKeyPEM := `-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAIvta56yduhTMVHSCkEHofIeBCzwtJVUHLQ7G8J0+vBLadxg6MWG0fAzMPPfzPzhz857OfVyBr0BClq6/cgkobcCAwEAAQ==
-----END PUBLIC KEY-----`

	// 2. 解码PEM格式的公钥
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		t.Fatal("failed to parse PEM block containing the public key")
	}

	// 3. 解析公钥
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		t.Fatal("failed to parse DER encoded public key:", err)
	}

	publicKey := pub.(*rsa.PublicKey)

	// 7. 将公钥转换为PEM格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatal("序列化公钥失败:", err)
	}
	publicKeyPEM2 := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	fmt.Printf("\n公钥PEM格式:\n%s\n", pem.EncodeToMemory(publicKeyPEM2))

	// 4. 要加密的消息
	message := []byte("Macro3_js")

	// 5. 使用公钥加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		t.Fatal("failed to encrypt:", err)
	}

	// 6. 输出加密后的数据
	fmt.Printf("Encrypted message: %x\n", ciphertext)

	// 7. 将加密后的数据转换为base64
	base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("Base64 Encrypted message: %s\n", base64Ciphertext)
}
