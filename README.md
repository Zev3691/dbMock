## dbMock

### 根据mock.json文件生产对应的假数据

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


