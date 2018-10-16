// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

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

func easyjsonCef4e921DecodeGithubComKamilskFormApiPkgDomain(in *jlexer.Lexer, out *Schema) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "lang":
			out.Language = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "action":
			out.Action = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "enctype":
			out.EncodingType = string(in.String())
		case "input":
			if in.IsNull() {
				in.Skip()
				out.Inputs = nil
			} else {
				in.Delim('[')
				if out.Inputs == nil {
					if !in.IsDelim(']') {
						out.Inputs = make([]Input, 0, 1)
					} else {
						out.Inputs = []Input{}
					}
				} else {
					out.Inputs = (out.Inputs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Input
					(v1).UnmarshalEasyJSON(in)
					out.Inputs = append(out.Inputs, v1)
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
func easyjsonCef4e921EncodeGithubComKamilskFormApiPkgDomain(out *jwriter.Writer, in Schema) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"lang\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Language))
	}
	{
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"action\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Action))
	}
	if in.Method != "" {
		const prefix string = ",\"method\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Method))
	}
	if in.EncodingType != "" {
		const prefix string = ",\"enctype\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EncodingType))
	}
	{
		const prefix string = ",\"input\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Inputs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Inputs {
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
func (v Schema) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCef4e921EncodeGithubComKamilskFormApiPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Schema) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCef4e921EncodeGithubComKamilskFormApiPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Schema) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCef4e921DecodeGithubComKamilskFormApiPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Schema) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCef4e921DecodeGithubComKamilskFormApiPkgDomain(l, v)
}
