### 功能描述

获取模板列表

### 请求参数

{{ common_args_desc }}

#### 接口参数

| 字段        |  类型     | 必选   |  描述    |
|-------------|-----------|--------|----------|
| biz_id      |  string   | 是     | 业务ID   |
| page        |  object   | 是     | 分页设置 |

#### page

| 字段         |  类型  | 必选   |  描述      |
|--------------|--------|--------|------------|
| return_total |  bool  | 否     | 是否返回总记录条数, 默认不返回 |
| start        |  int   | 是     | 记录开始位置 |
| limit        |  int   | 是     | 每页限制条数,最大100 |

### 请求参数示例

```json
{
    "bk_app_code": "xxx",
    "bk_app_secret": "xxx",
    "bk_token": "xxx",
    "biz_id": "xxx",
    "page": {
        "start": 0,
        "limit": 100
    }
}
```

### 返回结果示例

```json
{
    "result": true,
    "code": 0,
    "message": "OK",
    "data": {
        "total_count": 1,
        "info": [
            {
                "biz_id": "XXX",
                "template_id": "T-0b67a798-e9c1-11e9-8c23-525400f99278",
                "name": "server.yaml.template",
                "filt_name": "server.yaml",
                "file_path": "/etc",
                "user": "root",
                "user_group": "root",
                "file_privilege": "0755",
                "file_format": "unix",
                "file_mode": 1,
                "engine_type": 1,
                "creator": "melo",
                "last_modify_by": "melo",
                "memo": "my first config",
                "state": 0,
                "created_at": "2019-07-29 11:57:20",
                "updated_at": "2019-07-29 11:57:20"
            }
        ]
    }
}
```

### 返回结果参数

#### data

| 字段        | 类型      | 描述      |
|-------------|-----------|-----------|
| total_count | int       | 当前规则能匹配到的总记录条数 |
| info        | array     | 查询返回的数据 |

#### data.info[n]

| 字段           | 类型      | 描述    |
|----------------|-----------|---------|
| biz_id         |  string   | 业务ID  |
| template_id    |  string   | 模板ID  |
| name           |  string   | 模板名称|
| file_name      |  string   | 配置名称, 例如server.yaml   |
| file_path      |  string   | 配置相对路径, 例如/etc |
| user           |  string   | 归属用户信息, 例如root|
| user_group     |  string   | 归属用户组信息, 例如root |
| file_privilege |  string   | 文件权限，例如0755 |
| file_format    |  string   | 文件格式，例如unix |
| file_mode      |  integer  | 配置类型, 0: 文本文件  1: 二进制文件  2: 模板文件 |
| engine_type    |  integer  | 渲染引擎类型, 0: 非模板 1: Golang模板 2: Py Mako引擎 3: 外部引擎 |
| memo           |  string   | 备注 |
| state          |  integer  | 状态 默认0: 正常 |
| creator        |  string   | 创建者 |
| last_modify_by |  string   | 修改者 |
| created_at     |  string   | 创建时间 |
| updated_at     |  string   | 更新时间 |