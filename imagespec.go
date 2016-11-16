package oiio

/*
#include "stdlib.h"

#include "cpp/oiio.h"

*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

// ImageSpec describes the data format of an image – dimensions, layout,
// number and meanings of image channels.
type ImageSpec struct {
	ptr unsafe.Pointer
}

func newImageSpec(i unsafe.Pointer) *ImageSpec {
	spec := &ImageSpec{ptr: i}
	runtime.SetFinalizer(spec, deleteImageSpec)
	return spec
}

func deleteImageSpec(i *ImageSpec) {
	if i.ptr != nil {
		C.free(i.ptr)
		i.ptr = nil
	}
}

// given just the data format, set the default quantize and set all other channels to something reasonable.
func NewImageSpec(format TypeDesc) *ImageSpec {
	spec := C.ImageSpec_New((C.TypeDesc)(format))
	return newImageSpec(spec)
}

// for simple 2D scanline image with nothing special. If fmt is not supplied, default to unsigned 8-bit data.
func NewImageSpecSize(x, y, chans int, format TypeDesc) *ImageSpec {
	spec := C.ImageSpec_New_Size(C.int(x), C.int(y), C.int(chans), (C.TypeDesc)(format))
	return newImageSpec(spec)
}

// Set the channelnames to reasonable defaults ("R", "G", "B", "A"),
// and alpha_channel, based on the number of channels.
func (s *ImageSpec) DefaultChannelNames() {
	C.ImageSpec_default_channel_names(s.ptr)
}

// Return the number of bytes for each channel datum, assuming they are
// all stored using the data format given by format.
func (s *ImageSpec) ChannelBytes() int {
	return int(C.ImageSpec_channel_bytes(s.ptr))
}

// Return the number of bytes needed for the single specified channel.
// If native is false (default), compute the size of one channel of format,
// but if native is true, compute the size of the channel in terms of the "native"
// data format of that channel as stored in the file.
func (s *ImageSpec) ChannelBytesChan(chanNum int, native bool) int {
	return int(C.ImageSpec_channel_bytes_chan(s.ptr, C.int(chanNum), C.bool(native)))
}

// Return the number of bytes for each pixel (counting all channels). If native is
// false, assume all channels are in format, but if native is true,
// compute the size of a pixel in the "native" data format of the file (these may
// differ in the case of per-channel formats). This will return a max value
// in the event of an overflow where it's not representable in a int.
func (s *ImageSpec) PixelBytes(native bool) int {
	return int(C.ImageSpec_pixel_bytes(s.ptr, C.bool(native)))
}

// Return the number of bytes for just the subset of channels in each pixel described
// by [chanBegin, chanEnd). If native is false, assume all channels are in format,
// but if native is true, compute the size of a pixel in the "native" data format of
// the file (these may differ in the case of per-channel formats). This will return
// a max value in the event of an overflow where it's not representable in a int.
func (s *ImageSpec) PixelBytesChans(chanBegin, chanEnd int, native bool) int {
	return int(C.ImageSpec_pixel_bytes_chans(s.ptr, C.int(chanBegin), C.int(chanEnd), C.bool(native)))
}

// Return the number of bytes for each scanline.
// This will return a max value in the event of an overflow where it's not
// representable in an int. If native is false, assume all channels are in
// format, but if native is true, compute the size of a pixel in the "native"
// data format of the file (these may differ in the case of per-channel formats).
func (s *ImageSpec) ScanlineBytes(native bool) int {
	return int(C.ImageSpec_scanline_bytes(s.ptr, C.bool(native)))
}

// Return the number of pixels for a tile. This will return a max value in the event
// of an overflow where it's not representable in an int.
func (s *ImageSpec) TilePixels() int {
	return int(C.ImageSpec_tile_pixels(s.ptr))
}

// Return the number of bytes for each a tile of the image.
// This will return a max value in the event of an overflow where it's not
// representable in an imagesize_t. If native is false, assume all channels are
// in format, but if native is true, compute the size of a pixel in the "native"
// data format of the file (these may differ in the case of per-channel formats).
func (s *ImageSpec) TileBytes(native bool) int {
	return int(C.ImageSpec_tile_bytes(s.ptr, C.bool(native)))
}

// Return the number of pixels for an entire image.
// This will return a max value in the event of an overflow where
// it's not representable in an int.
func (s *ImageSpec) ImagePixels() int {
	return int(C.ImageSpec_image_pixels(s.ptr))
}

// Return the number of bytes for an entire image.
// This will return a max value in the event of an overflow where it's not
// representable in an int. If native is false, assume all channels are in
// format, but if native is true, compute the size of a pixel in the "native"
// data format of the file (these may differ in the case of per-channel formats).
func (s *ImageSpec) ImageBytes(native bool) int {
	return int(C.ImageSpec_image_bytes(s.ptr, C.bool(native)))
}

// Verify that on this platform, a size_t is big enough to hold the number of bytes
// (and pixels) in a scanline, a tile, and the whole image. If this returns false,
// the image is much too big to allocate and read all at once, so client apps beware
// and check these routines for overflows!
func (s *ImageSpec) SizeSafe() bool {
	return bool(C.ImageSpec_size_safe(s.ptr))
}

func (s *ImageSpec) ChannelFormat(chanNum int) TypeDesc {
	return (TypeDesc)(C.ImageSpec_channelformat(s.ptr, C.int(chanNum)))
}

// Properties
func (s *ImageSpec) X() int {
	return int(C.ImageSpec_x(s.ptr))
}

func (s *ImageSpec) SetX(val int) {
	C.ImageSpec_set_x(s.ptr, C.int(val))
}

func (s *ImageSpec) Y() int {
	return int(C.ImageSpec_y(s.ptr))
}

func (s *ImageSpec) SetY(val int) {
	C.ImageSpec_set_y(s.ptr, C.int(val))
}

// origin (upper left corner) of pixel data
func (s *ImageSpec) Z() int {
	return int(C.ImageSpec_z(s.ptr))
}

func (s *ImageSpec) SetZ(val int) {
	C.ImageSpec_set_z(s.ptr, C.int(val))
}

// width of the pixel data window
func (s *ImageSpec) Width() int {
	return int(C.ImageSpec_width(s.ptr))
}

func (s *ImageSpec) SetWidth(val int) {
	C.ImageSpec_set_width(s.ptr, C.int(val))
}

// height of the pixel data window
func (s *ImageSpec) Height() int {
	return int(C.ImageSpec_height(s.ptr))
}

func (s *ImageSpec) SetHeight(val int) {
	C.ImageSpec_set_height(s.ptr, C.int(val))
}

// depth of pixel data, >1 indicates a "volume"
func (s *ImageSpec) Depth() int {
	return int(C.ImageSpec_depth(s.ptr))
}

func (s *ImageSpec) SetDepth(val int) {
	C.ImageSpec_set_depth(s.ptr, C.int(val))
}

// origin of the full (display) window
func (s *ImageSpec) FullX() int {
	return int(C.ImageSpec_full_x(s.ptr))
}

func (s *ImageSpec) SetFullX(val int) {
	C.ImageSpec_set_full_x(s.ptr, C.int(val))
}

// origin of the full (display) window
func (s *ImageSpec) FullY() int {
	return int(C.ImageSpec_full_y(s.ptr))
}

func (s *ImageSpec) SetFullY(val int) {
	C.ImageSpec_set_full_y(s.ptr, C.int(val))
}

// origin of the full (display) window
func (s *ImageSpec) FullZ() int {
	return int(C.ImageSpec_full_z(s.ptr))
}

func (s *ImageSpec) SetFullZ(val int) {
	C.ImageSpec_set_full_z(s.ptr, C.int(val))
}

// width of the full (display) window
func (s *ImageSpec) FullWidth() int {
	return int(C.ImageSpec_full_width(s.ptr))
}

func (s *ImageSpec) SetFullWidth(val int) {
	C.ImageSpec_set_full_width(s.ptr, C.int(val))
}

// height of the full (display) window
func (s *ImageSpec) FullHeight() int {
	return int(C.ImageSpec_full_height(s.ptr))
}

func (s *ImageSpec) SetFullHeight(val int) {
	C.ImageSpec_set_full_height(s.ptr, C.int(val))
}

// depth of the full (display) window
func (s *ImageSpec) FullDepth() int {
	return int(C.ImageSpec_full_depth(s.ptr))
}

func (s *ImageSpec) SetFullDepth(val int) {
	C.ImageSpec_set_full_depth(s.ptr, C.int(val))
}

// tile width (0 for a non-tiled image)
func (s *ImageSpec) TileWidth() int {
	return int(C.ImageSpec_tile_width(s.ptr))
}

func (s *ImageSpec) SetTileWidth(val int) {
	C.ImageSpec_set_tile_width(s.ptr, C.int(val))
}

// tile height (0 for a non-tiled image)
func (s *ImageSpec) TileHeight() int {
	return int(C.ImageSpec_tile_height(s.ptr))
}

func (s *ImageSpec) SetTileHeight(val int) {
	C.ImageSpec_set_tile_height(s.ptr, C.int(val))
}

func (s *ImageSpec) TileDepth() int {
	return int(C.ImageSpec_tile_depth(s.ptr))
}

func (s *ImageSpec) SetTileDepth(val int) {
	C.ImageSpec_set_tile_depth(s.ptr, C.int(val))
}

// number of image channels, e.g., 4 for RGBA
func (s *ImageSpec) NumChannels() int {
	return int(C.ImageSpec_nchannels(s.ptr))
}

func (s *ImageSpec) SetNumChannels(val int) {
	C.ImageSpec_set_nchannels(s.ptr, C.int(val))
}

// data format of the channels
func (s *ImageSpec) Format() TypeDesc {
	return (TypeDesc)(C.ImageSpec_format(s.ptr))
}

// Set the data format, and as a side effect set quantize to good defaults for that format
func (s *ImageSpec) SetFormat(format TypeDesc) {
	C.ImageSpec_set_format(s.ptr, (C.TypeDesc)(format))
}

// Optional per-channel formats.
func (s *ImageSpec) ChannelFormats() []TypeDesc {
	formats := make([]TypeDesc, s.NumChannels())
	formats_ptr := (*C.TypeDesc)(unsafe.Pointer(&formats[0]))
	C.ImageSpec_channelformats(s.ptr, formats_ptr)
	return formats
}

func (s *ImageSpec) SetChannelFormats(formats []TypeDesc) {
	formats_ptr := (*C.TypeDesc)(unsafe.Pointer(&formats[0]))
	C.ImageSpec_set_channelformats(s.ptr, formats_ptr)
}

// String name of each channel
func (s *ImageSpec) ChannelNames() []string {
	names := make([]string, s.NumChannels())
	c_names := make([]*C.char, s.NumChannels())
	c_names_ptr := (**C.char)(unsafe.Pointer(&c_names[0]))
	C.ImageSpec_channelnames(s.ptr, c_names_ptr)
	for i, c := range c_names {
		names[i] = C.GoString(c)
	}
	return names
}

// SetChannelNames re-labels each existing channel,
// from a slice of string names.
func (s *ImageSpec) SetChannelNames(names []string) {
	c_names := make([]*C.char, len(names))
	for i, n := range names {
		c_names[i] = C.CString(n)
		defer C.free(unsafe.Pointer(c_names[i]))
	}
	c_names_ptr := (**C.char)(unsafe.Pointer(&c_names[0]))
	C.ImageSpec_set_channelnames(s.ptr, c_names_ptr)
}

// Convert ImageSpec class into XML string.
func (s *ImageSpec) ToXml() string {
	return C.GoString(C.ImageSpec_to_xml(s.ptr))
}

// Index of alpha channel, or -1 if not known.
func (s *ImageSpec) AlphaChannel() int {
	return int(C.ImageSpec_alpha_channel(s.ptr))
}

func (s *ImageSpec) SetAlphaChannel(val int) {
	C.ImageSpec_set_alpha_channel(s.ptr, C.int(val))
}

// Index of depth channel, or -1 if not known.
func (s *ImageSpec) ZChannel() int {
	return int(C.ImageSpec_z_channel(s.ptr))
}

// Set the index of the depth channel.
func (s *ImageSpec) SetZChannel(val int) {
	C.ImageSpec_set_z_channel(s.ptr, C.int(val))
}

// Contains deep data.
func (s *ImageSpec) Deep() bool {
	return bool(C.ImageSpec_deep(s.ptr))
}

func (s *ImageSpec) SetDeep(val bool) {
	C.ImageSpec_set_deep(s.ptr, C.bool(val))
}

// SetAttribute sets a metadata value in the extra attribs. Acceptable types are
// string, int, and float32.
//
// Example:
// 		s = NewImageSpec(...)
// 		s.SetAttribute("foo_str", "blah")
// 		s.SetAttribute("foo_int", 14)
// 		s.SetAttribute("foo_float", 3.14)
func (s *ImageSpec) SetAttribute(name string, val interface{}) error {
	c_str := C.CString(name)
	defer C.free(unsafe.Pointer(c_str))

	switch t := val.(type) {
	case string:
		c_val := C.CString(t)
		defer C.free(unsafe.Pointer(c_val))
		C.ImageSpec_attribute_char(s.ptr, c_str, c_val)
	case float32:
		C.ImageSpec_attribute_float(s.ptr, c_str, C.float(t))
	case int:
		C.ImageSpec_attribute_int(s.ptr, c_str, C.int(t))
	default:
		return fmt.Errorf("Value type %T is not one of (string, int, float32)", t)
	}
	return nil
}

// AttributeString looks up an existing attrib by name and returns
// the string value, or a default value if it does not exist.
// Default value is an empty string, if not specified.
func (s *ImageSpec) AttributeString(name string, defaultVal ...string) string {
	var defVal string
	if len(defaultVal) > 0 {
		defVal = defaultVal[0]
	}
	c_str := C.CString(name)
	c_val := C.CString(defVal)
	defer func() {
		C.free(unsafe.Pointer(c_str))
		C.free(unsafe.Pointer(c_val))
	}()

	return C.GoString(C.ImageSpec_get_string_attribute(s.ptr, c_str, c_val))
}

// AttributeFloat looks up an existing attrib by name and returns
// the float value, or a default value if it does not exist.
// Default value is 0, if not specified.
func (s *ImageSpec) AttributeFloat(name string, defaultVal ...float32) float32 {
	var defVal float32
	if len(defaultVal) > 0 {
		defVal = defaultVal[0]
	}
	c_str := C.CString(name)
	defer C.free(unsafe.Pointer(c_str))
	return float32(C.ImageSpec_get_float_attribute(s.ptr, c_str, C.float(defVal)))
}

// AttributeInt looks up an existing attrib by name and returns
// the int value, or a default value if it does not exist.
// Default value is 0, if not specified.
func (s *ImageSpec) AttributeInt(name string, defaultVal ...int) int {
	var defVal int
	if len(defaultVal) > 0 {
		defVal = defaultVal[0]
	}
	c_str := C.CString(name)
	defer C.free(unsafe.Pointer(c_str))
	return int(C.ImageSpec_get_int_attribute(s.ptr, c_str, C.int(defVal)))
}

// EraseAttribute removes the specified attribute from the list of extra_attribs.
// If not found, do nothing.
// If caseSensitive is true, the name search will be case-sensitive, otherwise the name
// search will be performed without regard to case
func (s *ImageSpec) EraseAttribute(name string, caseSensitive bool) {
	c_str := C.CString(name)
	defer C.free(unsafe.Pointer(c_str))

	C.ImageSpec_erase_attribute(s.ptr, c_str, C.TYPE_UNKNOWN, C.bool(caseSensitive))
}

// EraseAttributeType removes the specified attribute from the list of extra_attribs.
// If not found, do nothing.
// If searchtype is anything but TypeUnknown, restrict matches to only those of
// the given type.
// If caseSensitive is true, the name search will be case-sensitive, otherwise
// the name search will be performed without regard to case
func (s *ImageSpec) EraseAttributeType(name string, searchType TypeDesc, caseSensitive bool) {
	c_str := C.CString(name)
	defer C.free(unsafe.Pointer(c_str))

	C.ImageSpec_erase_attribute(s.ptr, c_str, C.TypeDesc(searchType), C.bool(caseSensitive))
}
