package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/liangzibo/go-excel/lzbExcel"
	"github.com/mitchellh/mapstructure"
	"os"
)

func main() {
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
	var arr []lzbExcel.ExcelTest
	err = lzbExcel.NewExcelStructDefault().SetPointerStruct(&lzbExcel.ExcelTest{}).RowsProcess(rows, func(maps map[string]interface{}) error {
		var ptr lzbExcel.ExcelTest
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
	var arr2 []lzbExcel.ExcelTest
	//StartRow 开始行,索引从 0开始
	//IndexMax  索引最大行,如果 结构体中的 index 大于配置的,那么使用结构体中的
	err = lzbExcel.NewExcelStruct(1, 10).SetPointerStruct(&lzbExcel.ExcelTest{}).RowsProcess(rows, func(maps map[string]interface{}) error {
		var ptr lzbExcel.ExcelTest
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
}
