package game

import (
	"errors"
	"github.com/Tactique/golib/logger"
)

func calculateDamage(attack *attack, armor *armor) int {
	damage := attack.power / armor.strength
	logger.Debugf("Calculated damage to be %d", damage)
	return damage
}

func damageUnit(attacker *unit, index int, defender *unit) (bool, error) {
	if len(attacker.attacks) > index {
		alive, err := defender.receiveDamage(calculateDamage(attacker.attacks[index], defender.armor))
		logger.Info("Defender is %t still alive with %d health", alive, defender.health)
		return alive, err
	} else {
		logger.Warn("Player does not have %d attacks (has %d)", index, len(attacker.attacks))
		return true, errors.New("Player does not have that many attacks")
	}
}
