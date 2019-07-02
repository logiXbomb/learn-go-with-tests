package main

const enUs = "Hello, "

func Hello(str string) string {
	if str == "" {
		return enUs + "world"
	}
	return enUs + str
}
