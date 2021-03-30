import socket

class nipoSocket:

    def __init__(self, sock=None):
        if sock is None:
            self.sock = socket.socket(
                            socket.AF_INET, socket.SOCK_STREAM)
        else:
            self.sock = sock

    def connect(self, host, port):
        self.sock.connect((host, port))
        err = self.sock.recv(1024)
        if err != None :
            print(repr(err))


    def send(self, msg):
        self.sock.send(bytes(msg , 'utf-8') )
        err = self.sock.recv(1024)
        if err != None :
            print(repr(err))





