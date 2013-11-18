from lib import contract

from . import consts, unit


@contract.accepts(unit.Unit, unit.Unit)
@contract.returns(None)
def DoDamage(attacker, receiver):
    weapon = 0
    attack = attacker.attacks[weapon]
    receiver.health -= int(
        receiver.armor[attack.attackType]
        * attack.power
        * (attacker.health / float(consts.MAX_HEALTH)))
