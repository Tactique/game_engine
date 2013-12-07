import Queue

from . import EBQP
from . import world


class GameRequestHandler:
    def __init__(self):
        self.worlld = None
        self.responses = {
            EBQP.new : self.respond_new,
        }

    def process(self, request):
        request_pieces = request.split(EBQP.packet_delimiter, 1)
        command = request_pieces[0]
        params = request_pieces[1] if len(request_pieces) > 1 else ''

        print self.responses
        try:
            response = self.responses[command]
        except KeyError:
            return "500"

        return response(params)

    def respond_new(self, args):
        #TODO PARSE JSON HERE DUMMY
        uid1, uid2 = 1, 0
        self.world = world.World([uid1, uid2])
        self.responses.update({
            EBQP.view: self.respond_view,
            EBQP.exit: self.respond_exit,
            EBQP.move: self.respond_move,
        })
        return 'ok'

    def respond_view(self, args):
        return self.world.to_json()

    #TODO
    def respond_move(self, args):
        return 'unimplemented'

    def respond_exit(self, args):
        return 'exit'

