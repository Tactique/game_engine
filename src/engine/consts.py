from lib import base, contract

MAX_HEALTH = 10


class Nation(base.BaseClass):
    def __init__(self, nation_id):
        self.nation_id = nation_id

    @contract.self_accepts(bool)
    @contract.returns(int)
    def serialize(self, public):
        return self.nation_id


class Team(base.BaseClass):
    def __init__(self, nation):
        self.team_id = nation.nation_id

    @contract.self_accepts(bool)
    @contract.returns(int)
    def serialize(self, public):
        return self.team_id

RED = Nation(0)
BLUE = Nation(1)
GREEN = Nation(2)
YELLOW = Nation(3)

NATIONS = [RED, BLUE, GREEN, YELLOW]
