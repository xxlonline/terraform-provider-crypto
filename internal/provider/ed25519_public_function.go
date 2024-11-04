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
	_ function.Function = Ed25519PublicFunction{}
)

func NewEd25519PublicFunction() function.Function {
	return Ed25519PublicFunction{}
}

type Ed25519PublicFunction struct{}

func (r Ed25519PublicFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ed25519_public"
}

func (r Ed25519PublicFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "ed25519 public",
		MarkdownDescription: "ed25519 public",
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

func (r Ed25519PublicFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
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
	pub, _, err := ed25519.GenerateKey(bytes.NewReader(bseed))
	if err != nil {
		resp.Error = function.NewFuncError("invalid seed")
		return
	}

	// 返回签名
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, b64Enc.EncodeToString(pub)))
}
