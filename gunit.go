// Package gunit provides "testing" package hooks and convenience
// functions for writing tests in an xUnit style.
// NOTE: Only some of the exported names in this package
// are meant to be referenced by users of this package:
//
// - Fixture              // (as an embedded field on your xUnit-style struct)
// - Fixture.So(...)      // (as a convenient assertion method: So(expected, should.Equal, actual))
// - Fixture.Ok(...)      // (as a convenient boolean assertion method: Ok(condition, optionalMessage))
// - Fixture.Error(...)   // (works just like *testing.T.Error(...))
// - Fixture.Errorf(...)  // (works just like *testing.T.Errorf(...))
// - Fixture.Print(...)   // (works just like fmt.Print)
// - Fixture.Printf(...)  // (works just like fmt.Printf)
// - Fixture.Println(...) // (works just like fmt.Println)
//
// The rest are called from code generated by the command at
// github.com/smartystreets/gunit/gunit.
// Please see the README file and the examples folder for examples.
package gunit

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/gunit/gunit/generate"
)

// TT represents the functional subset from *testing.T needed by Fixture.
type TT interface {
	Fail()
	Failed() bool
	SkipNow()
}

// Fixture keeps track of test status (failed, passed, skipped) and
// handles custom logging for xUnit style tests as an embedded field.
type Fixture struct {
	t       TT
	log     *bytes.Buffer
	skipped bool
}

// NewFixture is called by generated code.
func NewFixture(t TT) *Fixture {
	return &Fixture{t: t, log: &bytes.Buffer{}}
}

// So is a convenience method for reporting assertion failure messages,
// say from the assertion functions found in github.com/smartystreets/assertions/should.
// Example: self.So(actual, should.Equal, expected)
func (self *Fixture) So(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) {
	if ok, failure := assertions.So(actual, assert, expected...); !ok {
		self.t.Fail()
		self.reportFailure(failure)
	}
}

func (self *Fixture) Ok(condition bool, messages ...string) {
	if !condition {
		if len(messages) == 0 {
			messages = append(messages, "Expected condition to be true, was false instead.")
		}
		self.t.Fail()
		self.reportFailure(strings.Join(messages, ", "))
	}
}

func (self *Fixture) Error(args ...interface{}) {
	self.t.Fail()
	self.reportFailure(fmt.Sprint(args...))
}

func (self *Fixture) Errorf(format string, args ...interface{}) {
	self.t.Fail()
	self.reportFailure(fmt.Sprintf(format, args...))
}

func (self *Fixture) reportFailure(failure string) {
	_, file, line, _ := runtime.Caller(2) // 0: reportFailure + 1: Error/Errorf/So/Ok + 2: func Test...
	self.log.WriteString(fmt.Sprintf("\tX FAILED: %s:%d\n", file, line))
	self.print(failure)
}

// Print is analogous to fmt.Print and is ideal for printing in the middle of a test case.
func (self *Fixture) Print(a ...interface{}) (n int, err error) {
	return self.print(fmt.Sprint(a...))
}

// Printf is analogous to fmt.Printf and is ideal for printing in the middle of a test case.
func (self *Fixture) Printf(format string, a ...interface{}) (n int, err error) {
	return self.print(fmt.Sprintf(format, a...))
}

// Println is analogous to fmt.Println and is ideal for printing in the middle of a test case.
func (self *Fixture) Println(a ...interface{}) (n int, err error) {
	return self.print(fmt.Sprintln(a...))
}

// print ensures that the necessary indentation is prepended to every line of each printed message.
func (self *Fixture) print(message string) (n int, err error) {
	return fmt.Fprint(self.log, "\t\t"+strings.Replace(message, "\n", "\n\t\t", -1)+"\n")
}

// Describe is called by generated code.
func (self *Fixture) Describe(test string) {
	fmt.Fprintln(self.log, "\t- "+test)
}

// Skip is called by generated code.
func (self *Fixture) Skip(message string) {
	fmt.Fprintln(self.log, "\t? "+message)
	self.skipped = true
}

// Finalize is called by generated code.
func (self *Fixture) Finalize() {
	if verbose() || self.t.Failed() {
		io.WriteString(out, strings.TrimRight(self.log.String(), "\t"))
	}

	if r := recover(); r != nil { // TODO: unit test
		panic(r)
	}

	if self.skipped {
		self.t.SkipNow()
	}
}

//////////////////////////////////////////////////////////////////////////////

// Validate is called by generated code.
func Validate(checksum string) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		exit("Unable to resolve the test file from runtime.Caller(...).\n")
	}
	current, err := generate.Checksum(filepath.Dir(file))
	if err != nil {
		exit("Could not calculate checksum of current go files. Error: %s\n", err.Error())
	}
	if checksum != current {
		exit("The checksum provided [%s] does not match the current file listing [%s]. Please re-run the `gunit` command and try again.\n", checksum, current)
	}
}

func exit(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
	os.Exit(1)
}

var verbose = testing.Verbose
var out io.Writer = os.Stdout
