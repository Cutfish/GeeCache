主要的原理是：
一个api服务（选取一个节点开启api服务），n个节点（用户感知不到），通过…/api?key=Tom访问api
调用group.get：如果api服务所在的节点的maincache无数据，就调用一致性hash算法找到所在的远程节点