from lib import base, contract

MAX_HEALTH = 10


class Team(base.BaseClass):
    def __init__(self, team_id):
        self.team_id = team_id

    @contract.self_accepts(bool)
    @contract.returns(int)
    def serialize(self, public):
        return self.team_id

RED = Team(0)
BLUE = Team(1)
GREEN = Team(2)
YELLOW = Team(3)

TEAMS = [RED, BLUE, GREEN, YELLOW]
