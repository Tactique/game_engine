from lib import contract, base

from . import file_loader


class AttackTypes(base.BaseEnum):
    pass


class Attack(base.BaseClass):
    @contract.self_accepts(str, int, int)
    def __init__(self, name, power, attackType):
        self.name = name
        self.power = power
        self.attackType = attackType


@contract.returns(AttackTypes)
def load_attack_types():
    return AttackTypes(file_loader.load_enum('attack_types'))


@contract.accepts(AttackTypes)
def load_attacks(attack_types):
    @contract.accepts(str)
    @contract.returns(Attack)
    def attack_getter(name):
        attack_ = args[name]
        power = attack_['power']
        attack_type = attack_types[str(attack_['attack_type'])]
        return Attack(name, power, attack_type)

    args = file_loader.load_struct('attacks')
    return attack_getter, args.keys()
