package lzbExcel

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	DATE_PATTERN      = "2006-01-02"
	DATE_TIME_PATTERN = "2006-01-02 15:04:05"
)

//字段
type ExcelFields struct {
	Name      string            //名称
	Index     int               //索引  从0 开始
	Field     string            //json 字段名称
	FieldType string            //字段类型
	Tags      map[string]string // 保存所有tags
}

//综合
type ExcelStruct struct {
	MapIndex map[int]string         //按照 index 排序
	IndexMax int                    // index 最大
	Fields   map[string]ExcelFields //所有字段名
	StartRow int                    //第几行开始为具体数据
	Err      error                  //错误
}

//默认 从第一行开始,索引从 0开始
func NewExcelStructDefault() *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = 1
	n.IndexMax = 10
	return n
}

//StartRow 开始行,索引从 0开始
//IndexMax  索引最大行,如果 结构体中的 index 大于配置的,那么使用结构体中的
func NewExcelStruct(StartRow, IndexMax int) *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = StartRow
	n.IndexMax = IndexMax
	return n
}

type Callback func(maps map[string]interface{}) error


// 结构体 指针
func (c *ExcelStruct) SetPointerStruct(ptr interface{}) *ExcelStruct {
	// 获取入参的类型
	t := reflect.TypeOf(ptr)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		c.Err = fmt.Errorf("参数应该为结构体指针")
		return c
	}
	// 取指针指向的结构体变量
	v := reflect.ValueOf(ptr).Elem()
	// 解析字段
	for i := 0; i < v.NumField(); i++ {
		// 取tag
		fieldInfo := v.Type().Field(i)
		//
		fields := ExcelFields{}
		tag := fieldInfo.Tag
		// 解析label tag
		fields.Field = tag.Get("json")
		if fields.Field == "" {
			fields.Field = fieldInfo.Name
		}
		fields.Name = tag.Get("name")
		if fields.Name == "" {
			fields.Name = fieldInfo.Name
		}
		index := 0
		indexStr := tag.Get("index")
		if indexStr != "" {
			index, _ = strconv.Atoi(indexStr)
		}
		//如果索引大 那么赋值
		if c.IndexMax < index {
			c.IndexMax = index
		}
		fields.Index = index
		fields.FieldType = fieldInfo.Type.String()
		m := make(map[string]string)
		m["json"] = fields.Field
		m["name"] = fields.Name
		m["index"] = strconv.Itoa(i)
		fields.Tags = m
		//
		if c.Fields == nil {
			c.Fields = make(map[string]ExcelFields)
		}
		c.Fields[fields.Field] = fields
		if c.MapIndex == nil {
			c.MapIndex = make(map[int]string)
		}
		c.MapIndex[index] = fields.Field
		// 解析uppercase tag
		//value := fmt.Sprintf("%v", v.Field(i))
		//if fieldInfo.Type.Kind() == reflect.String {
		//	uppercase := tag.Get("uppercase")
		//	if uppercase == "true" {
		//		value = strings.ToUpper(value)
		//	} else {
		//		value = strings.ToLower(value)
		//	}
		//}
	}
	return c
}

//行处理
func (c *ExcelStruct) RowsProcess2(rows [][]string, callback Callback) error {
	if c.Fields == nil {
		return fmt.Errorf("请填写结构体指针")
	}
	if c.Err != nil {
		return c.Err
	}
	//data := []interface{}{}
	for index, row := range rows {
		//如果 索引 小于 已设置的 开始行,那么跳过
		if index < c.StartRow {
			continue
		}
		maps := make(map[string]interface{})
		for i, colCell := range row {
			//不能判断空值,否则
			if len(colCell) < 1 {
				continue
			}
			//判断键名是否存在
			if field, ok := c.MapIndex[i]; ok {
				maps[field] = ""
				//类型转换
				fields := c.Fields[field]
				//字符
				if fields.FieldType == "string" {
					maps[field] = colCell
					continue
				}
				//时间
				if fields.FieldType == "time.Time" && len(colCell) > 0 {
					if t, err := time.ParseInLocation(DATE_TIME_PATTERN, colCell, time.Local); err == nil {
						maps[field] = t
					}
				} else {
					//其他类型
					switch fields.FieldType {
					case "bool":
						lower := strings.ToLower(colCell)
						if lower == "true" {
							maps[field] = true
						} else {
							maps[field] = false
						}
					case "int":
						int, err := strconv.Atoi(colCell)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
					case "int8":
						int, err := strconv.ParseInt(colCell, 10, 8)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
					case "int16":
						int, err := strconv.ParseInt(colCell, 10, 16)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
					case "int32":
						int, err := strconv.ParseInt(colCell, 10, 32)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
					case "int64":
						int, err := strconv.ParseInt(colCell, 10, 64)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
						//fmt.Println("int64=", int)
					case "uint":
						int, err := strconv.Atoi(colCell)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = uint(int)
						}
					case "uint8":
						int, err := strconv.ParseUint(colCell, 10, 8)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
					case "uint16":
						int, err := strconv.ParseUint(colCell, 10, 16)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
					case "uint32":
						int, err := strconv.ParseUint(colCell, 10, 32)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
					case "uint64":
						int, err := strconv.ParseUint(colCell, 10, 64)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
					case "float32":
						int, err := strconv.ParseFloat(colCell, 32)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
					case "float64":
						int, err := strconv.ParseFloat(colCell, 64)
						if err != nil {
							maps[field] = 0
						} else {
							maps[field] = int
						}
					case "string":
						maps[field] = colCell
					}
				}
			}
		}
		//json1, err := convUtil.ObjToJson(maps)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//fmt.Println("MAP=>JSON")
		//fmt.Println(json1)
		//err = json.Unmarshal([]byte(json1), &tmp)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//fmt.Println("MAP=>JSON,JSON->STRUCT")
		//fmt.Println(tmp)
		//map 转为 struct
		//if err := mapstructure.Decode(maps, &ptr); err != nil {
		//	return nil, err
		//}
		//data = append(data, ptr)
		err := callback(maps)
		if err != nil {
			return err
		}
	}
	return nil
}
