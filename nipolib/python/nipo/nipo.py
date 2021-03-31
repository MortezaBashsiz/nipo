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

    def send(self, msg):
        self.sock.sendall(bytes(msg , 'utf-8') )
        response = self.sock.recv(1024)
        if response != None :
            print (repr(response))

class Config():
    def __init__(self, token, server, port):
        self.token = token
        self.server = server
        self.port = port

def CreateConfig(token, server, port):
    global conf
    conf = Config(token, server, port )

def CreateConnection() :
    global sock
    sock = nipoSocket()

def OpenConnection( ): 
    CreateConnection()
    sock.connect(conf.server, conf.port)

def ping():    
    OpenConnection()
    string = conf.token+" "+"ping"
    sock.send(string)

