package untils

//通用输出格式

func GenTip(code int,msg string,data interface{}) map[string]interface{}{
	mapdata := map[string]interface{}{}
	mapdata["code"] = code
	mapdata["msg"] = msg
	mapdata["data"] = data
	return mapdata
}