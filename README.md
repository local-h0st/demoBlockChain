# demo block chain project readme
## 0x00 开发环境配置
由于采用的是docker上的code-server image，因此采取的方式是挂载home目录到镜像的project目录，同时修改容器内的环境变量（配置GOROOT和GOPATH），实现一键运行
## 0x01 进度
* 2022.12.6 实现了proof of work机制，并且基于pow实现了区块的产出
* 2022.12.9 拆分main.go到package dem0chain里面，逻辑更加清晰

## 0x02 下一步
实现链的存储（序列化后存入数据库）以及从数据库中恢复
