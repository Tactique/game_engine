package api

type ResponseType struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}

type NewRequest struct {
	Uids    []int `json:"uids"`
	Debug   int   `json:"debug"`
	WorldId int   `json:"worldId"`
}

type NewResponse struct {
	Uids    []int `json:"uids"`
	Debug   int   `json:"debug"`
	WorldId int   `json:"worldId"`
}

type ViewWorldRequest struct {
}

type ViewWorldResponse struct {
	TerrainResponse ViewTerrainResponse
	UnitsResponse   ViewUnitsResponse
	PlayersResponse ViewPlayersResponse
}

type ViewTerrainRequest struct {
}

type ViewTerrainResponse struct {
	Terrain [][]int `json:"terrain"`
}

type ViewUnitsRequest struct {
}

type ViewUnitsResponse struct {
	Units []*UnitStruct `json:"units"`
}

type ViewPlayersRequest struct {
}

type ViewPlayersResponse struct {
	Me        *PlayerStruct   `json:"me"`
	TeamMates []*PlayerStruct `json:"teamMates"`
	Enemies   []*PlayerStruct `json:"enemies"`
	TurnOwner int             `json:"turnOwner"`
}

type MoveRequest struct {
	Move []LocationStruct `json:"move"`
}

type MoveResponse struct {
	Move []LocationStruct `json:"move"`
}

type AttackRequest struct {
	Attacker    LocationStruct `json:"attacker"`
	AttackIndex int            `json:"attackIndex"`
	Target      LocationStruct `json:"target"`
}

type AttackResponse struct {
	Attacker    LocationStruct `json:"attacker"`
	AttackIndex int            `json:"attackIndex"`
	Target      LocationStruct `json:"target"`
}

type EndTurnRequest struct {
}

type EndTurnResponse struct {
}

type ExitRequest struct {
	Reason string `json:"reason"`
}

type PlayerStruct struct {
	Nation int `json:"nation"`
	Team   int `json:"team"`
}

type UnitStruct struct {
	Name      string          `json:"name"`
	Health    int             `json:"health"`
	Nation    int             `json:"nation"`
	Movement  *MovementStruct `json:"movement"`
	Position  *LocationStruct `json:"position"`
	CanMove   bool            `json:"canMove"`
	Attacks   []*AttackStruct `json:"attacks"`
	CanAttack bool            `json:"canAttack"`
	Armor     *ArmorStruct    `json:"armor"`
}

type AttackStruct struct {
	Name       string `json:"name"`
	AttackType int    `json:"attackType"`
	Power      int    `json:"power"`
	MinRange   int    `json:"minRange"`
	MaxRange   int    `json:"maxRange"`
}

type ArmorStruct struct {
	Name      string `json:"name"`
	ArmorType int    `json:"armorType"`
	Strength  int    `json:"strength"`
}

type MovementStruct struct {
	Type     string             `json:"type"`
	Speeds   map[string]float64 `json:"speeds"`
	Distance int                `json:"distance"`
}

type LocationStruct struct {
	X int `json:"x"`
	Y int `json:"y"`
}
