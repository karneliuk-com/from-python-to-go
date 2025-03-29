# Code for blog 017

## Use case
Intreraction with device using NETCONF. Put credentails to environment variables.
```bash
$ export AUTOMATION_USER="user"
$ export AUTOMATION_PASS="password"
```

Extra dependencies are required.

### Python
Requirements:
```Bash
$ pip install pyyaml scrapli scrapli-netconf xmltodict ssh2-python
```

Execution result:
```bash
$ python main.py -i ../data/inventory.yaml
Device: go-blog-nexus
Config: {'config': {'interfaces': {'@xmlns': 'http://openconfig.net/yang/interfaces', 'interface': [{'name': 'Loopback 23', 'config': {'name': 'Loopback 23', 'description': 'Test-netconf-python-2'}}]}}}
Impact: ***
---
***************
*** 1,6 ****
  rpc-reply/@xmlns = urn:ietf:params:xml:ns:netconf:base:1.0
! rpc-reply/@message-id = 101
! rpc-reply/data/@time-modified = 2025-03-22T16:53:24.80910925Z
  rpc-reply/data/interfaces/@xmlns = http://openconfig.net/yang/interfaces
  rpc-reply/data/interfaces/interface/0/name = Management1
  rpc-reply/data/interfaces/interface/0/config/enabled = true
--- 1,6 ----
  rpc-reply/@xmlns = urn:ietf:params:xml:ns:netconf:base:1.0
! rpc-reply/@message-id = 104
! rpc-reply/data/@time-modified = 2025-03-22T16:53:25.701759874Z
  rpc-reply/data/interfaces/@xmlns = http://openconfig.net/yang/interfaces
  rpc-reply/data/interfaces/interface/0/name = Management1
  rpc-reply/data/interfaces/interface/0/config/enabled = true
***************
*** 134,140 ****
  rpc-reply/data/interfaces/interface/2/subinterfaces/subinterface/ipv6/config/enabled = false
  rpc-reply/data/interfaces/interface/2/subinterfaces/subinterface/ipv6/config/mtu = 1500
  rpc-reply/data/interfaces/interface/3/name = Loopback23
! rpc-reply/data/interfaces/interface/3/config/description = Test-netconf-golang-2
  rpc-reply/data/interfaces/interface/3/config/enabled = true
  rpc-reply/data/interfaces/interface/3/config/load-interval/@xmlns = http://arista.com/yang/openconfig/interfaces/augments
  rpc-reply/data/interfaces/interface/3/config/load-interval/#text = 300
--- 134,140 ----
  rpc-reply/data/interfaces/interface/2/subinterfaces/subinterface/ipv6/config/enabled = false
  rpc-reply/data/interfaces/interface/2/subinterfaces/subinterface/ipv6/config/mtu = 1500
  rpc-reply/data/interfaces/interface/3/name = Loopback23
! rpc-reply/data/interfaces/interface/3/config/description = Test-netconf-python-2
  rpc-reply/data/interfaces/interface/3/config/enabled = true
  rpc-reply/data/interfaces/interface/3/config/load-interval/@xmlns = http://arista.com/yang/openconfig/interfaces/augments
  rpc-reply/data/interfaces/interface/3/config/load-interval/#text = 300
***************
*** 148,154 ****
  rpc-reply/data/interfaces/interface/3/hold-time/config/down = 0
  rpc-reply/data/interfaces/interface/3/hold-time/config/up = 0
  rpc-reply/data/interfaces/interface/3/subinterfaces/subinterface/index = 0
! rpc-reply/data/interfaces/interface/3/subinterfaces/subinterface/config/description = Test-netconf-golang-2
  rpc-reply/data/interfaces/interface/3/subinterfaces/subinterface/config/enabled = true
  rpc-reply/data/interfaces/interface/3/subinterfaces/subinterface/config/index = 0
  rpc-reply/data/interfaces/interface/3/subinterfaces/subinterface/ipv4/@xmlns = http://openconfig.net/yang/interfaces/ip
--- 148,154 ----
  rpc-reply/data/interfaces/interface/3/hold-time/config/down = 0
  rpc-reply/data/interfaces/interface/3/hold-time/config/up = 0
  rpc-reply/data/interfaces/interface/3/subinterfaces/subinterface/index = 0
! rpc-reply/data/interfaces/interface/3/subinterfaces/subinterface/config/description = Test-netconf-python-2
  rpc-reply/data/interfaces/interface/3/subinterfaces/subinterface/config/enabled = true
  rpc-reply/data/interfaces/interface/3/subinterfaces/subinterface/config/index = 0
  rpc-reply/data/interfaces/interface/3/subinterfaces/subinterface/ipv4/@xmlns = http://openconfig.net/yang/interfaces/ip
Timestamp: 2025-03-22 16:53:26.975625
```

### Go
Requirements:
```bash
$ go get github.com/google/go-cmp/cmp
$ go get gopkg.in/yaml.v3
$ go get github.com/scrapli/scrapligo
```

Execution result:
```bash
$ go run . -i ../data/inventory.yaml
Config: {{http://openconfig.net/yang/interfaces interfaces} [{Loopback 23 {Loopback 23 Test-netconf-golang-2 true}}]}
Impact:   main.RPCResponse{
        XMLName: {Space: "urn:ietf:params:xml:ns:netconf:base:1.0", Local: "rpc-reply"},
        Data: struct{ Interfaces main.OpenConfigInterfaces }{
                Interfaces: main.OpenConfigInterfaces{
                        XMLName: {Space: "http://openconfig.net/yang/interfaces", Local: "interfaces"},
                        Interface: []main.OpenConfigInterface{
                                {Name: "Management1", Config: {Name: "Management1", Enabled: true}},
                                {Name: "Ethernet2", Config: {Name: "Ethernet2", Enabled: true}},
                                {Name: "Ethernet1", Config: {Name: "Ethernet1", Enabled: true}},
                                {
                                        Name: "Loopback23",
                                        Config: struct{ Name string "xml:\"name,omitempty\""; Description string "xml:\"description,omitempty\""; Enabled bool "xml:\"enabled\"" }{
                                                Name:        "Loopback23",
-                                               Description: "Test-netconf-python-2",
+                                               Description: "Test-netconf-golang-2",
                                                Enabled:     true,
                                        },
                                },
                                {Name: "Loopback0", Config: {Name: "Loopback0", Enabled: true}},
                                {Name: "Loopback51", Config: {Name: "Loopback51", Description: "pytest-update-test-33"}},
                        },
                },
        },
  }
Timestamp: 2025-03-22 18:28:01.696317104 +0000 GMT m=+2.236461044
```
