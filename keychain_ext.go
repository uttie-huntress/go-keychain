//go:build darwin
// +build darwin

package keychain

// See https://developer.apple.com/library/ios/documentation/Security/Reference/keychainservices/index.html for the APIs used below.

// Also see https://developer.apple.com/library/ios/documentation/Security/Conceptual/keychainServConcepts/01introduction/introduction.html .

/*
#cgo LDFLAGS: -framework CoreFoundation -framework Security

#include <CoreFoundation/CoreFoundation.h>
#include <Security/Security.h>
*/
import "C"
import (
	"unsafe"
)

// Keychain represents the path to a specific OSX keychain
type Keychain struct {
	path string
}

func NewKeychain(path string) Keychain {
	return Keychain{path: path}
}

// The returned SecKeychainRef, if non-nil, must be released via CFRelease.
func openKeychainRef(path string) (C.SecKeychainRef, error) {
	pathName := C.CString(path)
	defer C.free(unsafe.Pointer(pathName))

	var kref C.SecKeychainRef
	if err := checkError(C.SecKeychainOpen(pathName, &kref)); err != nil {
		return 0, err
	}
	return kref, nil
}

// The returned CFTypeRef, if non-nil, must be released via CFRelease.
func (kc Keychain) Convert() (C.CFTypeRef, error) {
	keyRef, err := openKeychainRef(kc.path)
	return C.CFTypeRef(keyRef), err
}

var KeychainKey = attrKey(C.CFTypeRef(C.kSecUseKeychain))

func (k *Item) UseKeychain(kc Keychain) {
	k.attr[KeychainKey] = kc
}

// Status returns the status of the keychain
func (kc Keychain) Status() error {
	// returns no error even if it doesn't exist
	kref, err := openKeychainRef(kc.path)
	if err != nil {
		return err
	}
	defer C.CFRelease(C.CFTypeRef(kref))

	var status C.SecKeychainStatus
	return checkError(C.SecKeychainGetStatus(kref, &status))
}

// releaseKey releases the memory - used for testing
func releaseKey(kref C.SecKeychainRef) {
	Release(C.CFTypeRef(kref))
}
