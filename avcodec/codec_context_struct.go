// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avformat provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
import "C"
import (
	"reflect"
	"unsafe"
)

func (cctxt *Context) Type() MediaType {
	return MediaType(cctxt.codec_type)
}

func (cctxt *Context) SetBitRate(br int64) {
	cctxt.bit_rate = C.int64_t(br)
}

func (cctxt *Context) GetCodecId() CodecId {
	return CodecId(cctxt.codec_id)
}

func (cctxt *Context) SetCodecId(codecId CodecId) {
	cctxt.codec_id = C.enum_AVCodecID(codecId)
}

func (cctxt *Context) GetCodecType() MediaType {
	return MediaType(cctxt.codec_type)
}

func (cctxt *Context) SetCodecType(ctype MediaType) {
	cctxt.codec_type = C.enum_AVMediaType(ctype)
}

func (cctxt *Context) GetTimeBase() Rational {
	return NewRational(int(cctxt.time_base.num), int(cctxt.time_base.den))
}

func (cctxt *Context) SetTimeBase(timeBase Rational) {
	cctxt.time_base.num = C.int(timeBase.Num())
	cctxt.time_base.den = C.int(timeBase.Den())
}

func (cctx *Context) GetWidth() int {
	return int(cctx.width)
}

func (cctx *Context) SetWidth(w int) {
	cctx.width = C.int(w)
}

func (cctx *Context) GetHeight() int {
	return int(cctx.height)
}

func (cctx *Context) SetHeight(h int) {
	cctx.height = C.int(h)
}

func (cctx *Context) GetPixelFormat() PixelFormat {
	return PixelFormat(C.int(cctx.pix_fmt))
}

func (cctx *Context) SetPixelFormat(fmt PixelFormat) {
	cctx.pix_fmt = C.enum_AVPixelFormat(C.int(fmt))
}

func (cctx *Context) GetFlags() int {
	return int(cctx.flags)
}

func (cctx *Context) SetFlags(flags int) {
	cctx.flags = C.int(flags)
}

func (cctx *Context) GetMeRange() int {
	return int(cctx.me_range)
}

func (cctx *Context) SetMeRange(r int) {
	cctx.me_range = C.int(r)
}

func (cctx *Context) GetMaxQDiff() int {
	return int(cctx.max_qdiff)
}

func (cctx *Context) SetMaxQDiff(v int) {
	cctx.max_qdiff = C.int(v)
}

func (cctx *Context) GetQMin() int {
	return int(cctx.qmin)
}

func (cctx *Context) SetQMin(v int) {
	cctx.qmin = C.int(v)
}

func (cctx *Context) GetQMax() int {
	return int(cctx.qmax)
}

func (cctx *Context) SetQMax(v int) {
	cctx.qmax = C.int(v)
}

func (cctx *Context) GetQCompress() float32 {
	return float32(cctx.qcompress)
}

func (cctx *Context) SetQCompress(v float32) {
	cctx.qcompress = C.float(v)
}

func (cctx *Context) GetExtraData() []byte {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cctx.extradata)),
		Len:  int(cctx.extradata_size),
		Cap:  int(cctx.extradata_size),
	}

	return *((*[]byte)(unsafe.Pointer(&header)))
}

func (cctx *Context) SetExtraData(data []byte) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))

	cctx.extradata = (*C.uint8_t)(unsafe.Pointer(header.Data))
	cctx.extradata_size = C.int(header.Len)
}

func (cctx *Context) Release() {
	C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(cctx)))
	C.av_freep(unsafe.Pointer(cctx))
}
