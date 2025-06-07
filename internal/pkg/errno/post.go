// Copyright 2025 码龙 <maloong2022@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/maloong2022/miniblog. The professional
// version of this repository is https://github.com/maloong2022/onex.

package errno

import (
	"github.com/maloong2022/miniblog/pkg/errorsx"
	"net/http"
)

// ErrPostNotFound 表示未找到指定的博客.
var ErrPostNotFound = &errorsx.ErrorX{Code: http.StatusNotFound, Reason: "NotFound.PostNotFound", Message: "Post not found."}
