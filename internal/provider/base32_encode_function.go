// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = Base32EncodeFunction{}
)

func NewBase32EncodeFunction() function.Function {
	return Base32EncodeFunction{}
}

type Base32EncodeFunction struct{}

func (r Base32EncodeFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "base32_encode"
}

func (r Base32EncodeFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "base32 encode",
		MarkdownDescription: "base32 encode",
		Parameters: []function.Parameter{
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

func (r Base32EncodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var message string
	var encoding string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &message, &encoding))
	if resp.Error != nil {
		return
	}

	// 解码消息
	msg, err := b64Enc.DecodeString(message)
	if err != nil {
		resp.Error = function.NewFuncError("message is invalid base64 string")
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, b32Enc.EncodeToString(msg)))
}
