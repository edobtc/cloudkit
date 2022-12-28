package logging

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var structuredFormatterTests = []struct {
	in  string
	out string
}{
	{"fluentd", "*log.Formatter"},
	{"stackdriver", "*log.Formatter"},
	{"structured", "*log.Formatter"},
}

func TestStackDriverSwitching(t *testing.T) {
	for _, tt := range structuredFormatterTests {
		t.Run(tt.in, func(t *testing.T) {
			formatter := loadFormatter(tt.in)
			concreteType := reflect.TypeOf(formatter)
			result := fmt.Sprintf("%v", concreteType)

			if result != tt.out {
				t.Errorf("got %s, expected %v", formatter, tt.out)
			}
		})
	}
}

var structuredFormatterTypeTests = []struct {
	in     string
	out    string
	result bool
}{
	// Success
	{"fluentd", `{"message":"test this","severity":"INFO","timestamp"`, true},
	{"stackdriver", `{"message":"test this","severity":"INFO","timestamp"`, true},
	{"structured", `{"message":"test this","severity":"INFO","timestamp"`, true},
	{"json", `{"level":"info","msg":"test this"`, true},
	{"default", `level=info msg="test this"`, true},

	// Fails
	// fluentD json formatter
	{"fluentd", `level=info msg="test this"`, false},
	{"stackdriver", `level=info msg="test this"`, false},
	{"structured", `level=info msg="test this"`, false},

	// JSON fails
	{"json", `{"message":"test this","severity":"INFO","timestamp"`, false},
	{"json", `level=info msg="test this"`, false},

	// Text with pure string
	{"default", `rando"`, false},
}

func TestFormatterTypes(t *testing.T) {
	for _, tt := range structuredFormatterTypeTests {
		t.Run(tt.in, func(t *testing.T) {
			var b bytes.Buffer
			wr := io.Writer(&b)
			log.SetOutput(wr)

			formatter := loadFormatter(tt.in)

			log.SetFormatter(formatter)

			log.Info("test this")

			match := strings.Contains(b.String(), tt.out)

			assert.Equal(t, match, tt.result)
		})
	}
}
