# falcon-agent

This is an agent port of open-falcon, just like the origin agent module, but can running on both windows and linux.

## Features

* besic data collection(cpu, mem, disk and etc.)
* process cpu, mem, nums collection
* http api to push
* single execute file, can working with windows service or linux systemd

## Installation

it is a golang classic project

``` shell
# set GOPATH and GOROOT
go get github.com/geekerlw/falcon-agent
cd $GOPATH/src/github.com/geekerlw/falcon-agent
go build -o falcon-agent.exe 	# for windows
go build -o falcon-agent 		# for linux
```



## Support metrics

### common

|  Counters   | Type  | Tag  |     Notes      |
| :---------: | :---: | :--: | :------------: |
| agent.alive | GAUGE |  /   | agent is alive |

### cpu

|     Counters     | Type  | Tag  |       Notes        |
| :--------------: | :---: | :--: | :----------------: |
|     cpu.user     | GAUGE |  /   |   cpu user time    |
|    cpu.system    | GAUGE |  /   |  cpu system time   |
|     cpu.idle     | GAUGE |  /   |   cpu idle time    |
|     cpu.nice     | GAUGE |  /   |   cpu nice time    |
|    cpu.iowait    | GAUGE |  /   |  cpu iowait time   |
|     cpu.irq      | GAUGE |  /   |    cpu irq time    |
|   cpu.softirq    | GAUGE |  /   |  cpu softirq time  |
|    cpu.steal     | GAUGE |  /   |   cpu steal time   |
|  cpu.guestnice   | GAUGE |  /   | cpu gusesnice time |
|    cpu.stolen    | GAUGE |  /   |  cpu stolen time   |
| cpu.used.percent | GAUGE |  /   |  cpu used percent  |

## memory

|       Counters        | Type  | Tag  |             Notes              |
| :-------------------: | :---: | :--: | :----------------------------: |
|    mem.swap.total     | GAUGE |  /   |       total swap memory        |
|     mem.swap.used     | GAUGE |  /   |        used swap memory        |
|     mem.swap.free     | GAUGE |  /   |        free swap memory        |
| mem.swap.used.percent | GAUGE |  /   |    swap memory used percent    |
|       mem.total       | GAUGE |  /   |      total virtual memory      |
|     mem.available     | GAUGE |  /   | total available virtual memory |
|       mem.used        | GAUGE |  /   |          used memory           |
|       mem.free        | GAUGE |  /   |          free memory           |
|   mem.used.percent    | GAUGE |  /   |      memory used percent       |

### Disk

|     Counters      | Type  |     Tag     |    Notes     |
| :---------------: | :---: | :---------: | :----------: |
|    disk.total     | GAUGE | diskpath=%s |    total     |
|     disk.free     | GAUGE | diskpath=%s |     free     |
|     disk.used     | GAUGE | diskpath=%s |     used     |
| disk.used.percent | GAUGE | diskpath=%s | used percent |

### Process

|     Counters     | Type  |            Tag            |       Notes        |
| :--------------: | :---: | :-----------------------: | :----------------: |
|     proc.num     | GAUGE | name=name,cmdline=cmdline |   process number   |
| proc.cpu.percent | GAUGE | name=name,cmdline=cmdline |  process cpu use   |
| proc.mem.percent | GAUGE | name=name,cmdline=cmdline | process memory use |

### Snmp

| Counters | Type  |         Tag          |           Notes            |
| :------: | :---: | :------------------: | :------------------------: |
| snmp.get | GAUGE | addr=address,oid=oid | get oid value from address |


## Configuration

* **heartbeat**: heartbeat server rpc address
* **transfer**: transfer rpc address
* **collector**: metric configs
* **ignore**: the metrics should be ignored

Refer to `cfg.example.json`, modify the file name to `cfg.json` :

```config
{
    "debug": true,
    "hostname": "",
    "ip": "",
    "plugin": {
        "enabled": false,
        "dir": "./plugin",
        "git": "https://github.com/open-falcon/plugin.git",
        "logs": "./logs"
    },
    "heartbeat": {
        "enabled": true,
        "addr": "127.0.0.1:6030",
        "interval": 60,
        "timeout": 1000
    },
    "transfer": {
        "enabled": true,
        "addrs": [
            "127.0.0.1:8433",
            "127.0.0.1:8433"
        ],
        "interval": 60,
        "timeout": 1000
    },
    "http": {
        "enabled": true,
        "listen": ":1988",
        "backdoor": false
    },
    "collector": {
        "ifacePrefix": ["eth", "em"],
        "mountPoint": [],
        "snmpAddr": `192.168.1.1`,	// snmp target address
        "snmpOids": ["1.3.6.1.2.1.2.1.0", "1.3.6.1.2.1.1.3.0"]	// snmp target oids
    },
    "default_tags": {
    },
    "ignore": {
    }
}

```

## License

This software is licensed under the Apache License. See the LICENSE file in the top distribution directory for the full license text.