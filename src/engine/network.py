import socket
import string

from lib import contract

from . import world


@contract.accepts(str, int, world.World)
@contract.returns(None)
def listen(host, port, world_):
    sock = socket.socket()
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    sock.bind((host, port))
    sock.listen(1)

    running = True
    while running:
        conn, addr = sock.accept()
        request = conn.recv(1024)
        resp = respond(world_, request)
        conn.send(resp)
        conn.close()
        if 'exit' in resp:
            running = False

    sock.close()


@contract.accepts(world.World, str)
@contract.returns(str)
def respond(world_, request):
    request_pieces = string.split(request, ' ', maxsplit=1)
    command = request_pieces[0]
    if len(request_pieces) > 1:
        params = request_pieces[1]
    else:
        params = ''
    return responses[command](world_, params)


@contract.accepts(world.World, str)
@contract.returns(str)
def respond_view_world(world_, args):
    return world_.to_json()


#TODO
@contract.accepts(world.World, str)
@contract.returns(str)
def respond_move(world_, args):
    return 'unimplemented'


@contract.accepts(world.World, str)
@contract.returns(str)
def respond_exit(world_, args):
    return 'exit'

responses = {
    'view': respond_view_world,
    'exit': respond_exit,
    'move': respond_move,
}
