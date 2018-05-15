package info

import (
	"strconv"
	"encoding/base32"
	"fmt"
)

type UserNonce struct {
	UserInfo
}

const (
	UserNonce_Encoding_Head        = "[USER'S NONCE] "
	UserNonce_Encoding_Tail        = "[/USER'S NONCE]"
	UserNonce_Encoding_Head_Length = len(UserNonce_Encoding_Head)
	UserNonce_Encoding_Tail_Length = len(UserNonce_Encoding_Tail)
)

func (nonce *UserNonce) Encode(val int) string {
	raw := ([]byte)(UserNonce_Encoding_Head + strconv.Itoa(val) + UserNonce_Encoding_Tail)
	return base32.HexEncoding.EncodeToString(raw) //or other encoding forms
}

func (nonce *UserNonce) Decode(str string) (int, error) {
	raw, err := base32.HexEncoding.DecodeString(str)
	if err != nil {
		return 0, fmt.Errorf("Failed in decoding.")
	} else if string(raw[:UserNonce_Encoding_Head_Length]) != UserNonce_Encoding_Head {
		return 0, fmt.Errorf("Data format error: it's not a valid user nonce data")
	} else if string(raw[len(raw)-UserNonce_Encoding_Tail_Length:]) != UserNonce_Encoding_Tail {
		return 0, fmt.Errorf("Data format error: it's not a valid user nonce data")
	} else {
		val, err := strconv.Atoi(string(raw[UserNonce_Encoding_Head_Length : len(raw)-UserNonce_Encoding_Tail_Length]))
		if err != nil {
			return 0, fmt.Errorf("Data format error: it's not a valid user nonce data")
		} else {
			return val, nil
		}
	}
}

