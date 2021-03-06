// Copyright 2017 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Code generated by "stringer -type=callType"; DO NOT EDIT

package test

import "fmt"

const _callType_name = "callStatreturnStatcallGetreturnGetcallPutreturnPutcallWriteToreturnWriteTocallReadFromreturnReadFrom"

var _callType_index = [...]uint8{0, 8, 18, 25, 34, 41, 50, 61, 74, 86, 100}

func (i callType) String() string {
	if i < 0 || i >= callType(len(_callType_index)-1) {
		return fmt.Sprintf("callType(%d)", i)
	}
	return _callType_name[_callType_index[i]:_callType_index[i+1]]
}
