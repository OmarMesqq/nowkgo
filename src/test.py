import unittest
from unittest.mock import patch
from server import bindPort, HOST
import socket

class TestServer(unittest.TestCase):
    @patch('socket.socket.bind')
    def test_bindPort(self, mock_bind):
        mock_bind.side_effect = [OSError, None]
        server = socket.socket()
        port = 8000
        bindPort(server, port)
        mock_bind.assert_called_with((HOST, port + 1))

if __name__ == "__main__":
    unittest.main()

# Teste do strict no zip do punchlines