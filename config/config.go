package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var store map[string]string // config store

func init() {
	store = map[string]string{}
}

// parse parses a given env file and transforms the to a map
func Parse(filePath string) (err error) {
	env, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer env.Close()

	scanner := bufio.NewScanner(env)
	for scanner.Scan() {
		results := strings.Split(scanner.Text(), "=")

		// incorrect format, skip
		if len(results) <= 1 {
			continue
		}

		key := results[0]
		value := results[1]

		// if "=" in value, concat all strings after first "="
		if len(results) > 2 {
			value = strings.Join(results[1:len(results)], "=")
		}

		value = strings.Replace(value, "\"", "", -1)

		store[key] = value
	}

	if err = scanner.Err(); err != nil {
		return
	}
	return
}

func GetEnv(key string) string {
	value := store[key]
	return value
}

func GetBool(key string) bool {
	value := store[key]
	result, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return result
}
