package api

const (
	STATUS_OK int = iota
	STATUS_FAILURE
	STATUS_MALFORMED_RESPONSE
	STATUS_BAD_REQUEST
	STATUS_UNKNOWN_REQUEST
	STATUS_UNSERIALIZEABLE_RESPONSE
)

const (
	COMMAND_NEW          string = "new"
	COMMAND_EXIT         string = "exit"
	COMMAND_MOVE         string = "move"
	COMMAND_TURN         string = "turn"
	COMMAND_ATTACK       string = "attack"
	COMMAND_VIEW_WORLD   string = "viewWorld"
	COMMAND_VIEW_TERRAIN string = "viewTerrain"
	COMMAND_VIEW_UNITS   string = "viewUnits"
	COMMAND_VIEW_PLAYERS string = "viewPlayers"
)
