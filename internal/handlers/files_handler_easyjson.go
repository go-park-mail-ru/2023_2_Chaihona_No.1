// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package handlers

import (
	json "encoding/json"
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
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

func easyjson348e992eDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(in *jlexer.Lexer, out *FileBody) {
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
		case "path":
			out.Path = string(in.String())
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
func easyjson348e992eEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(out *jwriter.Writer, in FileBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"path\":"
		out.RawString(prefix[1:])
		out.String(string(in.Path))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FileBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson348e992eEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FileBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson348e992eEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FileBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson348e992eDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FileBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson348e992eDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(l, v)
}
func easyjson348e992eDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(in *jlexer.Lexer, out *BodyAttaches) {
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
		case "attaches":
			if in.IsNull() {
				in.Skip()
				out.Attaches = nil
			} else {
				in.Delim('[')
				if out.Attaches == nil {
					if !in.IsDelim(']') {
						out.Attaches = make([]model.Attach, 0, 0)
					} else {
						out.Attaches = []model.Attach{}
					}
				} else {
					out.Attaches = (out.Attaches)[:0]
				}
				for !in.IsDelim(']') {
					var v1 model.Attach
					(v1).UnmarshalEasyJSON(in)
					out.Attaches = append(out.Attaches, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson348e992eEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(out *jwriter.Writer, in BodyAttaches) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"attaches\":"
		out.RawString(prefix[1:])
		if in.Attaches == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Attaches {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BodyAttaches) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson348e992eEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BodyAttaches) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson348e992eEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BodyAttaches) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson348e992eDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BodyAttaches) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson348e992eDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(l, v)
}
