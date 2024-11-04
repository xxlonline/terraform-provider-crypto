// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = Base32DecodeFunction{}
)

func NewBase32DecodeFunction() function.Function {
	return Base32DecodeFunction{}
}

type Base32DecodeFunction struct{}

func (r Base32DecodeFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "base32_decode"
}

func (r Base32DecodeFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "base32 decode",
		MarkdownDescription: "base32 decode",
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

func (r Base32DecodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var message string
	var encoding string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &message, &encoding))
	if resp.Error != nil {
		return
	}

	// 解码消息
	msg, err := b32Enc.DecodeString(message)
	if err != nil {
		resp.Error = function.NewFuncError("message is invalid base32 string")
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, b64Enc.EncodeToString(msg)))

}
