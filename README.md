## dbMock

### 根据json文件生产对应的假数据

#### 使用：
```
1. 首先安装go环境，项目使用go.mod,需配置相关选项 //百度可得
2. 进入项目跟目录，执行： go build -o dbMock，等待编译完成
3. 编译完成后的到二进制文件，使用 ./dbMock -h可查看帮助
ps: 
    1. 帮助中的path参数路径使用绝对路径
    2. 密码建议使用''包裹
    3. 所有操作在项目跟目录执行
eg:
    ./dbmock -u root -p 'root' -path /home/project/dbMock/mock.json
```


```json
{
    // 数据库host
    "host": "192.168.2.122",
    // 数据库端口
    "port": "3306",
    //数据库databasename
    "dbname": "pbx",
    // 表名
    "table_name": "pbx.clip_list",
    // 表结构，目前支持varchar，char，时间戳三种类型的数据
    // varchar，char 长度为10，不可定义
    "cols":{
        "name": "varchar",
        "uuid": "char",
        "trunk": "varchar",
        "update_time": "int"
    },
    // 需要生产的数据量
    "max_count": 100000
}
```
