from lib import contract, base

from . import attack, file_loader


class Armor(base.BaseDictionary):
    pass


@contract.accepts(attack.AttackTypes)
def load_armors(attack_types):
    args = {}
    for armor_ in file_loader.read_and_parse_json('armors'):
        name = str(armor_['name'])
        resists = {}
        for attack_type, multiplier in armor_['resists'].items():
            resists[attack_types[str(attack_type)]] = multiplier
        args[name] = resists

    @contract.accepts(str)
    def armor_getter(name):
        return Armor(dict(args[name]))
    return armor_getter, args.keys()
