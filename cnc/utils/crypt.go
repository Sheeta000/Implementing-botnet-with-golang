package utils

import "encoding/base64"

var s = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()+/")

func MessageCrypt(str string) string {
	x1 := s.EncodeToString([]byte(str))
	x2 := encryptCaesar(x1, 6)
	x3 := encryptROT13(x2)
	return x3
}
func MessageDeclassify(str string) string {
	x1 := decryptROT13(str)
	x2 := decryptCaesar(x1, 6)
	x3, _ := s.DecodeString(x2)
	return string(x3)
}

func encryptCaesar(plaintext string, shift int) string {
	ciphertext := ""
	for _, c := range plaintext {
		if c >= 'a' && c <= 'z' {
			c = 'a' + (c-'a'+rune(shift))%26
		} else if c >= 'A' && c <= 'Z' {
			c = 'A' + (c-'A'+rune(shift))%26
		}
		ciphertext += string(c)
	}
	return ciphertext
}

func decryptCaesar(ciphertext string, shift int) string {
	plaintext := ""
	for _, c := range ciphertext {
		if c >= 'a' && c <= 'z' {
			c = 'a' + (c-'a'-rune(shift)+26)%26
		} else if c >= 'A' && c <= 'Z' {
			c = 'A' + (c-'A'-rune(shift)+26)%26
		}
		plaintext += string(c)
	}
	return plaintext
}

func encryptROT13(plaintext string) string {
	ciphertext := ""
	for _, c := range plaintext {
		switch {
		case c >= 'a' && c <= 'm':
			ciphertext += string(c + 13)
		case c >= 'n' && c <= 'z':
			ciphertext += string(c - 13)
		case c >= 'A' && c <= 'M':
			ciphertext += string(c + 13)
		case c >= 'N' && c <= 'Z':
			ciphertext += string(c - 13)
		default:
			ciphertext += string(c)
		}
	}
	return ciphertext
}

func decryptROT13(ciphertext string) string {
	return encryptROT13(ciphertext)
}
