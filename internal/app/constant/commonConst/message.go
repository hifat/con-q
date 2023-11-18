package commonConst

type msgType struct {
	DUPLICATE_RECORD      string
	RECORD_NOTFOUND       string
	INTERNAL_SERVER_ERROR string
	TOO_MANY_REQUEST      string
}

var Msg = msgType{
	RECORD_NOTFOUND:       "record not found",
	DUPLICATE_RECORD:      "duplicate record",
	INTERNAL_SERVER_ERROR: "internal server error",
	TOO_MANY_REQUEST:      "too many request",
}
