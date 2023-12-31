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

func easyjsonE01deb61DecodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(in *jlexer.Lexer, out *Goal) {
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
		case "goal_type":
			out.GoalType = string(in.String())
		case "currency":
			out.Currency = string(in.String())
		case "current":
			out.Current = float64(in.Float64())
		case "goal_value":
			out.GoalValue = float64(in.Float64())
		case "description":
			out.Description = string(in.String())
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
func easyjsonE01deb61EncodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(out *jwriter.Writer, in Goal) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"goal_type\":"
		out.RawString(prefix)
		out.String(string(in.GoalType))
	}
	if in.Currency != "" {
		const prefix string = ",\"currency\":"
		out.RawString(prefix)
		out.String(string(in.Currency))
	}
	{
		const prefix string = ",\"current\":"
		out.RawString(prefix)
		out.Float64(float64(in.Current))
	}
	{
		const prefix string = ",\"goal_value\":"
		out.RawString(prefix)
		out.Float64(float64(in.GoalValue))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Goal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE01deb61EncodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Goal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE01deb61EncodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Goal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE01deb61DecodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Goal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE01deb61DecodeGithubComGoParkMailRu20232ChaihonaNo1InternalModel(l, v)
}
