# go-excel
go excel

待转换的数据 类型 为  [][]string,都可以转换为结构体

# go excel 处理,返回结构体数组
核心部分 没有使用任何 第三方包,引入第三方包都是测试和转换使用的

待转换的数据 类型 为  [][]string,都可以转换为结构体

# 使用方式
项目中执行引入包
```bash
go get github.com/liangzibo/go-excel
```
结构体TAG 中 `json`,`index` 必须存在

json: 字段名称

index: 索序号

name: 名称

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

excel 读取 出的格式如下
```go
var demoAll = [][]string{
		{"赵三",
			"30",
			"100",
			"1970-01-01 12:50:01",
			"1980-11-21 15:20:01"},
		{"赵六",
			"40",
			"30",
			"1970-01-01 12:50:01",
			"1980-11-21 15:20:01"},
	}
```

```go
	xlsx, err := excelize.OpenFile("./excel.xlsx")
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
    err = lzbExcel.NewExcelStructDefault().SetPointerStruct(&ExcelTest{}).RowsAllProcess(rows, func(maps map[string]interface{}) error {
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
    //rows 为[][]string 类型
    //结果在  arr 中
    var arr2 []ExcelTest
    //StartRow 开始行,索引从 0开始
    //IndexMax  索引最大行,如果 结构体中的 index 大于配置的,那么使用结构体中的
    err = lzbExcel.NewExcelStruct(1, 10).SetPointerStruct(&ExcelTest{}).RowsAllProcess(rows, func(maps map[string]interface{}) error {
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


# Go excel processing, return structure array
The core part does not use any third-party packages, and the introduction of third-party packages is used for testing and transformation

The data type to be converted is [][]string, which can be converted to struct


# Usage
Execute the import package in the project
```bash
go get github.com/liangzibo/go-excel
```
struct TAG 中 `json`,`index` Must exist

json: Field name

index: Index ordinal

name: name


# case
```go
type ExcelTest struct {
Name     string    `json:"name" name:"名称" index:"0"`
Age      int64     `json:"age" name:"年龄" index:"1"`
Score    int64     `json:"score" name:"分数" index:"2"`
Birthday time.Time `json:"birthday" name:"生日" index:"4"`
}
```
If you want to execute a case, introduce the case dependency first
```shell
go get  github.com/360EntSecGroup-Skylar/excelize/v2
go get  github.com/mitchellh/mapstructure
```
Specific cases are as follows or see the test file

```go
	xlsx, err := excelize.OpenFile("./excel.xlsx")
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
    //result  arr
    var arr []ExcelTest
    err = lzbExcel.NewExcelStructDefault().SetPointerStruct(&ExcelTest{}).RowsAllProcess(rows, func(maps map[string]interface{}) error {
        var ptr ExcelTest
		// map to struct
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
    //rows  :=[][]string{}
    //result  arr 
    var arr2 []ExcelTest
    //StartRow starts row, index starts from 0
    //IndexMax index the maximum row, if the index in the structure is larger than the configured, then use the structure in the
    err = lzbExcel.NewExcelStruct(1, 10).SetPointerStruct(&ExcelTest{}).RowsAllProcess(rows, func(maps map[string]interface{}) error {
        var ptr ExcelTest
		// map to struct
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
    fmt.Println("arr2")
    fmt.Println(arr2)
    /////////////
    ////////////
    ////////////
    var arr3 []ExcelTest
    excel := lzbExcel.NewExcelStruct(1, 10).SetPointerStruct(&ExcelTest{})
    for i, row := range rows {
        //If the index is less than the set start row, skip
        //如果 索引 小于 已设置的 开始行,那么跳过
        if i < excel.StartRow {
            continue
        }
        //单行处理
        m, err3 := excel.Row(row)
        if err3 != nil {
            fmt.Println(err3)
        }
        var tmp ExcelTest
        // map 转 结构体
        if err2 := mapstructure.Decode(m, &tmp); err2 != nil {
            fmt.Println(err2)
        } else {
            arr3 = append(arr3, tmp)
        }
    }
    fmt.Println("arr3")
    fmt.Println(arr3)
    //单行处理
    var demo = []string{
        "赵六",
        "40",
        "30",
        "1970-01-01 12:50:01",
        "1980-11-21 15:20:01",
    }
    row, err := lzbExcel.NewExcelStructDefault().SetPointerStruct(&ExcelTest{}).Row(demo)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(row)
    var demoStruct ExcelTest
    // map 转 结构体
    if err2 := mapstructure.Decode(row, &demoStruct); err2 != nil {
        fmt.Println(err2)
    }
    fmt.Println(demoStruct)
```

打印结果
```bash
[{张三 30 100 1999-02-01 15:20:31 +0800 CST} {李四 31 99 0001-01-01 00:00:00 +0000 UTC} {王五 33 0 0001-01-01 00:00:00 +0000 UTC}]
[{张三 30 100 1999-02-01 15:20:31 +0800 CST} {李四 31 99 0001-01-01 00:00:00 +0000 UTC} {王五 33 0 0001-01-01 00:00:00 +0000 UTC}]
```