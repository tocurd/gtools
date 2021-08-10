### gtools是一款集合型的工具包


#### Array
```
// 查找某个值是否在数组内
gtools.Array.Find(target interface{}, value interface{}) bool

// 数组转字符串
gtools.Array.Join(a []string, sep string) string

// 获取数组的指定字段
gtools.Array.Field(array interface{}, key ...string) (result []interface{}, err error)

// 将某个字段栏目里的数据提取
gtools.Array.Column(array interface{}, key string) (result []interface{}, err error)
```

#### base64
#### cmd
#### dir
#### file
#### hash
#### interface
#### ip
#### log
#### network
#### path
#### quick
#### rand
#### string
#### struct
#### time
#### turn
#### verify
#### yaml
#### zip
