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
    attack_types = AttackTypes()
    attack_list = file_loader.read_and_parse_json('attack_types')[0]
    for enumeration, attack_type in enumerate(attack_list):
        attack_types[str(attack_type)] = enumeration
    return attack_types


@contract.accepts(AttackTypes)
def load_attacks(attack_types):
    args = {}
    for attack_ in file_loader.read_and_parse_json('attacks'):
        name = str(attack_['name'])
        power = attack_['power']
        attack_type = attack_types[str(attack_['attack_type'])]
        args[name] = [name, power, attack_type]

    @contract.accepts(str)
    @contract.returns(Attack)
    def attack_getter(name):
        return Attack(*args[name])
    return attack_getter, args.keys()
