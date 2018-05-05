

## 基于etcd3实现的gRPC负载均衡


- 代码非我原创，在此文的基础之上修改： [gRPC服务发现&负载均衡](https://segmentfault.com/a/1190000008672912)
- 原文中的代码有一些问题
    1. 判断serviceKey是否存在有问题
    2. 注册时重复PUT，KeepAliveOnce
    3. UnRegister函数bug修复
- 同时参考了<https://github.com/wwcd/grpc-lb>,实现方式略有不同