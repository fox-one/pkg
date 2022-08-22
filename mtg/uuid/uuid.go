package uuid

import (
	"crypto/md5"
	"io"

	"github.com/google/uuid"
)

type UUID = uuid.UUID

func New() UUID {
	return uuid.New()
}

func NewString() string {
	return uuid.NewString()
}

func IsUUID(id string) bool {
	_, err := FromString(id)
	return err == nil
}

func Modify(base, modifier string) string {
	ns := uuid.MustParse(base)
	return uuid.NewSHA1(ns, []byte(modifier)).String()
}

func MD5(input string) string {
	h := md5.New()
	_, _ = io.WriteString(h, input)
	sum := h.Sum(nil)
	sum[6] = (sum[6] & 0x0f) | 0x30
	sum[8] = (sum[8] & 0x3f) | 0x80
	id, _ := uuid.ParseBytes(sum)
	return id.String()
}

func FromString(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

func IsNil(id string) bool {
	uid, err := FromString(id)
	return err != nil || uid == uuid.Nil
}
