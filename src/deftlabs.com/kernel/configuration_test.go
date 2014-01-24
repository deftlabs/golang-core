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

package deftlabskernel

import "testing"

func TestLoadConfiguration(t *testing.T) {

	configuration, err := NewConfiguration("../../../test/configuration.json")

	if err != nil {
		t.Errorf("NewConfiguration is broken: %v", err)
		return
	}

	if configuration.Int("server.http.port", 0) != 9999 {
		t.Errorf("Configuration server.http.port is broken - expected 9999 - received: %d", configuration.Int("server.http.port", 0))
		return
	}

	/*
	if appConfiguration.SocketTimeout != 40 {

	}

	if appConfiguration.LdapServerPort != 3890 {
		t.Errorf("LdapServerPort expected: %d - received: %d", 3890, appConfiguration.LdapServerPort)
	}

	if appConfiguration.LdapServerBindAddr != "127.0.0.1" {
		t.Errorf("LdapServerBindAddr expected: %s - received: %s", "127.0.0.1", appConfiguration.LdapServerBindAddr)
	}
	*/
}

