// Copyright 2016 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package scanner_test

import (
	"testing"

	"github.com/savaki/jq/scanner"
	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkNumber(t *testing.B) {
	data := []byte(`12.34e+9`)

	for i := 0; i < t.N; i++ {
		end, err := scanner.Number(data, 0)
		if err != nil {
			t.FailNow()
			return
		}

		if end == 0 {
			t.FailNow()
			return
		}
	}
}

func TestNumber(t *testing.T) {
	Convey("Verify String", t, func() {
		testCases := map[string]struct {
			In     string
			Out    string
			HasErr bool
		}{
			"simple": {
				In:  `1234`,
				Out: `1234`,
			},
			"decimal": {
				In:  `1.234`,
				Out: `1.234`,
			},
			"spaced": {
				In:  `  1.234   `,
				Out: `  1.234`,
			},
			"kitchen-sink": {
				In:  `  +-123.25eE10 `,
				Out: `  +-123.25eE10`,
			},
		}

		for label, tc := range testCases {
			Convey(label, func() {
				end, err := scanner.Number([]byte(tc.In), 0)
				if tc.HasErr {
					So(err, ShouldNotBeNil)
				} else {
					data := tc.In[0:end]
					So(string(data), ShouldEqual, tc.Out)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}