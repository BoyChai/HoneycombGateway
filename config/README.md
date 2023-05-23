```yaml
# 全局配置
global:
	encryption: "none | gob"
# 内网穿透配置
penetrate:
  - listen: [ 地址池中的地址名称 ]
  - nodes: 
    - name: [NAME]
      protocol: [ TCP | UDP ]
      addr: [ 地址池中的地址名称 ]
# 内网穿透客户端配置
client:
  - server: [ 地址池中的地址名称 ]
  - service:
    - name: [NAME]
      protocol: [ TCP | UDP ]
      addr: [ 地址池中的地址名称 ]
    - name: [NAME]
      protocol: [ TCP | UDP ]
      addr: [ 地址池中的地址名称 ]
# 流量转发配置
forward:
  - name: [NAME]
    protocol: [ TCP | UDP ]
    # 填写地址名称(address里面配置)
    listen: [ 地址池中的地址名称 ]  
    target: [ 地址池中的地址名称 ]

# 负载配置
load:
  - name: [NAME]
    protocol: [ TCP | http ]
    servers:
      - [ 地址池中的地址名称 ]
      - [ 地址池中的地址名称 ]
# 地址池
address:
  - name: addr1
    addr: "127.0.0.1"
    port: "80"
  - name: addr2
    addr: "localhost"
    port: "81"
  - name: port2
    port: "8080"
```