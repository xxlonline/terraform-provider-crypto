// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"bytes"
	"context"
	"crypto/ed25519"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = Ed25519PrivateFunction{}
)

func NewEd25519PrivateFunction() function.Function {
	return Ed25519PrivateFunction{}
}

type Ed25519PrivateFunction struct{}

func (r Ed25519PrivateFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ed25519_private"
}

func (r Ed25519PrivateFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "ed25519 private",
		MarkdownDescription: "ed25519 private",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "seed",
				MarkdownDescription: "seed",
				AllowNullValue:      false,
				AllowUnknownValues:  false,
			},
		},
		Return: function.StringReturn{},
	}
}

func (r Ed25519PrivateFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var seed string

	// 获取参数
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &seed))
	if resp.Error != nil {
		return
	}

	// 解码种子
	bseed, err := b64Enc.DecodeString(seed)
	if err != nil {
		resp.Error = function.NewFuncError("invalid seed")
		return
	}

	// 获取公钥
	_, priv, err := ed25519.GenerateKey(bytes.NewReader(bseed))
	if err != nil {
		resp.Error = function.NewFuncError("invalid seed")
		return
	}

	// 返回签名
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, b64Enc.EncodeToString(priv)))
}
