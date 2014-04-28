package template_generator

import (
    "fmt"
    "encoding/json"
    "requests"
)

func GenerateAllTemplates() {
    ret_game := &requests.WorldStruct{
        Terrain: [][]int{
            []int{0, 0, 0},
            []int{0, 0, 0},
            []int{0, 0, 0}},
        Units: []*requests.UnitStruct{
            &requests.UnitStruct{
                Name: "Warrior", Health: 10, Nation: 0,
                Movement: &requests.MovementStruct{
                    Type: "Legs", Speeds: map[string]float64{
                        "0": 1}},
                Position: &requests.LocationStruct{X: 0, Y: 0},
                Distance: 10,
                CanAttack: true,
                CanMove: true,
                Attacks: []*requests.AttackStruct{
                    &requests.AttackStruct{
                        0, 5}}}},
        Players: []*requests.PlayerStruct{
            &requests.PlayerStruct{
                PlayerId: 26, Nation: 0, Team: 0},
            &requests.PlayerStruct{
                PlayerId: 13, Nation: 1, Team: 1}},
        TurnOwner: 0}
    jsonString, err := json.MarshalIndent(ret_game, "", "    ")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(jsonString))
}
