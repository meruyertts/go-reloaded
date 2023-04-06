package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1:]
	var s, r string
	if len(input) == 1 {
		s = input[0]
	} else if len(input) == 2 {
		r = input[1]
		s = input[0]
	}

	if len(s) > 0 {
		input, err := os.Open(s)
		if err != nil {
			fmt.Println(err)
		}
		data := make([]byte, 1000)
		count, err := input.Read(data)
		if err != nil {
			fmt.Println(err)
		}
		if count > 0 {

			newArr := []string{}
			arr := strings.Split(string(data[:count]), " ")
			arr = mergeBrackets(arr)
			arr = removeEmptyString(arr)

			for _, word := range arr {
				lenArr := len(newArr)
				if word == "(hex)" {
					n, _ := strconv.ParseInt(newArr[lenArr-1], 16, 32)
					newArr[lenArr-1] = strconv.Itoa(int(n))
					if checkAfterBrackets(word) {
						newArr = append(newArr, addAfterBrackets(word))
					}
				} else if word == "(bin)" {
					n, _ := strconv.ParseInt(newArr[lenArr-1], 2, 32)
					newArr[lenArr-1] = strconv.Itoa(int(n))
					if checkAfterBrackets(word) {
						newArr = append(newArr, addAfterBrackets(word))
					}
				} else if strings.HasPrefix(word, "(low") {

					re := regexp.MustCompile("[0-9]+")
					mynum := re.FindAllString(word, 1)
					if mynum != nil {
						num, _ := strconv.Atoi(mynum[0])
						if num > len(newArr) {
							num = len(newArr)
						}
						for i := 1; i <= num; i++ {
							newArr[lenArr-i] = strings.ToLower(newArr[lenArr-i])
						}
					} else {
						newArr[lenArr-1] = strings.ToLower(newArr[lenArr-1])
					}

					if checkAfterBrackets(word) {
						newArr = append(newArr, addAfterBrackets(word))
					}
				} else if strings.HasPrefix(word, "(up") {
					re := regexp.MustCompile("[0-9]+")
					mynum := re.FindAllString(word, 1)
					if mynum != nil {
						num, _ := strconv.Atoi(mynum[0])
						if num >= 1 {
							if num > len(newArr) {
								num = len(newArr)
							}
							for i := 1; i <= num; i++ {
								newArr[lenArr-i] = strings.ToUpper(newArr[lenArr-i])
							}
						}
					} else {
						newArr[lenArr-1] = strings.ToUpper(newArr[lenArr-1])
					}

					if checkAfterBrackets(word) {
						newArr = append(newArr, addAfterBrackets(word))
					}
				} else if strings.HasPrefix(word, "(cap") {
					re := regexp.MustCompile("[0-9]+")
					mynum := re.FindAllString(word, 1)
					if mynum != nil {
						num, _ := strconv.Atoi(mynum[0])
						if num > len(newArr) {
							num = len(newArr)
						}
						for i := 1; i <= num; i++ {
							newArr[lenArr-i] = strings.ToLower(newArr[lenArr-i])
							newArr[lenArr-i] = strings.Title(newArr[lenArr-i])
						}
					} else {
						newArr[lenArr-1] = strings.ToLower(newArr[lenArr-1])
						newArr[lenArr-1] = strings.Title(newArr[lenArr-1])
					}
					if checkAfterBrackets(word) {
						newArr = append(newArr, addAfterBrackets(word))
					}

				} else {
					newArr = append(newArr, word)
				}

			}
			result := strings.Join(newArr, " ")
			result = containsPunctuation(result)
			result = singleQuot1(result)
			result = singleQuot2(result)
			result = singleQuot3(result)
			result = checkArticle(result)

			if len(r) > 0 {
				file, err := os.Create(string(r))
				if err != nil {
					fmt.Println(err)
				} else {
					result = strings.TrimSpace(result)
					result += "\n"
					file.WriteString(result)
				}
				file.Close()
			}

		}

	}
}

func removeEmptyString(s []string) []string {
	var r []string
	for _, word := range s {
		if word != "" {
			r = append(r, word)
		}
	}
	return r
}

func mergeBrackets(s []string) []string {
	for index, word := range s {
		if strings.HasPrefix(word, "(") {
			if index != len(s)-1 {
				if strings.Contains(s[index+1], ")") {
					s[index] = s[index] + s[index+1]
					s[index+1] = ""
				}
			}
		}
	}
	return s
}

func containsPunctuation(s string) string {
	result := ""
	re := regexp.MustCompile("[[:punct:]]+")
	punct := re.FindAllString(s, -1)
	text := re.Split(s, -1)
	myPunct := []string{".", ",", "!", "?", ":", ";"}
	for i := 0; i < len(punct); i++ {
		result += strings.TrimSpace(text[i])
		result += strings.TrimSpace(punct[i])
		if len(punct[i]) == 1 {
			for _, symbol := range myPunct {
				if symbol == punct[i] {
					result += " "
				}
			}
		} else if len(punct[i]) > 1 {
			result += " "
		}

	}
	result += strings.TrimSpace(text[len(text)-1])
	return result
}

func singleQuot1(s string) string {
	result := ""
	re := regexp.MustCompile("'(.*?)'")
	punct := re.FindAllString(s, -1)
	quot_word := ""
	replaced_word := []string{}

	for _, word := range punct {
		quot_word = strings.ReplaceAll(word, "' ", "'")
		quot_word = strings.ReplaceAll(quot_word, " '", "'")
		replaced_word = append(replaced_word, quot_word)
	}

	text := re.Split(s, -1)
	for i := 0; i < len(replaced_word); i++ {
		result += strings.TrimSpace(text[i])
		result += " "
		result += strings.TrimSpace(replaced_word[i])

	}
	result += text[len(text)-1]
	return result
}

func singleQuot2(s string) string {
	result := ""
	re := regexp.MustCompile("‘(.*?)’")
	punct := re.FindAllString(s, -1)
	quot_word := ""
	replaced_word := []string{}

	for _, word := range punct {
		quot_word = strings.ReplaceAll(word, "‘ ", "‘")
		quot_word = strings.ReplaceAll(quot_word, " ’", "’")
		replaced_word = append(replaced_word, quot_word)
	}

	text := re.Split(s, -1)
	for i := 0; i < len(replaced_word); i++ {
		result += strings.TrimSpace(text[i])
		result += " "
		result += strings.TrimSpace(replaced_word[i])

	}
	result += text[len(text)-1]
	return result
}

func singleQuot3(s string) string {
	result := ""
	re := regexp.MustCompile("‛(.*?)‛")
	punct := re.FindAllString(s, -1)
	quot_word := ""
	replaced_word := []string{}

	for _, word := range punct {
		quot_word = strings.ReplaceAll(word, "‛ ", "‛")
		quot_word = strings.ReplaceAll(quot_word, " ‛", "‛")
		replaced_word = append(replaced_word, quot_word)
	}

	text := re.Split(s, -1)
	for i := 0; i < len(replaced_word); i++ {
		result += strings.TrimSpace(text[i])
		result += " "
		result += strings.TrimSpace(replaced_word[i])

	}
	result += text[len(text)-1]
	return result
}

func checkArticle(s string) string {
	myResult := ""
	myArr := strings.Split(s, " ")

	for index, word := range myArr {
		articleChange := false
		if word == "a" && index != len(myArr)-1 {
			vowels := []string{"a", "e", "i", "o", "u", "h"}
			for _, letter := range vowels {
				if strings.HasPrefix(myArr[index+1], letter) {
					word = "an"
					articleChange = true
					myResult += word
					myResult += " "
				}
			}
			if !articleChange {
				myResult += word
				myResult += " "
			}

		} else if word == "A" && index != len(myArr)-1 {
			vowels := []string{"a", "e", "i", "o", "u", "h"}
			for _, letter := range vowels {
				if strings.HasPrefix(myArr[index+1], letter) {
					word = "An"
					articleChange = true
					myResult += word
					myResult += " "
				}
			}
			if !articleChange {
				myResult += word
				myResult += " "
			}

		} else {
			myResult += word
			if index != len(myArr) {
				myResult += " "
			}
		}
	}
	return myResult
}

func checkAfterBrackets(s string) bool {
	result := strings.Split(s, ")")
	return result[1] != ""
}

func addAfterBrackets(s string) string {
	result := strings.Split(s, ")")
	return result[1]
}
