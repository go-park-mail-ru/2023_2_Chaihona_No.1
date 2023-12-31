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

func easyjson29bda1bDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(in *jlexer.Lexer, out *Profiles) {
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
		case "profiles":
			if in.IsNull() {
				in.Skip()
				out.Profiles = nil
			} else {
				in.Delim('[')
				if out.Profiles == nil {
					if !in.IsDelim(']') {
						out.Profiles = make([]model.Profile, 0, 0)
					} else {
						out.Profiles = []model.Profile{}
					}
				} else {
					out.Profiles = (out.Profiles)[:0]
				}
				for !in.IsDelim(']') {
					var v1 model.Profile
					(v1).UnmarshalEasyJSON(in)
					out.Profiles = append(out.Profiles, v1)
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
func easyjson29bda1bEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(out *jwriter.Writer, in Profiles) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"profiles\":"
		out.RawString(prefix[1:])
		if in.Profiles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Profiles {
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
func (v Profiles) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29bda1bEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Profiles) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29bda1bEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Profiles) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29bda1bDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Profiles) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29bda1bDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers(l, v)
}
func easyjson29bda1bDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(in *jlexer.Lexer, out *BodyProfile) {
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
		case "profile":
			(out.Profile).UnmarshalEasyJSON(in)
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
func easyjson29bda1bEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(out *jwriter.Writer, in BodyProfile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"profile\":"
		out.RawString(prefix[1:])
		(in.Profile).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BodyProfile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29bda1bEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BodyProfile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29bda1bEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BodyProfile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29bda1bDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BodyProfile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29bda1bDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers1(l, v)
}
func easyjson29bda1bDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers2(in *jlexer.Lexer, out *Analytics) {
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
		case "analytics":
			(out.Analytics).UnmarshalEasyJSON(in)
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
func easyjson29bda1bEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers2(out *jwriter.Writer, in Analytics) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"analytics\":"
		out.RawString(prefix[1:])
		(in.Analytics).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Analytics) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29bda1bEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Analytics) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29bda1bEncodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Analytics) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29bda1bDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Analytics) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29bda1bDecodeGithubComGoParkMailRu20232ChaihonaNo1InternalHandlers2(l, v)
}
