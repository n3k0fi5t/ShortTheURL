package base62

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint(len(alphabet))
)

func Encode(number uint, nLength int) (string, error) {
	var builder strings.Builder
	maxEncodableNumber := uint(math.Pow(float64(length), float64(nLength)) - 1)

	if number > maxEncodableNumber {
		return "", errors.New("cannot encode number exceed: " + strconv.FormatUint(uint64(maxEncodableNumber), 10))
	}

	builder.Grow(nLength)

	for l := 0; number > 0; number, l = number/length, l+1 {
		builder.WriteByte(alphabet[number%length])
	}

	for builder.Len() < builder.Cap() {
		builder.WriteByte(alphabet[0])
	}

	return builder.String(), nil
}

func Decode(encoded string) (uint, error) {
	var number uint

	for i, symbol := range encoded {
		pos := strings.IndexRune(alphabet, symbol)

		if pos == -1 {
			return uint(pos), errors.New("invalid charactor: " + string(symbol))
		}
		number += uint(pos) * uint(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}
