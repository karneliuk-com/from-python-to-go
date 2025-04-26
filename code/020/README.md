# Code for blog 020

## Use case
Concurency in Python and Golang
```bash
$ export AUTOMATION_INVENTORY_URL="https://demo.netbox.dev"
$ export AUTOMATION_INVENTORY_TOKEN="token"
$ export AUTOMATION_USER="user"
$ export AUTOMATION_PASS="password"
```

Extra dependencies are required.

### Python
Requirements:
```Bash
$ pip install pygnmi httpx
```

Execution result:
```bash
$ python main.py
ssl_target_name_override is applied, should be used for testing only!
ssl_target_name_override is applied, should be used for testing only!
config_result={'timestamp': 1745696412856245684, 'prefix': None, 'response': [{'path': 'interfaces', 'op': 'UPDATE'}]}
config_result={'timestamp': 1745696415292591043, 'prefix': None, 'response': [{'path': 'interfaces', 'op': 'UPDATE'}]}
Device: ka-blog-001
Config: [('/interfaces', {'interface': [{'name': 'Loopback 23', 'config': {'name': 'Loopback 23', 'description': 'Test-gnmi-python-23'}}]})]
Impact: ***
---
***************
*** 1,4 ****
! notification/0/timestamp = 1745696415279571842
  notification/0/prefix = None
  notification/0/alias = None
  notification/0/atomic = False
--- 1,4 ----
! notification/0/timestamp = 1745696415677702589
  notification/0/prefix = None
  notification/0/alias = None
  notification/0/atomic = False
***************
*** 48,58 ****
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/openconfig-if-ip:ipv4/addresses/address/0/config/prefix-length = 31
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/openconfig-if-ip:ipv4/addresses/address/0/ip = 10.0.0.1
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/openconfig-if-ip:ipv4/config/enabled = True
! notification/0/update/0/val/openconfig-interfaces:interface/3/config/description = Test-gnmi-python-2
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/loopback-mode = FACILITY
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/name = Loopback23
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/type = iana-if-type:softwareLoopback
  notification/0/update/0/val/openconfig-interfaces:interface/3/name = Loopback23
! notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/config/description = Test-gnmi-python-2
  notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/index = 0
  notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/openconfig-if-ip:ipv4/config/enabled = True
--- 48,58 ----
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/openconfig-if-ip:ipv4/addresses/address/0/config/prefix-length = 31
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/openconfig-if-ip:ipv4/addresses/address/0/ip = 10.0.0.1
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/openconfig-if-ip:ipv4/config/enabled = True
! notification/0/update/0/val/openconfig-interfaces:interface/3/config/description = Test-gnmi-python-23
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/loopback-mode = FACILITY
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/name = Loopback23
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/type = iana-if-type:softwareLoopback
  notification/0/update/0/val/openconfig-interfaces:interface/3/name = Loopback23
! notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/config/description = Test-gnmi-python-23
  notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/index = 0
  notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/openconfig-if-ip:ipv4/config/enabled = True
Timestamp: 2025-04-26 20:40:17.240665
Device: ka-blog-002
Config: [('/interfaces', {'interface': [{'name': 'Loopback 23', 'config': {'name': 'Loopback 23', 'description': 'Test-gnmi-python-23'}}]})]
Impact: ***
---
***************
*** 818,824 ****
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/openconfig-if-ip:ipv6/state/mtu = 1500
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/state/index = 0
! notification/0/update/0/val/openconfig-interfaces:interface/5/config/description = Test-gnmi-python-2
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/arista-intf-augments:load-interval = 300
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/loopback-mode = True
--- 818,824 ----
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/openconfig-if-ip:ipv6/state/mtu = 1500
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/state/index = 0
! notification/0/update/0/val/openconfig-interfaces:interface/5/config/description = Test-gnmi-python-23
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/arista-intf-augments:load-interval = 300
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/loopback-mode = True
***************
*** 833,839 ****
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/name = Loopback23
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/openconfig-vlan:tpid = openconfig-vlan-types:TPID_0X8100
! notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/description = Test-gnmi-python-2
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/index = 0
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/index = 0
--- 833,839 ----
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/name = Loopback23
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/openconfig-vlan:tpid = openconfig-vlan-types:TPID_0X8100
! notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/description = Test-gnmi-python-23
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/index = 0
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/index = 0
Timestamp: 2025-04-26 20:40:17.138377
```

### Go
Requirements:
```bash
$ go get github.com/google/go-cmp/cmp
$ go get github.com/openconfig/gnmic/api
$ go get google.golang.org/protobuf/encoding/prototext
```

Execution result:
```bash
$ go run .
2025/04/26 20:56:36 response: {
  path: {
    elem: {
      name: "interfaces"
    }
  }
  op: UPDATE
}
timestamp: 1745697392374689242

2025/04/26 20:56:36 response: {
  path: {
    elem: {
      name: "interfaces"
    }
  }
  op: UPDATE
}
timestamp: 1745697394881050173

Config: {/interfaces map[interface:[map[config:map[description:Test-gnmi-golang-23 name:Loopback 23] name:Loopback 23]]]}
Impact:   main.OpenConfigInterfaces{
        Interface: []main.OpenConfigInterface{
                ... // 3 identical elements
                {Name: "Loopback51", Config: {Name: "Loopback51", Description: "pytest-update-test-33"}},
                {Name: "Loopback0", Config: {Name: "Loopback0", Enabled: true}},
                {
                        Name: "Loopback23",
                        Config: struct{ Name string "xml:\"name,omitempty\" json:\"name,omitempty\""; Description string "xml:\"description,omitempty\" json:\"description,omitempty\""; Enabled bool "xml:\"enabled,omitempty\" json:\"enabled,omitempty\"" }{
                                Name:        "Loopback23",
-                               Description: "Test-gnmi-golang-3",
+                               Description: "Test-gnmi-golang-23",
                                Enabled:     true,
                        },
                },
        },
  }

Timestamp: 2025-04-26 20:56:36.703237967 +0100 BST m=+1.125671022
Config: {/interfaces map[interface:[map[config:map[description:Test-gnmi-golang-23 name:Loopback 23] name:Loopback 23]]]}
Impact:   main.OpenConfigInterfaces{
        Interface: []main.OpenConfigInterface{
                {Name: "Management1", Config: {Name: "Management1"}},
                {Name: "Ethernet2", Config: {Name: "Ethernet2"}},
                {Name: "Ethernet1", Config: {Name: "Ethernet1"}},
                {
                        Name: "Loopback23",
                        Config: struct{ Name string "xml:\"name,omitempty\" json:\"name,omitempty\""; Description string "xml:\"description,omitempty\" json:\"description,omitempty\""; Enabled bool "xml:\"enabled,omitempty\" json:\"enabled,omitempty\"" }{
                                Name:        "Loopback23",
-                               Description: "Test-gnmi-golang-3",
+                               Description: "Test-gnmi-golang-23",
                                Enabled:     false,
                        },
                },
        },
  }

Timestamp: 2025-04-26 20:56:37.002338539 +0100 BST m=+1.424771596
```
