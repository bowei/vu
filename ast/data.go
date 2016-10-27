/*
Copyright 2016 Bowei Du

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package ast

type Store interface {
	Put(key string, value string)
	Get(key string) (string, bool)
	List() []string
}

// TheStore is the global store.
var TheStore = new()

type mapStore struct {
	d map[string]string
}

func new() Store {
	return &mapStore{
		d: make(map[string]string),
	}
}

func (ms *mapStore) Put(key string, value string) {
	ms.d[key] = value
}

func (ms *mapStore) Get(key string) (string, bool) {
	key, ok := ms.d[key]
	return key, ok
}

func (ms *mapStore) List() []string {
	ret := make([]string, 0, len(ms.d))
	for key := range ms.d {
		ret = append(ret, key)
	}
	return ret
}
