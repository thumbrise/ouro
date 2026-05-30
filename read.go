// Copyright 2026 thumbrise
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/thumbrise/ouro/contract"
	"github.com/thumbrise/ouro/reference/pkg/reflection"
	stringsutil "github.com/thumbrise/ouro/reference/pkg/strings"
)

type Reader struct {
	validator       contract.Validator
	loader          contract.Loader
	hookSuccessRead contract.HookSuccessRead
}

func NewReader(hookSuccessLoad contract.HookSuccessRead, loader contract.Loader, validator contract.Validator) *Reader {
	return &Reader{hookSuccessRead: hookSuccessLoad, loader: loader, validator: validator}
}

func (c *Reader) Read(ctx context.Context, out interface{}, key string) error {
	loadContext := contract.LoadContext{Key: key, Data: out}

	err := c.loader.Load(ctx, loadContext)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrLoad, err)
	}

	err = c.validator.Validate(ctx, contract.ValidateContext{Key: key, Data: out})
	if err != nil {
		return fmt.Errorf("%w: %w", ErrValidate, err)
	}

	c.hookSuccessRead(ctx, loadContext)

	return nil
}

var (
	ErrLoad     = errors.New("failed to load configuration")
	ErrValidate = errors.New("validate")
)

// Read uses go generics for inspect needed type and config key and return new instance.
//
// Using reflection to retrieve name of type.
// First letter of type lowercased to satisfy go naming convention of package exported types.
// Expects config naming in lowerCamelCase.
// Expects using Struct with name equals key in root of config.
//
// Always returning pointer of type. Even if you use own type for root config string/int/boolean values. So you need dereference result value in case of primitive types.
// See examples.
//
// Working example: `
//
//	myParams:
//		a: 5
//		b: "s"
//
// ...
//
//	type MyParams struct {
//			A int
//			B string
//		}
//
// `
//
// Working example: `
//
//	rootA: 5
//	rootB: "s"
//
// ...
//
//	type RootA int
//	type RootB string
//
// `
//
// Was inspired by gorm.G.
func Read[T any](ctx context.Context, reader *Reader) (*T, error) {
	v := *new(T)

	key := stringsutil.LowerFirst(reflection.TypeName(v))

	err := reader.Read(ctx, &v, key)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
