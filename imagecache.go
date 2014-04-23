package oiio

/*
#include "stdlib.h"

#include "cpp/oiio.h"

*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

// Define an API to an abstract class that manages image files, caches of open
// file handles as well as tiles of pixels so that truly huge amounts of image
// data may be accessed by an application with low memory footprint.
type ImageCache struct {
	ptr unsafe.Pointer
}

func newImageCache(i unsafe.Pointer) *ImageCache {
	x := new(ImageCache)
	x.ptr = i
	runtime.SetFinalizer(x, destroyImageCache)
	return x
}

// Finalizer that makes sure to destroy the ImageCache, in case
// the user forgets.
func destroyImageCache(i *ImageCache) {
	if i.ptr != nil {
		C.ImageCache_Destroy(i.ptr, C.bool(false))
		i.ptr = nil
	}
}

// Destroy a ImageCache that was created using CreateImageCache().
// When 'teardown' parameter is set to true, it will fully destroy even a "shared" ImageCache.
func (i *ImageCache) Destroy(teardown bool) {
	if i.ptr != nil {
		C.ImageCache_Destroy(i.ptr, C.bool(teardown))
		i.ptr = nil
	}
}

// Return the last error generated by API calls.
// An nil error will be returned if no error has occured.
func (i *ImageCache) LastError() error {
	err := C.GoString(C.ImageCache_geterror(i.ptr))
	if err == "" {
		return nil
	}
	return errors.New(err)
}

// Create an ImageCache. This should be freed by calling ImageCache.destroy()
//
// If shared==true, it's intended to be shared with other like-minded owners
// in the same process who also ask for a shared cache.
//
// If false, a private image cache will be created.
func CreateImageCache(shared bool) *ImageCache {
	ptr := C.ImageCache_Create(C.bool(false))
	return newImageCache(ptr)
}

// Close everything, free resources, start from scratch.
func (i *ImageCache) Clear() {
	C.ImageCache_clear(i.ptr)
}

// Return the statistics output as a huge string.
// Suitable default for level == 1
func (i *ImageCache) GetStats(level int) string {
	c_stats := C.ImageCache_getstats(i.ptr, C.int(level))
	stats := C.GoString(c_stats)
	return stats
}

func (i *ImageCache) ResetStats() {
	C.ImageCache_reset_stats(i.ptr)
}

func (i *ImageCache) Invalidate(filename string) {
	c_str := C.CString(filename)
	defer C.free(unsafe.Pointer(c_str))

	C.ImageCache_invalidate(i.ptr, c_str)
}

func (i *ImageCache) InvalidateAll(force bool) {
	C.ImageCache_invalidate_all(i.ptr, C.bool(force))
}