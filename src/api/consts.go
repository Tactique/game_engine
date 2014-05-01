package api

const (
	STATUS_OK int = iota
	STATUS_FAILURE
	STATUS_MALFORMED_RESPONSE
	STATUS_BAD_REQUEST
	STATUS_UNKNOWN_REQUEST
	STATUS_UNSERIALIZEABLE_RESPONSE
)