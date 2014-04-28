package requests

type NewCommandRequest struct {
    Uids []int `json:"uids"`
    Debug int `json:"debug"`
}

type ViewCommandRequest struct {
}

type MoveCommandRequest struct {
    Move []LocationStruct `json:"move"`
}

type EndTurnCommandRequest struct {
    PlayerId int `json:"playerId"`
}

type ExitCommandRequest struct {
    Reason string `json:"reason"`
}

type WorldStruct struct {
    Terrain [][]int `json:"terrain"`
    Units []*UnitStruct `json:"units"`
    Players []*PlayerStruct `json:"players"`
    TurnOwner int `json:"turnOwner"`
}

type PlayerStruct struct {
    PlayerId int `json:"playerId"`
    Nation int `json:"nation"`
    Team int `json:"team"`
}

type UnitStruct struct {
    Name string `json:"name"`
    Health int `json:"health"`
    Nation int `json:"nation"`
    Movement *MovementStruct `json:"movement"`
    Position *LocationStruct `json:"position"`
    Distance int `json:"distance"`
    CanMove bool `json:"canMove"`
}

type MovementStruct struct {
    Type string `json:"type"`
    Speeds map[string]float64 `json:"speeds"`
}

type LocationStruct struct {
    X int `json:"x"`
    Y int `json:"y"`
}
