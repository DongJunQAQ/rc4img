package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"rc4img/cmd"
)

func processCrypt(args []string, operation string) { // 加/解密共用的处理函数
	inputPath := args[0]
	outputPath := cmd.OutputFlag
	key := cmd.KeyFlag
	// 读取文件
	data, err := readFile(inputPath)
	if err != nil {
		fmt.Printf("%s失败：%v\n", operation, err)
		return
	}
	// 执行RC4加解密（本身算法对称）
	processedData := rc4Crypt(data, []byte(key))
	// 写入输出文件
	if err := writeFile(outputPath, processedData); err != nil {
		fmt.Printf("%s失败：%v\n", operation, err)
		return
	}
	fmt.Printf("%s成功！\n输入：%s\n输出：%s\n", operation, inputPath, outputPath)
}

func runEncrypt(_ *cobra.Command, args []string) { // 加密命令实现
	processCrypt(args, "加密")
}

func runDecrypt(_ *cobra.Command, args []string) { // 解密命令实现
	processCrypt(args, "解密")
}
