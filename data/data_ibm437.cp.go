// This file is automatically generated by generate-charset-data.
// Do not hand-edit.

package data

import (
	"github.com/ezkl/go-charset/charset"
	"io"
	"io/ioutil"
	"strings"
)

func init() {
	charset.RegisterDataFile("ibm437.cp", func() (io.ReadCloser, error) {
		r := strings.NewReader("\x00\x01\x02\x03\x04\x05\x06\a\b\t\n\v\f\r\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\u007fÇüéâäàåçêëèïîìÄÅÉæÆôöòûùÿÖÜ¢£¥₧ƒáíóúñÑªº¿⌐¬½¼¡«»░▒▓│┤╡╢╖╕╣║╗╝╜╛┐└┴┬├─┼╞╟╚╔╩╦╠═╬╧╨╤╥╙╘╒╓╫╪┘┌█▄▌▐▀αßΓπΣσµτΦΘΩδ∞∅∈∩≡±≥≤⌠⌡÷≈°•·√ⁿ²∎\u00a0")
		return ioutil.NopCloser(r), nil
	})
}
