/**
 * (C) Copyright 2013, Deft Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package deftlabsutil

import (
	"fmt"
	"crypto/md5"
	"deftlabs.com/log"
)

func Md5HexFromBytes(v []byte) (string, error) {

	if len(v) == 0 {
		return "", slogger.NewStackError("Value cannot be nil/empty")
	}

	h := md5.New()

	written, err := h.Write(v)

	if err != nil {
		return "", err
	}

	if written != len(v) {
		return "", slogger.NewStackError("Written does not equal length - written: %d - len: %d", written, len(v))
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func Md5Hex(v string) (string, error) {
	return Md5HexFromBytes([]byte(v))
}

