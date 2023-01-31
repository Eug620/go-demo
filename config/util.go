package projectConfig

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

// 创建.env文件
func createEnv() {
	envConfig := "ENV = pro"
	buf := &bytes.Buffer{}
	buf.WriteString(envConfig)
	env, err := godotenv.Parse(buf)
	if err != nil {
		log.Fatal(err)
	}
	path := getAppPath() + "/.env"
	err = godotenv.Write(env, path)
	if err != nil {
		log.Fatal(err)
	}
}

// 获取当前环境
func getEnv() string {
	// 加载.env文件配置
	err := godotenv.Load(".env")
	// 获取运行环境
	env := os.Getenv("ENV")
	if err != nil {
		fmt.Println("配置文件加载失败,现在新建一个配置文件")
		env = "pro"
	}

	fmt.Println("当前环境开发:", env)

	return env
}

// 获取当前路径
func getAppPath() string {
	var _, b, _, _ = runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../")
}
