import socket
import string
import threading

from lib import contract

from . import respond_handler

class Connection(threading.Thread):
    def __init__(self, conn, addr, request_handler):
        threading.Thread.__init__(self)
        print 'received connection'
        self.conn = conn
        self.addr = addr
        self.request_handler = request_handler

    def read_loop(self):
        while True:
            #TODO handle more than 1024
            request = self.conn.recv(1024)
            if not request:
                print 'client closed socket'
                break
            print 'received request %s' % (request,)
            response = self.request_handler.process(request)
            print 'responding :', response
            self.conn.send(response)

    def run(self):
        self.read_loop()


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
        conn = Connection(conn, addr, respond_handler.GameRequestHandler())
        conn.start()

    sock.close()
