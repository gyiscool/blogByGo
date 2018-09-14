package models

import(
	"fmt"
	"reflect"
	//"strconv"
)


//提取某个struct 的数组中某个字段值 把字符串和整型都转换成 string
func ArrayColums(list interface{},fieldStr string) []string{
//func ArrayColums(list interface{},fieldStr string) ([]string){

	v := reflect.ValueOf(list)

	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}

	l := v.Len()

	inputs := make([]reflect.Value, l)

	ret := make([]interface{}, l)
	res := make([]string,l)

	typeStr := "int"

	//判断字段是什么类型的
	if l >= 1 {

		fmt.Println(fieldStr)
		tempType := (fmt.Sprintf("%s", ((v.Index(0).FieldByName(fieldStr)).Type())))

		switch  tempType {
		case "string":
				typeStr = "string"
		case "int":
				typeStr = "int"
		default :
				typeStr = "string"
		}

		fmt.Printf("我的类型%+v\n\n\n",typeStr)

		fmt.Println(typeStr)
	}


	//循环提取
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
		inputs[i] = (v.Index(i).FieldByName("Iid"))

		if typeStr == "int" {
			res[i] = fmt.Sprintf("%d",(v.Index(i).FieldByName(fieldStr).Int()))
		}else if typeStr == "string" {
			res[i] = fmt.Sprintf("%s",(v.Index(i).FieldByName(fieldStr).String()))
		}
	}

	return res

}

