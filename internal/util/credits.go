package util

const (
	OneCreditsEqLetter = 150
)

func CreditsForString(content string) int32 {

	if content == "" {
		return 0
	}

	return CreditsForStringLen(int32(len(content))) // 64
}

func CreditsForStringLen(str_len int32) int32 {

	if str_len == 0 {
		return 0
	}

	return int32(str_len / OneCreditsEqLetter) // 64
}
