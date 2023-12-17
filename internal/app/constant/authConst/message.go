package authConst

var Msg = authConstType{
	INVALID_CREDENTIALS:       "invalid credentials",
	UNAUTHORIZED:              "unauthorized",
	NOT_FOUND_BEARER:          "could not find bearer token in authorization header",
	BROKEN_TOKEN:              "the token is broken",
	MAX_DEVICES_LOGIN:         "maximum devices limit reached. please log out from one of your other devices before attempting to log in again",
	INVALID_TOKEN:             "invalid token",
	TOKEN_EXPIRED:             "token expired",
	NO_AUTHORIZATION_HEADER:   "no authorization header provided",
	NO_X_REFRESH_TOKEN_HEADER: "no X-Refresh-Token header provided",
	REVOKED_TOKEN:             "token has been revoked",
}
