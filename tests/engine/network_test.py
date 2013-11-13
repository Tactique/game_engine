import unittest
import socket
import time
from threading import Thread

from engine import network
from engine import world


class NetworkTest(unittest.TestCase):
    def setUp(self):
        self.world = world.World([1, 2])

    def testSocket(self):
        self.host = '127.0.0.1'
        self.port = 8080

        def send_message(message):
            sock = socket.socket()
            sock.connect((self.host, self.port))
            sock.send(message)
            resp = sock.recv(1024)
            sock.close()
            return resp

        thread = Thread(
            target=network.listen,
            args=(self.host, self.port, self.world))
        thread.daemon = True
        thread.start()

        time.sleep(.1)

        resp = send_message('view')
        resp = send_message('exit')
        self.assertEqual(resp, 'exit')

    def testExit(self):
        resp = network.respond(self.world, 'exit')
        self.assertEqual(resp, 'exit')

    def testView(self):
        resp = network.respond(self.world, 'view')
        expected_resp_piece = '[%s0]' % ('0, ' * 9)
        expected_resp = '[%s%s]' % (
            (expected_resp_piece + ', ') * 9, expected_resp_piece)
        self.assertEqual(resp, expected_resp)

    def testMove(self):
        resp = network.respond(self.world, 'move')
        self.assertEqual(resp, 'unimplemented')

        resp = network.respond(self.world, 'move asdf')
        self.assertEqual(resp, 'unimplemented')

if __name__ == '__main__':
    unittest.main()
