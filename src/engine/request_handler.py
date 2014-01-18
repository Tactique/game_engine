import Queue
import json

import EBQP

from . import world
from . import types
from . import consts


class GameRequestHandler:
    def __init__(self):
        self.world = None
        self.responses = {
            EBQP.new: self.respond_new,
        }

    def process(self, request):
        request_pieces = request.split(EBQP.packet_delimiter, 1)
        command = request_pieces[0]
        params = request_pieces[1].strip() if len(request_pieces) > 1 else ''
        try:
            json_args = json.loads(params)
        except Exception as e:
            return "process:failure:bad json"

        if command in self.responses:
            return self.responses[command](json_args)
        else:
            return "process:failure:unsupported command"

    def respond_new(self, args):
        if 'uid1' not in args or 'uid2' not in args:
            return 'new:failure:missing uid'

        uid1 = args['uid1']
        uid2 = args['uid2']
        self.world = world.World([uid1, uid2])
        self.world.add_unit(uid1, types.new_unit('Tank', consts.RED))

        self.responses = {
            EBQP.view: self.respond_view,
            EBQP.move: self.respond_move,
        }
        return 'new:success'

    def respond_view(self, args):
        return 'view:success:%s' % self.world.to_json()

    #TODO
    def respond_move(self, args):
        return 'move:failure:unimplemented'
