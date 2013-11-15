from . import attack

attack_types = attack.load_attack_types()

new_attack, attacks = attack.load_attacks(attack_types)
