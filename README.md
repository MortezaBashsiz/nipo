# Welcome to NIPO

Nipo is going to be a powerful, fast, multi-thread, clustered and in-memory key-value database, written by GO.
With several mathematical and aggregation functionalities on batch of keys and values.

# Config file
## global
Global section defines some global parameters

`authorization (string): [true/false]`

Defines that the clients must work with token or not. If set "true" you have to define users section

`master (string): [true/false]`

Defines that this server has some slaves. If set "true" you have to define slaves section

`checkinterval (int):`

Defines the interval of slaves healthcheck in milliseconds

    global:  
      authorization: "false"
      master: "true"
      checkinterval: 1000



## slaves
Slave section defines parameter about slaves of this server

`id (int)` : defines the id of slave. Master will sync the slaves by id priority.

`ip (string)` : is the IP of slave

`port (string)` : is the listen port of destination IP

`authorization (string) : [true/false]` defines if the destination slave uses token or not

`token (string)` : in case of authorization is true, you need to define token


    slaves:
      - slave:
        id : 1
        ip : "127.0.0.1"
        port : "2324"
        authorization: "false"
        token: "061b30a7-1a12-4280-8e3c-6bc9a19b1683"
      - slave:
        id : 2
        ip : "127.0.0.1"
        port : "2325"
        authorization: "false"
        token: "061b30a7-1a12-4280-8e3c-6bc9a19b1683"


## proc
Proc section defines parameters for multi-threading and multi-processing

`cores (int)` : the count of cores you want to used by nipo

`threads (int)` : the count of threads you want to created by nipo

**NOTE** : the best practice is using threads two times of cores

    proc:
      cores: 2
      threads: 4


## listen
At this section you can configure your server side listen IP and PORT, currently only TCP is allowed.

    listen:
      ip: "0.0.0.0"
      port: "2323"
      protocol: "tcp"

## log
Log section defines parameters for logging

    level (int) :
      0 - no log
      1 - info
      2 - debug

`path (string)` : defines the path of log file

    log:
      level: 1
      path: "/tmp/nipo.log"

## users
Users section defines parameters for authorization. 
If authorization in global section is true, this section had to be defined
you can define several users

`name (string)` : just is metadata for name of user

`token (string)` : used for authorization

`keys (string)` : the regex of keys which user should have access.
                if you have several regexes you can separate them with delimiter "||"

`cmds (string)` : the list of commands that user should have access to execute
                if you have several commands you can separate them with delimiter "||"

    users:
      - user:
        name: "admin"
        token: "061b30a7-1a12-4280-8e3c-6bc9a19b1683"
        keys: ".*"
        cmds: "all"
      - user:
        name: "readonly"
        token: "0517376d-49c1-40eb-a8fc-fd73b70a4ce9"
        keys: "name.*||.*log.*"
        cmds: "get||select||avg"

# CLI

To introduce with CLI `nipocli` please visit the [LINK](nipocli).

# GO library

To introduce how to use nipo with GO please visit the [LINK](nipolib/go).