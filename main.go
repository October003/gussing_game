package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	//用程序启动的时间戳初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The scretNumber is ", secretNumber)
	fmt.Println("Pleasr input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		//读取一行输入
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input,Please try again", err)
			continue
		}
		//去掉换行符
		input = strings.TrimSuffix(input, "\r\n")
		//转换成数字
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input.Please enter an integer value", err)
			continue
		}
		fmt.Println("Your guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number.Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number.Please try again")
		} else {
			fmt.Println("Correct,you Legend")
			break
		}
	}
}
