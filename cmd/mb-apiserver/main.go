// Copyright 2025 码龙 <maloong2022@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/maloong2022/miniblog. The professional
// version of this repository is https://github.com/maloong2022/onex.

package main

import (
	"os"

	"github.com/maloong2022/miniblog/cmd/mb-apiserver/app"
	// 导入 automaxprocs 包，可以在程序启动时自动设置 GOMAXPROCS 配置，
	// 使其与 Linux 容器的 CPU 配额相匹配。
	// 这避免了在容器中运行时，因默认 GOMAXPROCS 值不合适导致的性能问题，
	// 确保 Go 程序能够充分利用可用的 CPU 资源，避免 CPU 浪费。
	_ "go.uber.org/automaxprocs"
)

func main() {
	command := app.NewMiniBlogCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
