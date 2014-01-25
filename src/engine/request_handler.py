import Queue
import json

import EBQP

from . import world
from . import types
from . import consts
from . import loc


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
        if 'uids' not in args:
            return 'new:failure:missing uid'
        uids = args['uids']
        self.world = world.World(uids)
        if 'debug' in args:
            unit_ = types.new_unit('Tank', consts.RED)
            loc_ = loc.Loc(3, 3)
            self.world.add_unit(unit_, loc_)

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
