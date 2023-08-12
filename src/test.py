import unittest
from unittest.mock import patch
from server import bindPort, HOST
import socket
from jokes import setups, punchlines

class TestServer(unittest.TestCase):
    """
    Testa a resiliência do servidor caso o SO não deixe bindar a porta de operação
    """
    @patch('socket.socket.bind')
    def test_bindPort(self, mock_bind):
        mock_bind.side_effect = [OSError, None]
        with socket.socket() as server:
            port = 9001
            bindPort(server, port)
            mock_bind.assert_called_with((HOST, port + 1))


class TestJokeDb(unittest.TestCase):
    def test_strictZipDict(self):
        """
        Avalia se a taxa setup/punchline está correta (1:1) 
        """
        try:
            mockDict = {setup: punchline for setup, punchline in zip(setups, punchlines, strict=True)}
        except ValueError:
            print("O número de setups não bate com o de punchlines!")


if __name__ == "__main__":
    unittest.main()
    