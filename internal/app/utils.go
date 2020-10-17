package class

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"regexp"

	"github.com/ELQASASystem/app/internal/app/qq"
)

// chainToString 消息链转文本
func chainToString(chain []qq.Chain) (fullText string) {

	for _, element := range chain {
		fullText += element.Text
	}
	return

}

// isAnswer 是否为合法答案 [选择题]
func isAnswer(answer string) (ok bool) {
	ok, _ = regexp.MatchString("^[A-Za-z]$", answer)
	return
}

// HashSHA1 将答题数据散列
func HashSHA1(data interface{}) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%v", data)))
	return hex.EncodeToString(h.Sum(nil))
}
