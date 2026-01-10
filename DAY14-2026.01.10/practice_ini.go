// package main
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Config 配置结构体，通过 ini 标签映射到配置文件
type Config struct {
	Database struct {
		Host     string `ini:"host"`
		Port     int    `ini:"port"`
		Username string `ini:"username"`
		Debug    bool   `ini:"debug"`
	} `ini:"database"`

	Server struct {
		Port int    `ini:"port"`
		Name string `ini:"name"`
	} `ini:"server"`
}

// IniData 存储解析后的 INI 数据
// 格式: map[section]map[key]value
type IniData map[string]map[string]string

// ParseIni 解析 INI 配置文件并填充到结构体
// filename: 配置文件路径
// v: 必须传结构体指针，通过反射修改其值
func ParseIni(filename string, v interface{}) error {
	// ========== 第一步：参数校验 ==========
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return errors.New("必须传入指针类型")
	}

	// Elem() 解引用指针，获取指针指向的值
	if rv.Elem().Kind() != reflect.Struct {
		return errors.New("必须指向结构体类型")
	}

	structValue := rv.Elem()

	// ========== 第二步：解析 INI 文件 ==========
	data, err := parseFile(filename)
	if err != nil {
		return fmt.Errorf("解析文件失败: %w", err)
	}

	// ========== 第三步：通过反射填充结构体 ==========
	// 遍历结构体的每一个字段
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)      // 字段的值（可修改）
		fieldType := structValue.Type().Field(i) // 字段的类型（只读，含标签）

		// 获取字段的 ini 标签，对应 section 名称
		section := fieldType.Tag.Get("ini")

		// 跳过没有标签的字段
		if section == "" {
			continue
		}

		// 检查该字段是否是结构体（嵌套结构体）
		if field.Kind() == reflect.Struct {
			// 递归填充嵌套结构体
			if err := fillStruct(data[section], field); err != nil {
				return fmt.Errorf("填充 [%s] 失败: %w", section, err)
			}
		}
	}

	return nil
}

// parseFile 读取并解析 INI 文件
func parseFile(filename string) (IniData, error) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // 确保函数返回时关闭文件

	data := make(IniData)
	var currentSection string

	// 使用带缓冲的 Scanner 按行读取
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 跳过空行和注释行
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 解析 [section]
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			// 提取 section 名称，去掉 [ 和 ]
			currentSection = line[1 : len(line)-1]
			// 初始化该 section 的 map
			data[currentSection] = make(map[string]string)
			continue
		}

		// 解析 key = value
		if idx := strings.Index(line, "="); idx != -1 {
			key := strings.TrimSpace(line[:idx])      // 等号左边
			value := strings.TrimSpace(line[idx+1:])  // 等号右边

			// 确保有当前 section
			if currentSection != "" && data[currentSection] != nil {
				data[currentSection][key] = value
			}
		}
	}

	// 检查扫描过程是否有错误
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// fillStruct 用数据填充结构体
// sectionData: 当前 section 的 key-value 数据
// structValue: 要填充的结构体值（可修改）
func fillStruct(sectionData map[string]string, structValue reflect.Value) error {
	// 遍历嵌套结构体的每个字段
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := structValue.Type().Field(i)

		// 获取字段的 ini 标签，对应配置文件中的 key
		key := fieldType.Tag.Get("ini")
		if key == "" {
			continue
		}

		// 从解析的数据中查找对应的值
		value, exists := sectionData[key]
		if !exists {
			continue // 配置文件中没有这个 key，跳过
		}

		// 根据字段类型设置值
		if err := setFieldValue(field, value); err != nil {
			return fmt.Errorf("设置字段 %s 失败: %w", key, err)
		}
	}

	return nil
}

// setFieldValue 根据字段类型设置值
func setFieldValue(field reflect.Value, value string) error {
	// 检查字段是否可设置（指针解引用后的字段才可设置）
	if !field.CanSet() {
		return errors.New("字段不可设置")
	}

	switch field.Kind() {
	case reflect.String:
		// 直接设置字符串
		field.SetString(value)

	case reflect.Int, reflect.Int32, reflect.Int64:
		// 字符串转整数
		num, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("无法转换为整数: %s", value)
		}
		field.SetInt(num)

	case reflect.Bool:
		// 字符串转布尔值
		// strconv.ParseBool 支持: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("无法转换为布尔值: %s", value)
		}
		field.SetBool(b)

	case reflect.Float32, reflect.Float64:
		// 字符串转浮点数
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("无法转换为浮点数: %s", value)
		}
		field.SetFloat(f)

	default:
		return fmt.Errorf("不支持的类型: %s", field.Kind())
	}

	return nil
}

func main() {
	var cfg Config

	// 解析配置文件
	err := ParseIni("config.ini", &cfg)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	// 打印解析结果
	fmt.Println("========== 解析结果 ==========")
	fmt.Printf("Database.Host: %s\n", cfg.Database.Host)
	fmt.Printf("Database.Port: %d\n", cfg.Database.Port)
	fmt.Printf("Database.Username: %s\n", cfg.Database.Username)
	fmt.Printf("Database.Debug: %v\n", cfg.Database.Debug)
	fmt.Println("-----------------------------")
	fmt.Printf("Server.Port: %d\n", cfg.Server.Port)
	fmt.Printf("Server.Name: %s\n", cfg.Server.Name)
	fmt.Println("=============================")

	// 打印完整结构体
	fmt.Printf("\n完整结构体: %+v\n", cfg)
}
