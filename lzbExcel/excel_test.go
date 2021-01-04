package lzbExcel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/mitchellh/mapstructure"
	"os"
	"testing"
	"time"
)

type ExcelTest struct {
	Name     string    `json:"name" name:"名称" index:"0"`
	Age      int64     `json:"age" name:"年龄" index:"1"`
	Score    int64     `json:"score" name:"分数" index:"2"`
	Birthday time.Time `json:"birthday" name:"生日" index:"4"`
}

func TestExcel(t *testing.T) {
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
	var arr []ExcelTest
	err = NewExcelStructDefault().SetPointerStruct(&ExcelTest{}).RowsProcess2(rows, func(maps map[string]interface{}) error {
		var ptr ExcelTest
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
	//StartRow 开始行,索引从 0开始
	//IndexMax  索引最大行,如果 结构体中的 index 大于配置的,那么使用结构体中的
	var arr2 []ExcelTest
	err = NewExcelStruct(1, 10).SetPointerStruct(&ExcelTest{}).RowsProcess2(rows, func(maps map[string]interface{}) error {
		var ptr ExcelTest
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
}
