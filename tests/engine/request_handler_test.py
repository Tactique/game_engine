import unittest

from engine import request_handler


class RequestHandler(unittest.TestCase):
    def setUp(self):
        self.handler = request_handler.GameRequestHandler()
        self.handler.process('new:{"uids": [0, 1]}')

    def test_request_new(self):
        self.assertEqual(
            request_handler.GameRequestHandler().respond_new({"uids": [0, 1]}),
            'new:success')

    #TODO support move
    def test_request_move(self):
        self.assertEqual(self.handler.respond_move({}), 'move:failure:unimplemented')

    def test_request_view(self):
        self.assertEqual(self.handler.respond_view({}).split(':')[:2], ['view', 'success'])

    def test_process_bad_jsons(self):
        self.assertEqual(self.handler.process('new:{"'), 'process:failure:bad json')
        self.assertEqual(
            request_handler.GameRequestHandler().process('new:{"uid": []}'),
            'new:failure:missing uid')

    def test_process_new_second_fails(self):
        self.assertEqual(
            self.handler.process('new:{"uids": [0, 1]}'),
            'process:failure:unsupported command')

    def test_send_debug(self):
        self.assertEqual(
            request_handler.GameRequestHandler().respond_new({"uids": [0, 1], 'debug': 0}),
            'new:success')

if __name__ == '__main__':
    unittest.main()
