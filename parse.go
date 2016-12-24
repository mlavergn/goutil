// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goutil

import (
	. "golog"
	"strings"
)

const PARSE_MAXLEN int = 100000

func FixedLengthSplit(positions [][]int, input string) []string {
	var result = []string{}

	LogDebug(input)
	maxend := len(input)

	for _, pos := range positions {
		start := pos[0]
		end := pos[1]

		if start > maxend || (end > maxend && end != PARSE_MAXLEN) {
			// line length is too short to proceed
			break
		}

		value := ""
		if end == PARSE_MAXLEN {
			value = input[start:]
		} else {
			value = input[start:end]
		}

		value = strings.TrimSpace(value)
		result = append(result, value)
	}
	LogDebug(result)

	return result
}
