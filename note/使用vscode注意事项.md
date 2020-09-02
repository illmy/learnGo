# 使用vscode注意事项

1. 在开始使用安装gocode等依赖时可能会一直失败，建议重新打开vscode再安装，或者手动一个个安装
2. 运行go test时，t.Log和fmt.Println都不会有输出，在.vscode目录下创建settings.json文件，新增配置项"go.testFlags": ["-v"]即可以在控制台输出中打印信息