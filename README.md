# loginservice


1.提交到任意GitHub，git提交需要注意有意义，不要随意提交 
2.注意安全可能性
3.注意选择的技术栈，说明为什么这样选择
4.注意权限验证的便捷性，使得其他地方也能使用
5.考虑性能、可扩展性，说明如何达到的


注意安全可能性
* jwt
* bcrypt(pwd) to db
* https
* 验证码
* 登入频次，锁定账号，暴力破解
* 部署相关 WAP等
* 密码强度
* SQL 注入等


注意选择的技术栈，说明为什么这样选择
* golang
  * 轻量
  * 快速
  * 性能好
  * 部署简单
* vue
  * 轻量
* mobile
  * xxx

注意权限验证的便捷性，使得其他地方也能使用
* 登入用户密码匹配 or 用户角色权限管理
* 登入用户密码匹配
  * 统一接口
  * 独立部署
  * 多实例部署 + 服务发现
  * 负载策略


考虑性能、可扩展性，说明如何达到的
* 存储
  * 按照需求规划分库分表
  * 索引优化
* 服务
  * 无状态
  * 多实例 + 服务发现
  * 负载均衡