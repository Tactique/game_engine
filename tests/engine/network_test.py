import unittest
import socket
import time
import json
from threading import Thread

from engine import network


class NetworkTest(unittest.TestCase):
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
            args=(self.host, self.port))
        thread.daemon = True
        thread.start()

        time.sleep(1)

        resp = send_message('new:{"uids": [0, 1]}')

if __name__ == '__main__':
    unittest.main()
