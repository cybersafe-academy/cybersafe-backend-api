package users

type mbti string

const (
	INTJ mbti = "INTJ"
	INTP mbti = "INTP"
	ENTJ mbti = "ENTJ"
	ENTP mbti = "ENTP"
	INFJ mbti = "INFJ"
	INFP mbti = "INFP"
	ENFJ mbti = "ENFJ"
	ENFP mbti = "ENFP"
	ISTJ mbti = "ISTJ"
	ISFJ mbti = "ISFJ"
	ESTJ mbti = "ESTJ"
	ESFJ mbti = "ESFJ"
	ISTP mbti = "ISTP"
	ISFP mbti = "ISFP"
	ESTP mbti = "ESTP"
	ESFP mbti = "ESFP"
)

func isValidMBTIType(mbtiType string) bool {
	switch mbti(mbtiType) {
	case INTJ, INTP, ENTJ, ENTP, INFJ, INFP, ENFJ, ENFP, ISTJ, ISFJ, ESTJ, ESFJ, ISTP, ISFP, ESTP, ESFP:
		return true
	default:
		return false
	}
}
