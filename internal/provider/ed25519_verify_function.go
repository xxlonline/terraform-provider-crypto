// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"crypto/ed25519"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = Ed25519VerifyFunction{}
)

func NewEd25519VerifyFunction() function.Function {
	return Ed25519VerifyFunction{}
}

type Ed25519VerifyFunction struct{}

func (r Ed25519VerifyFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ed25519_verify"
}

func (r Ed25519VerifyFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "ed25519 verify",
		MarkdownDescription: "ed25519 verify",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "publicKey",
				MarkdownDescription: "publicKey",
				AllowNullValue:      false,
				AllowUnknownValues:  false,
			},
			function.StringParameter{
				Name:                "message",
				MarkdownDescription: "message",
				AllowNullValue:      false,
				AllowUnknownValues:  false,
			},
			function.StringParameter{
				Name:                "sig",
				MarkdownDescription: "sig",
				AllowNullValue:      false,
				AllowUnknownValues:  false,
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r Ed25519VerifyFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var publicKey string
	var message string
	var sig string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &publicKey, &message, &sig))
	if resp.Error != nil {
		return
	}

	// 解码公钥
	key, err := b64Enc.DecodeString(publicKey)
	if err != nil {
		resp.Error = function.NewFuncError("invalid public key")
		return
	}

	// 解码消息
	msg, err := b64Enc.DecodeString(message)
	if err != nil {
		resp.Error = function.NewFuncError("invalid seed")
		return
	}

	// 解码签名
	bsig, err := b64Enc.DecodeString(sig)
	if err != nil {
		resp.Error = function.NewFuncError("invalid seed")
		return
	}
	ret := ed25519.Verify(key, msg, bsig)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, ret))
}
