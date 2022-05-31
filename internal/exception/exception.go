package exception

import (
	"encoding/json"
	"log"
)

type Exception struct {
	ErrorCode int    `json:"errorCode"`
	Msg       string `json:"msg"`
}

func NewApiException(errorCode int, msg string) string {
	e := Exception{ErrorCode: errorCode, Msg: msg}
	data, err := json.Marshal(e)
	if err != nil {
		log.Panic(err)
	}
	return string(data)

}
