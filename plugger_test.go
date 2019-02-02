// Copyright 2019 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugger

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("plugin registering", func() {

	var oldPluginGroups map[string]*PluginGroup

	BeforeEach(func() {
		// First save the old state of the registered plugin groups, and then
		// reset, so each test here runs on an empty plugin group map. This is
		// needed because we will otherwise trash the static plugin example,
		// making it fail depending on the sequence of tests and examples.
		oldPluginGroups = pluginGroups
		pluginGroups = map[string]*PluginGroup{}
	})

	AfterEach(func() {
		pluginGroups = oldPluginGroups
	})

	It("panics when unable to fetch runtime caller data", func() {
		Expect(func() {
			registerPlugin(
				&PluginSpec{},
				func(int) (uintptr, string, int, bool) {
					return uintptr(0), "", 0, false
				})
		}).To(Panic())
	})

	It("handles arcane caller data in registration", func() {
		Expect(func() {
			registerPlugin(&PluginSpec{},
				func(int) (uintptr, string, int, bool) {
					return uintptr(0), "plug.go", 0, true
				})
		}).To(Panic())

		Expect(func() {
			registerPlugin(&PluginSpec{},
				func(int) (uintptr, string, int, bool) {
					return uintptr(0), "/plug.go", 0, true
				})
		}).To(Panic())

		Expect(func() {
			registerPlugin(&PluginSpec{},
				func(int) (uintptr, string, int, bool) {
					return uintptr(0), "foo/plug.go", 0, true
				})
		}).To(Panic())
	})

	It("handles correct caller data in registration", func() {
		Expect(func() {
			registerPlugin(&PluginSpec{},
				func(int) (uintptr, string, int, bool) {
					return uintptr(0), "plugins/foo/plug.go", 0, true
				})
		}).ToNot(Panic())
		p := New("plugins")
		Expect(p.plugins).To(HaveLen(1))
		Expect(p.plugins[0].Name).To(Equal("foo"))
	})

})