import socket
import string
import threading

from lib import contract

from . import request_handler


def connect(conn, addr, request_handler):
    print 'received connection'
    while True:
        #TODO handle more than 1024
        request = conn.recv(1024)
        if not request:
            print 'client closed socket'
            break
        print 'received request \'%s\'' % request
        response = request_handler.process(request)
        print 'responding : \'%s\'' % response
        conn.send(response)


@contract.accepts(str, int)
@contract.returns(None)
def listen(host, port):
    sock = socket.socket()
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    sock.bind((host, port))
    sock.listen(1)

    while True:
        print 'listening'
        conn, addr = sock.accept()
        thread = threading.Thread(
            target=connect,
            args=(conn, addr, request_handler.GameRequestHandler()))
        thread.daemon = True
        thread.start()

    #sock.close()
