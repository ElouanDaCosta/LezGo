package pkg

import (
	"bufio"
	"bytes"
	"container/ring"
	"fmt"
	"os"
	"time"
	"unicode/utf8"
)

var (
	space           = []byte{32}
	backSpace       = []byte{8}
	clearFromCursor = []byte{27, 91, 75}

	//Spin is a standard spinner
	Spin = []string{"-", "\\", "|", "/"}

	//Dot is an ellipsis sequence
	Dot = []string{"", ".", "..", "..."}

	//Grow is a . that grows up to be an O
	Grow = []string{".", "ₒ", "ⱺ", "O", "ⱺ", "ₒ", "."}
)

// Throbber starts a throbber
func Throbber(
	vals []string,
	speed time.Duration,
	c chan bool,
) {
	throb("", "", vals, speed, c)
}

// ThrobberPrefixed starts a throbber with a given static prefix
func ThrobberPrefixed(
	prefix string,
	vals []string,
	speed time.Duration,
	c chan bool,
) {
	throb(prefix, "", vals, speed, c)
}

// ThrobberPostfixed starts a throbber with a given static postfix
func ThrobberPostfixed(
	postfix string,
	vals []string,
	speed time.Duration,
	c chan bool,
) {
	throb("", postfix, vals, speed, c)
}

// ThrobberPrefixedAndPostfixed starts a throbber with a given static prefix and
// postfix
func ThrobberPrefixedAndPostfixed(
	prefix string,
	postfix string,
	vals []string,
	speed time.Duration,
	c chan bool,
) {
	throb(prefix, postfix, vals, speed, c)
}

func throb(
	prefix string,
	postfix string,
	vals []string,
	speed time.Duration,
	c chan bool,
) {
	var (
		buffer        = bufio.NewWriter(os.Stdout)
		sequence, max = ringFromStringSlice(vals)
	)
	for {
		select {
		case <-c:
			return
		default:
			step := sequence.Value.(string)
			replaceLineAfter(
				buffer,
				[]byte(fmt.Sprintf(
					"%s%s%s%s",
					prefix,
					step,
					bytes.Repeat(space, max-utf8.RuneCountInString(step)),
					postfix,
				)),
				speed,
			)
			sequence = sequence.Next()
		}
	}
}

func replaceLineAfter(
	buffer *bufio.Writer,
	s []byte,
	wait time.Duration,
) {
	var err error
	defer buffer.Flush()
	_, err = buffer.WriteString(fmt.Sprintf(
		"%s%s",
		bytes.Repeat(backSpace, len(s)),
		clearFromCursor,
	))
	if err != nil {
		return
	}
	_, err = buffer.Write(s)
	if err != nil {
		return
	}
	time.Sleep(wait)
}

func ringFromStringSlice(vals []string) (*ring.Ring, int) {
	var (
		r   = ring.New(len(vals))
		max int
	)
	for _, v := range vals {
		if l := utf8.RuneCountInString(v); l > max {
			max = l
		}
		r.Value = v
		r = r.Next()
	}
	return r, max
}
