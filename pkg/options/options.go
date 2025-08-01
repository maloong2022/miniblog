// Copyright 2025 码龙 <maloong2022@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/maloong2022/miniblog. The professional
// version of this repository is https://github.com/maloong2022/onex.

package options

import "github.com/spf13/pflag"

// IOptions defines methods to implement a generic options.
type IOptions interface {
	// Validate validates all the required options.
	// It can also used to complete options if needed.
	Validate() []error

	// AddFlags adds flags related to given flagset.
	AddFlags(fs *pflag.FlagSet, prefixes ...string)
}
