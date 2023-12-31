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

func easyjson13673cd6DecodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(in *jlexer.Lexer, out *Tag) {
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
		case "post_id":
			out.PostId = uint(in.Uint())
		case "name":
			out.Name = string(in.String())
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
func easyjson13673cd6EncodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(out *jwriter.Writer, in Tag) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"post_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.PostId))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Tag) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson13673cd6EncodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Tag) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson13673cd6EncodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Tag) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson13673cd6DecodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Tag) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson13673cd6DecodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(l, v)
}