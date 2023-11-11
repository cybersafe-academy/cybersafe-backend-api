package helpers

import "strings"

type MBTI string

const (
	INTJ MBTI = "INTJ"
	INTP MBTI = "INTP"
	ENTJ MBTI = "ENTJ"
	ENTP MBTI = "ENTP"
	INFJ MBTI = "INFJ"
	INFP MBTI = "INFP"
	ENFJ MBTI = "ENFJ"
	ENFP MBTI = "ENFP"
	ISTJ MBTI = "ISTJ"
	ISFJ MBTI = "ISFJ"
	ESTJ MBTI = "ESTJ"
	ESFJ MBTI = "ESFJ"
	ISTP MBTI = "ISTP"
	ISFP MBTI = "ISFP"
	ESTP MBTI = "ESTP"
	ESFP MBTI = "ESFP"
)

func IsValidMBTIType(mbtiType string) bool {
	switch MBTI(strings.ToUpper(mbtiType)) {
	case INTJ, INTP, ENTJ, ENTP, INFJ, INFP, ENFJ, ENFP, ISTJ, ISFJ, ESTJ, ESFJ, ISTP, ISFP, ESTP, ESFP:
		return true
	default:
		return false
	}
}
