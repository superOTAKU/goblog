package conf

import "strconv"

var conf = make(map[string]string)

func GetStr(key string) string {
	return conf[key]
}

func GetInt(key string) int {
	i, err := strconv.Atoi(conf[key])
	if err != nil {
		panic(err)
	}
	return i
}

func init() {
	conf["DbUrl"] = "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	conf["MaxIdle"] = "10"
	conf["MaxOpen"] = "100"
}
