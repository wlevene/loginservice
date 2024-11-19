package util

// func GetUserInfoFormAuth0(jwt string) map[string]interface{} {

// 	url := "https://insouslide.us.auth0.com/userinfo"
// 	client := resty.New()
// 	resp, err := client.R().
// 		SetHeader("Authorization", jwt).
// 		Get(url)

// 	if err != nil {
// 		logx.Error(err)
// 		return nil
// 	}

// 	values, err := Json2map(string(resp.Body()))

// 	if err != nil {
// 		logx.Error(err)
// 		return nil
// 	}

// 	return values
// }

// func GetUserFormAuth0(jwt string) *dao_user.User {

// 	user := GetUserInfoFormAuth0(jwt)
// 	if user == nil {
// 		return nil
// 	}

// 	result := &dao_user.User{}

// 	if v, ok := user["sub"].(string); ok {
// 		result.UserID = v
// 	}

// 	if v, ok := user["nickname"].(string); ok {
// 		result.Nick = v
// 	}

// 	if v, ok := user["email"].(string); ok {
// 		result.EMail = v
// 	}

// 	if v, ok := user["picture"].(string); ok {
// 		result.AvatarURL = v
// 	}

// 	if v, ok := user["picture"].(string); ok {
// 		result.AvatarURL = v
// 	}

// 	return result
// }
