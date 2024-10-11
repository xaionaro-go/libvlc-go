//go:build libvlc_static
// +build libvlc_static

package vlc

// #cgo LDFLAGS: -Wl,-Bstatic -lvlc -Wl,-Bdynamic
import "C"
