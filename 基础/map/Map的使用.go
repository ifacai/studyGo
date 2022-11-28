package main

import "fmt"

func main() {
	//直接创建初始化一个mao
	var mapInit = map[string]string{"xiaoli": "湖南", "xiaoliu": "天津"}
	fmt.Println(mapInit)
	//声明一个map类型变量,
	//map的key的类型是string，value的类型是string
	var mapTemp map[string]string
	//使用make函数初始化这个变量,并指定大小(也可以不指定)
	mapTemp = make(map[string]string, 10)
	//存储key ，value
	mapTemp["xiaoming"] = "北京"
	mapTemp["xiaowang"] = "河北"
	//根据key获取value,
	//如果key存在，则ok是true，否则是flase
	//v1用来接收key对应的value,当ok是false时，v1是nil
	v1, ok := mapTemp["xiaoming"]
	fmt.Println(ok, v1)
	//当key=xiaowang存在时打印value
	if v2, ok := mapTemp["xiaowang"]; ok {
		fmt.Println(v2)
	}
	//遍历map,打印key和value
	for k, v := range mapTemp {
		fmt.Println(k, v)
	}
	//删除map中的key
	delete(mapTemp, "xiaoming")
	//获取map的大小
	l := len(mapTemp)
	fmt.Println(l)
}
