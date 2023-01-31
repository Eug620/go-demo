package JSONHandle

// 反序列化的方法
import (
	"encoding/json"
	"fmt"
	"os"
)

// 解码json文件为map
func jsonFileToMap(fileUrl string) map[string]interface{} {
	//打开json文件
	srcFile, _ := os.Open(fileUrl)
	defer srcFile.Close()
	//创建接收数据的map类型数据
	datamap := make(map[string]interface{})
	//创建解码器
	decoder := json.NewDecoder(srcFile)
	//解码
	err := decoder.Decode(&datamap)
	if err != nil {
		fmt.Println("解码失败,err=", err)
		return datamap
	}
	fmt.Println("解码成功", datamap)
	return datamap
}

// 将json数据转为map
func jsonToMap(jsonStr string) map[string]interface{} {

	//将json数据转为字节数据
	jsonbytes := []byte(jsonStr)
	dataMap := make(map[string]interface{})

	err := json.Unmarshal(jsonbytes, &dataMap)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
	}
	fmt.Println(dataMap)
	return dataMap
}

// 将json数据转换为map切片
func jsonToMapSlice(jsonStr string) []map[string]interface{} {
	jsonBytes := []byte(jsonStr)
	dataSlice := make([]map[string]interface{}, 0)
	err := json.Unmarshal(jsonBytes, &dataSlice)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
	}
	fmt.Println(dataSlice)
	return dataSlice
}
