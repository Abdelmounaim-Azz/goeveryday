package main
import (
"bytes"
"crypto/aes"
"crypto/cipher"
"encoding/hex"
"fmt"
"io"
)
func Hash(keyValue, secret string) string {
key, _ := hex.DecodeString(keyValue)
bReader := bytes.NewReader([]byte(secret))
block, err := aes.NewCipher(key)
if err != nil {
panic(err)
}
var iv [aes.BlockSize]byte
stream := cipher.NewOFB(block, iv[:])
var out bytes.Buffer
writer := &cipher.StreamWriter{S: stream, W: &out}
if _, err := io.Copy(writer, bReader); err != nil {
panic(err)
}
return fmt.Sprintf("%x\n", out.Bytes())
}

func main(){
	key := "123456789088"
	res:=Hash("6368616e676520746869732070617373", key)
	fmt.Println(res)
}