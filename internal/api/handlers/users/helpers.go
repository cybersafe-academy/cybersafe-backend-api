package users

var validMBTITypes = map[string]bool{
	"INTJ": true,
	"INTP": true,
	"ENTJ": true,
	"ENTP": true,
	"INFJ": true,
	"INFP": true,
	"ENFJ": true,
	"ENFP": true,
	"ISTJ": true,
	"ISFJ": true,
	"ESTJ": true,
	"ESFJ": true,
	"ISTP": true,
	"ISFP": true,
	"ESTP": true,
	"ESFP": true,
}

func isValidMBTIType(mbtiType string) bool {
	_, exists := validMBTITypes[mbtiType]

	return exists
}
