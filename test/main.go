package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/liangzibo/go-excel/lzbExcel"
	"github.com/mitchellh/mapstructure"
	"os"
)

func main() {
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
	fmt.Println("arr")
	fmt.Println(arr)

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
}
