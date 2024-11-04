// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"crypto/ed25519"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = Ed25519SignFunction{}
)

func NewEd25519SignFunction() function.Function {
	return Ed25519SignFunction{}
}

type Ed25519SignFunction struct{}

func (r Ed25519SignFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ed25519_sign"
}

func (r Ed25519SignFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "ed25519 sign",
		MarkdownDescription: "ed25519 sign",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "privateKey",
				MarkdownDescription: "privateKey",
				AllowNullValue:      false,
				AllowUnknownValues:  false,
			},
			function.StringParameter{
				Name:                "message",
				MarkdownDescription: "message",
				AllowNullValue:      false,
				AllowUnknownValues:  false,
			},
		},
		Return: function.StringReturn{},
	}
}

func (r Ed25519SignFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var privateKey string
	var message string

	// 获取参数
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &privateKey, &message))
	if resp.Error != nil {
		return
	}

	// 解码公钥
	key, err := b64Enc.DecodeString(privateKey)
	if err != nil {
		resp.Error = function.NewFuncError("invalid private key")
		return
	}

	// 解码消息
	msg, err := b64Enc.DecodeString(message)
	if err != nil {
		resp.Error = function.NewFuncError("invalid message")
		return
	}

	// 计算签名
	sig := ed25519.Sign(key, msg)

	// 返回签名
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, b64Enc.EncodeToString(sig)))
}
