package util

import (
	"bytes"
	"encoding/base64"
	"github.com/aofei/cameron"
	"go.uber.org/zap"
	"image"
	"image/png"
)

const defaultAvatar = "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEU5gWfGfph2LxzOAAAAeklEQVR4nOzWMQoCQQwF0IAH8+oebCGCjjuG0Wpg07zfpMhrw0/s5555vEfN8dkBANACZm6Z+TjHjwAA0AJiPVcAANpALdZxvP+bFwAAAACAV8Ym59MLAEAfiHKucW6+HAAArWANAAB9oH67deRoXgAArgfbeQYAAP//KgOqkcdmwTgAAAAASUVORK5CYII="
const size = 256

func GenerateAvatar(username string) image.Image {
	return cameron.Identicon([]byte(username), size, size/9)
}

func ImageToBase64(img image.Image) string {
	var buf bytes.Buffer

	err := png.Encode(&buf, img)
	if err != nil {
		zap.L().Error("Failed to encode image to base64", zap.Error(err))
		return defaultAvatar
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
