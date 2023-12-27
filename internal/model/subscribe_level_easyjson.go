// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson83a9dd39DecodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(in *jlexer.Lexer, out *SubscribeLevel) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint(in.Uint())
		case "level":
			out.Level = uint(in.Uint())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "currency":
			out.Currency = string(in.String())
		case "cost_integer":
			out.CostInteger = uint(in.Uint())
		case "cost_fractional":
			out.CostFractional = uint(in.Uint())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson83a9dd39EncodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(out *jwriter.Writer, in SubscribeLevel) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"level\":"
		out.RawString(prefix)
		out.Uint(uint(in.Level))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"currency\":"
		out.RawString(prefix)
		out.String(string(in.Currency))
	}
	{
		const prefix string = ",\"cost_integer\":"
		out.RawString(prefix)
		out.Uint(uint(in.CostInteger))
	}
	{
		const prefix string = ",\"cost_fractional\":"
		out.RawString(prefix)
		out.Uint(uint(in.CostFractional))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SubscribeLevel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson83a9dd39EncodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SubscribeLevel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson83a9dd39EncodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SubscribeLevel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson83a9dd39DecodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SubscribeLevel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson83a9dd39DecodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(l, v)
}