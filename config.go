package main

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var Config config

type config struct {
	InputPath          string `default:"input/friends.json"`
	OutputJsonPath     string `default:"output/friends.json"`
	OutputMarkdownPath string `default:"output/friends.md"`
	TemplatePath       string `default:"template/template.md"`
	NumOutputPost      int    `default:"20"`
	ContentLength      int    `default:"200"`
}

// Get configs from ENV, you set them in GitHub Actions Workflow
func init() {
	Config = loadConfigFromEnv()
}

func loadConfigFromEnv() config {
	cfg := config{}
	t := reflect.TypeOf(cfg)
	v := reflect.ValueOf(&cfg)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		defaultValue := field.Tag.Get("default")
		envValue := getEnv(ConvertToSnakeCase(field.Name), defaultValue)
		setFieldValue(v.Elem().Field(i), envValue)
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Warning: Environment variable %s not set. Using default value: %s", key, defaultValue)
		return defaultValue
	}
	return value
}

func setFieldValue(field reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	// Add support for other data types here if needed
	// For example:
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err == nil {
			field.SetInt(int64(intValue))
		} else {
			log.Printf("Warning: Environment variable %s is not a int. Using default value: %s", field, value)
		}
		// case reflect.Bool:
		//     boolValue, err := strconv.ParseBool(value)
		//     if err == nil {
		//         field.SetBool(boolValue)
		//     }
	}
}

// ConvertToSnakeCase converts a string to snake case
// InputPath -> INPUT_PATH
func ConvertToSnakeCase(input string) string {
	var result strings.Builder
	for i, c := range input {
		if unicode.IsUpper(c) {
			if i > 0 {
				result.WriteByte('_')
			}
			result.WriteRune(unicode.ToUpper(c))
		} else {
			result.WriteRune(unicode.ToUpper(c))
		}
	}
	return result.String()
}
