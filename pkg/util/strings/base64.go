// Copyright 2025 码龙 <maloong2022@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/maloong2022/miniblog. The professional
// version of this repository is https://github.com/maloong2022/onex.

package strings

import (
	"bytes"
	"encoding/base64"
	"io"
)

func DecodeBase64(i string) ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(i)))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
