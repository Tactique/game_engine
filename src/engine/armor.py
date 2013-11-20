from lib import contract, base

from . import attack, file_loader


class Armor(base.BaseMultiplier):
    pass


@contract.accepts(attack.AttackTypes)
def load_armors(attack_types):
    @contract.accepts(str)
    @contract.returns(Armor)
    def armor_getter(name):
        armor_ = args[name]
        resists = {}
        for attack_type, multiplier in armor_['resists'].items():
            resists[attack_types[str(attack_type)]] = multiplier
        return Armor(resists)

    args = file_loader.load_struct('armors')
    return armor_getter, args.keys()
