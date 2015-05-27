package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"github.com/smartystreets/gunit/gunit/parse"
)

// TestFile generates complete source code for a _test.go file from the provided fixtures.
func TestFile(packageName string, fixtures []*parse.Fixture, checksum string, assertions map[string]map[int]string) ([]byte, error) {
	buffer := bytes.NewBufferString(fmt.Sprintf(header, packageName))
	buffer.WriteString("\n///////////////////////////////////////////////////////////////////////////////\n\n")
	for _, fixture := range fixtures {
		function, err := TestFunction(fixture)
		if err != nil {
			return nil, err
		}
		buffer.Write(function)
		buffer.WriteString("\n\n///////////////////////////////////////////////////////////////////////////////\n\n")
	}
	buffer.WriteString(fmt.Sprintf(footer, assertions, checksum))
	return format.Source(buffer.Bytes())
}

// TestFunction generates a test function based solely on whether the fixture has test cases that are Skipped or not.
func TestFunction(fixture *parse.Fixture) ([]byte, error) {
	if fixture.Skipped {
		return executeTemplate(skippedFixtureTemplate, fixture)
	}
	return executeTemplate(testFixtureTemplate, fixture)
}

func executeTemplate(template *template.Template, data interface{}) ([]byte, error) {
	writer := &bytes.Buffer{}
	template.Execute(writer, data)
	return format.Source(writer.Bytes())
}

var skippedFixtureTemplate = template.Must(template.
	New("testFunction").Funcs(map[string]interface{}{"sentence": toSentence}).
	Parse(rawSkippedFixture))

var testFixtureTemplate = template.Must(template.
	New("testFunction").Funcs(map[string]interface{}{"sentence": toSentence}).
	Parse(rawTestFunction))
