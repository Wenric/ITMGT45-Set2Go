package main

import (
	"strings"
)

func shiftLetter(letter string, shift int) string {
	if letter == " " {
		return " "
	}
	
	base := int('A')
	letterCode := int(letter[0]) - base
	newCode := (letterCode + shift) % 26
	return string(rune(base + newCode))
}

func caesarCipher(message string, shift int) string {
	result := ""
	for _, ch := range message {
		letter := string(ch)
		result += shiftLetter(letter, shift)
	}
	return result
}

func shiftByLetter(letter string, letterShift string) string {
	if letter == " " {
		return " "
	}
	
	shift := int(letterShift[0]) - int('A')
	return shiftLetter(letter, shift)
}

func vigenereCipher(message string, key string) string {
	result := ""
	keyIndex := 0
	keyLen := len(key)

	for _, ch := range message {
		letter := string(ch)
		if letter == " " {
			result += " "
		} else {
			shift := int(key[keyIndex%keyLen]) - int('A')
			result += shiftLetter(letter, shift)
			keyIndex++
		}
	}
	return result
}

func scytaleCipher(message string, shift int) string {
	msgLen := len(message)
	
	if msgLen%shift != 0 {
		padLen := shift - (msgLen % shift)
		message += strings.Repeat("_", padLen)
		msgLen += padLen
	}

	cols := msgLen / shift
	encoded := make([]byte, msgLen)

	for i := 0; i < msgLen; i++ {
		fromIdx := (i / shift) + cols*(i%shift)
		encoded[i] = message[fromIdx]
	}

	return string(encoded)
}

func scytaleDecipher(message string, shift int) string {
	msgLen := len(message)
	cols := msgLen / shift
	decoded := make([]byte, msgLen)

	for i := 0; i < msgLen; i++ {
		toIdx := (i / cols) + shift*(i%cols)
		decoded[toIdx] = message[i]
	}

	return string(decoded)
}

func main() {
	
}
