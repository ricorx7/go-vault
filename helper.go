package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

// genToken will create a unique token
// for a web form.  The token will ensure
// a double submit will not be sent.
func genToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

// OptionItem will help with comboboxes.
// A a selected option needs to be known.
// Use this to create a slice of ListOptions
// for the combobox to display.  The selected flag
// will be set for the selected item.  Only 1 selected
// flag should be set in a slice.
type OptionItem struct {
	Value    string // The value to equal when the option is selected
	Text     string // The text to display for the option
	Selected bool   // Flag if the option is selected
}
