# go-excel
go excel

# go excel 处理,返回结构体数组
核心部分 没有使用任何 第三方包,引入第三方包都是测试和转换使用的

# 使用方式
项目中执行引入包
```bash
go get github.com/liangzibo/go-excel
```
结构体TAG 中 `json`,`index` 必须存在

json: 字段名称

index: 索引名称

name: 中文名称

# 案例
```go
type ExcelTest struct {
Name     string    `json:"name" name:"名称" index:"0"`
Age      int64     `json:"age" name:"年龄" index:"1"`
Score    int64     `json:"score" name:"分数" index:"2"`
Birthday time.Time `json:"birthday" name:"生日" index:"4"`
}
```
如果要执行案例,先引入案例依赖
```shell
go get  github.com/360EntSecGroup-Skylar/excelize/v2
go get  github.com/mitchellh/mapstructure
```
具体案例如下 或 看测试文件 

```go
	xlsx, err := excelize.OpenFile("./lzbExcel/excel.xlsx")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    // Get all the rows in a sheet.
    rows, err := xlsx.GetRows("Sheet1")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    //结果在  arr 中
    var arr []ExcelTest
    err = lzbExcel.NewExcelStructDefault().SetPointerStruct(&ExcelTest{}).RowsProcess(rows, func(maps map[string]interface{}) error {
        var ptr ExcelTest
		// map 转 结构体
        if err2 := mapstructure.Decode(maps, &ptr); err2 != nil {
            return err2
        }
        arr = append(arr, ptr)
        return nil
    })
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(arr)
    
    //结果在  arr 中
    var arr2 []ExcelTest
    //StartRow 开始行,索引从 0开始
    //IndexMax  索引最大行,如果 结构体中的 index 大于配置的,那么使用结构体中的
    err = lzbExcel.NewExcelStruct(1, 10).SetPointerStruct(&ExcelTest{}).RowsProcess(rows, func(maps map[string]interface{}) error {
        var ptr ExcelTest
		// map 转 结构体
        if err2 := mapstructure.Decode(maps, &ptr); err2 != nil {
            return err2
        }
        arr2 = append(arr2, ptr)
        return nil
    })
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(arr)
```

打印结果
```bash
[{张三 30 100 1999-02-01 15:20:31 +0800 CST} {李四 31 99 0001-01-01 00:00:00 +0000 UTC} {王五 33 0 0001-01-01 00:00:00 +0000 UTC}]
[{张三 30 100 1999-02-01 15:20:31 +0800 CST} {李四 31 99 0001-01-01 00:00:00 +0000 UTC} {王五 33 0 0001-01-01 00:00:00 +0000 UTC}]
```