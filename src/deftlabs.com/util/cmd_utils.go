/**
 * (C) Copyright 2013 Deft Labs
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
	"io"
	"time"
	"os/exec"
)

// TimedCmdExec allows you to execute commands for a "max" period of time. After
// the specified amount of time, the command will be killed. If you
// want access to the stdout/err then you need to pass in a Writer to receive the data.
// This does not support sending data to the command executing. This method will panic if
// maxTimeMs is <= 0. If the cmdName length is zero, it will also panic.
func CmdExecWithMaxTime(cmdName string, maxTimeMs int64, stdOutPipe io.Writer, stdErrPipe io.Writer, args ...string) error {

	if maxTimeMs <= 0 {
		panic("You must set maxTimeMs to greater than zero")
	}

	if len(cmdName) == 0 {
		panic("You must set cmdName to a non-mepty string")
	}

	cmd := exec.Command(cmdName, args...)
	done := make(chan error)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	if stdErrPipe != nil {
		go io.Copy(stdErrPipe, stderr)
	}

	if stdOutPipe != nil {
		go io.Copy(stdOutPipe, stdout)
	}

	go func() { done <- cmd.Wait() }()

	select {
		case <-time.After(time.Duration(maxTimeMs) * time.Millisecond): {
			err = cmd.Process.Kill()

			// The goroutine will be able to exit because wait returns after the kill
			<-done
			return err
		}

		case err = <-done:
			return err
	}

	return err
}

