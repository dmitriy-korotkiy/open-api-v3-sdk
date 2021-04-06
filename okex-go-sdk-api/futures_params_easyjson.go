// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package okex

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

func easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi(in *jlexer.Lexer, out *FuturesOrdersParams) {
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
		case "Currency":
			out.Currency = string(in.String())
		case "Status":
			out.Status = int(in.Int())
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
func easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi(out *jwriter.Writer, in FuturesOrdersParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Currency\":"
		out.RawString(prefix[1:])
		out.String(string(in.Currency))
	}
	{
		const prefix string = ",\"Status\":"
		out.RawString(prefix)
		out.Int(int(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FuturesOrdersParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FuturesOrdersParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FuturesOrdersParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FuturesOrdersParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi(l, v)
}
func easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi1(in *jlexer.Lexer, out *FuturesNewOrderParams) {
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
		case "instrument_id":
			out.InstrumentId = string(in.String())
		case "leverage":
			out.Leverage = string(in.String())
		case "client_oid":
			out.ClientOid = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "price":
			out.Price = string(in.String())
		case "size":
			out.Size = string(in.String())
		case "match_price":
			out.MatchPrice = string(in.String())
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
func easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi1(out *jwriter.Writer, in FuturesNewOrderParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"instrument_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.InstrumentId))
	}
	{
		const prefix string = ",\"leverage\":"
		out.RawString(prefix)
		out.String(string(in.Leverage))
	}
	{
		const prefix string = ",\"client_oid\":"
		out.RawString(prefix)
		out.String(string(in.ClientOid))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.String(string(in.Price))
	}
	{
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.String(string(in.Size))
	}
	{
		const prefix string = ",\"match_price\":"
		out.RawString(prefix)
		out.String(string(in.MatchPrice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FuturesNewOrderParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FuturesNewOrderParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FuturesNewOrderParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FuturesNewOrderParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi1(l, v)
}
func easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi2(in *jlexer.Lexer, out *FuturesFillsParams) {
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
		case "order_id":
			out.OrderId = string(in.String())
		case "instrument_id":
			out.InstrumentId = string(in.String())
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
func easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi2(out *jwriter.Writer, in FuturesFillsParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"order_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.OrderId))
	}
	{
		const prefix string = ",\"instrument_id\":"
		out.RawString(prefix)
		out.String(string(in.InstrumentId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FuturesFillsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FuturesFillsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FuturesFillsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FuturesFillsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi2(l, v)
}
func easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi3(in *jlexer.Lexer, out *FuturesClosePositionParams) {
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
		case "ClosePositionData":
			if in.IsNull() {
				in.Skip()
				out.ClosePositionData = nil
			} else {
				in.Delim('[')
				if out.ClosePositionData == nil {
					if !in.IsDelim(']') {
						out.ClosePositionData = make([]ClosePositionData, 0, 1)
					} else {
						out.ClosePositionData = []ClosePositionData{}
					}
				} else {
					out.ClosePositionData = (out.ClosePositionData)[:0]
				}
				for !in.IsDelim(']') {
					var v1 ClosePositionData
					(v1).UnmarshalEasyJSON(in)
					out.ClosePositionData = append(out.ClosePositionData, v1)
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
func easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi3(out *jwriter.Writer, in FuturesClosePositionParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ClosePositionData\":"
		out.RawString(prefix[1:])
		if in.ClosePositionData == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.ClosePositionData {
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
func (v FuturesClosePositionParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FuturesClosePositionParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FuturesClosePositionParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FuturesClosePositionParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi3(l, v)
}
func easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi4(in *jlexer.Lexer, out *FuturesBatchNewOrderParams) {
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
		case "instrument_id":
			out.InstrumentId = string(in.String())
		case "leverage":
			out.Leverage = string(in.String())
		case "orders_data":
			out.OrdersData = string(in.String())
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
func easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi4(out *jwriter.Writer, in FuturesBatchNewOrderParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"instrument_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.InstrumentId))
	}
	{
		const prefix string = ",\"leverage\":"
		out.RawString(prefix)
		out.String(string(in.Leverage))
	}
	{
		const prefix string = ",\"orders_data\":"
		out.RawString(prefix)
		out.String(string(in.OrdersData))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FuturesBatchNewOrderParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FuturesBatchNewOrderParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FuturesBatchNewOrderParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FuturesBatchNewOrderParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi4(l, v)
}
func easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi5(in *jlexer.Lexer, out *FuturesBatchNewOrderItem) {
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
		case "client_oid":
			out.ClientOid = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "price":
			out.Price = string(in.String())
		case "size":
			out.Size = string(in.String())
		case "match_price":
			out.MatchPrice = string(in.String())
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
func easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi5(out *jwriter.Writer, in FuturesBatchNewOrderItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"client_oid\":"
		out.RawString(prefix[1:])
		out.String(string(in.ClientOid))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.String(string(in.Price))
	}
	{
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.String(string(in.Size))
	}
	{
		const prefix string = ",\"match_price\":"
		out.RawString(prefix)
		out.String(string(in.MatchPrice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FuturesBatchNewOrderItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FuturesBatchNewOrderItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FuturesBatchNewOrderItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FuturesBatchNewOrderItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi5(l, v)
}
func easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi6(in *jlexer.Lexer, out *ClosePositionData) {
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
		case "instrument_id":
			out.InstrumentId = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "lever_rate":
			out.LeverRate = string(in.String())
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
func easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi6(out *jwriter.Writer, in ClosePositionData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"instrument_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.InstrumentId))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"lever_rate\":"
		out.RawString(prefix)
		out.String(string(in.LeverRate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClosePositionData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClosePositionData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson420bb3e9EncodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClosePositionData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClosePositionData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson420bb3e9DecodeGithubComDarkfoxs96OpenApiV3SdkOkexGoSdkApi6(l, v)
}
