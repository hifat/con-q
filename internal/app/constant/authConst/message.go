package authConst

type msgType struct {
	INVALID_CREDENTIALS     string
	UNAUTHORIZED            string
	NO_AUTHORIZATION_HEADER string
	NOT_FOUND_BEARER        string
	BROKEN_TOKEN            string
	MAX_DEVICES_LOGIN       string
	INVALID_TOKEN           string
	TOKEN_EXPIRED           string
}

var Msg = msgType{
	INVALID_CREDENTIALS:     "invalid credentials",
	UNAUTHORIZED:            "unauthorized",
	NO_AUTHORIZATION_HEADER: "no authorization header provided",
	NOT_FOUND_BEARER:        "could not find bearer token in authorization header",
	BROKEN_TOKEN:            "the token is broken",
	MAX_DEVICES_LOGIN:       "maximum devices limit reached. please log out from one of your other devices before attempting to log in again",
	INVALID_TOKEN:           "invalid token",
	TOKEN_EXPIRED:           "token expired",
}
