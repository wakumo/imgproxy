package main

import (
	"math"
	"golang.org/x/crypto/salsa20"
	"os"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minNonZeroInt(a, b int) int {
	switch {
	case a == 0:
		return b
	case b == 0:
		return a
	}

	return minInt(a, b)
}

func roundToInt(a float64) int {
	return int(math.Round(a))
}

func scaleInt(a int, scale float64) int {
	if a == 0 {
		return 0
	}

	return roundToInt(float64(a) * scale)
}

func salsa20_decode(decoded_in []byte) string {
	var nonce = []byte(os.Getenv("IMGPROXY_SALT")[:8])
  var arr [32]byte
  var key_str = os.Getenv("IMGPROXY_KEY")[:32]
  copy(arr[:], key_str)
	// var in = []byte(decoded_in)
	var out = make([]byte,len(decoded_in))
	salsa20.XORKeyStream(out, decoded_in, nonce, &arr)
	return string(out)
}
