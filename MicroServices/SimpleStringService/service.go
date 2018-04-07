/*
Copyright 2018 Moshe Tubul
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
package SimpleStringService

import (
	"errors"
	"strings"
)

var ErrEmptyString = errors.New("Empty String")

// service operation interface
type StringService interface {
	Uppercase(string) (string, error)
	Lowercase(string) (string, error)
	CutSub(string, string) (string, error)
	Count(string) int
}

type stringService struct {}

func (p stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	return strings.ToUpper(s), nil
}

func (p stringService) Lowercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	return strings.ToLower(s), nil
}

func (p stringService) CutSub(src string, sub string) (string, error) {
	if src == "" {
		return "", ErrEmptyString
	}
	replacer := strings.NewReplacer(sub, "")
	return replacer.Replace(src), nil
}

func (p stringService) Count(s string) int {
	return len(s)
}
