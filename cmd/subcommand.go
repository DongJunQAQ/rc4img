package cmd

import (
	"github.com/spf13/cobra"
	"rc4ctl/internal"
)

func init() {
	// 将加密子命令添加到根命令
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
}

var encryptCmd = &cobra.Command{ //定义加密子命令
	Use:   "encrypt [input-file-path]",
	Short: "使用RC4算法加密文件",
	Long:  `将指定图片文件通过RC4算法加密，输出为不可直接查看的文件（需用decrypt子命令解密）`,
	Args:  cobra.ExactArgs(1), // 必须传入1个位置参数：输入要加密的文件路径
	Run: func(_ *cobra.Command, args []string) { // 加密逻辑执行函数
		internal.ProcessCrypt(args, "加密", outputFlag, keyFlag)
	},
}

var decryptCmd = &cobra.Command{ //定义解密子命令
	Use:   "decrypt [input-file-path]",
	Short: "使用RC4算法解密文件",
	Long:  `将通过encrypt子命令加密的文件解密，恢复为原始可查看的图片文件（需使用加密时的相同密钥）`,
	Args:  cobra.ExactArgs(1), // 必须传入1个位置参数：输入要解密的文件路径
	Run: func(_ *cobra.Command, args []string) { // 解密逻辑执行函数
		internal.ProcessCrypt(args, "解密", outputFlag, keyFlag)
	},
}
