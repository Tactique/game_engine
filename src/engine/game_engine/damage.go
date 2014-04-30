package game_engine

import (
    "errors"
)

func calculateDamage(attack *attack, armor *armor) int {
    return attack.power / armor.strength
}

func damageUnit(attacker *unit, index int, defender *unit) (bool, error) {
    if len(attacker.attacks) > index {
        return defender.receiveDamage(calculateDamage(attacker.attacks[index], defender.armor))
    } else {
        return true, errors.New("Player does not have that many attacks")
    }
}
