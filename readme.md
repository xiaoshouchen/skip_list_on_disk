# skip list

本项目用golang实现一个skip list，支持超过内存大小的存储级别

## 实现方式

将key存储到内存中，将value写入硬盘，skip list 原本的 value 存储offset，通过offset直接从硬盘上读取消息

## 文件架构

- disk_io 主要负责硬盘的读写
- sorted_string_table 是对外开放接口
- skip_list 跳跃表，为sst提供数据结构上的支持