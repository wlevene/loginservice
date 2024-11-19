package user

type UserProTyp int32

const (
	UserProTyp_Free = iota
	UserProTyp_Plus
	UserProTyp_Pro
	UserProTyp_SAP
)

func UserProTypToString(pro_typ UserProTyp) string {

	switch pro_typ {

	case UserProTyp_Free:
		return "free"

	case UserProTyp_Plus:
		return "plus"

	case UserProTyp_Pro:
		return "pro"

	case UserProTyp_SAP:
		return "sap"
	}

	return "free"

}

func UserProTypFromString(pro_typ string) UserProTyp {

	switch pro_typ {

	case "free":
		return UserProTyp_Free

	case "plus":
		return UserProTyp_Plus

	case "pro":
		return UserProTyp_Pro

	case "sap":
		return UserProTyp_Pro
	}

	return UserProTyp_Free

}
