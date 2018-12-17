package compatible

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/uber/prototool/internal/location"
)

type methods map[string]*method

var _ descriptorProtoGroup = (methods)(nil)

func (ms methods) Items() map[string]descriptorProto {
	items := make(map[string]descriptorProto)
	for i, m := range ms {
		items[i] = m
	}
	return items
}

// method represents a *descriptor.MethodDescriptorProto.
type method struct {
	path            location.Path
	name            string
	input           string
	output          string
	clientStreaming bool
	serverStreaming bool
}

var _ descriptorProto = (*method)(nil)

func (m *method) Name() string        { return m.name }
func (m *method) Path() location.Path { return m.path }
func (m *method) Type() string        { return fmt.Sprintf("Method %q", m.name) }

func newMethod(md *descriptor.MethodDescriptorProto, p location.Path) *method {
	return &method{
		path:            p,
		name:            md.GetName(),
		input:           strings.TrimPrefix(md.GetInputType(), "."),
		output:          strings.TrimPrefix(md.GetOutputType(), "."),
		clientStreaming: md.GetClientStreaming(),
		serverStreaming: md.GetServerStreaming(),
	}
}

// checkMethods verifies that,
//  - None of the methods were removed.
//  - None of the methods' request types were updated.
//  - None of the methods' response types were updated.
//  - None of the methods' client streaming capabilities were updated.
//  - None of the methods' server streaming capabilities were updated.
func (c *fileChecker) checkMethods(original, updated methods) {
	c.checkRemovedItems(original, updated, location.Name)
	for i, um := range updated {
		if om, ok := original[i]; ok {
			c.checkMethod(om, um)
		}
	}
}

func (c *fileChecker) checkMethod(original, updated *method) {
	c.checkUpdatedAttribute(
		original,
		Wire,
		"request type",
		original.input,
		updated.input,
		location.MethodRequest,
	)
	c.checkUpdatedAttribute(
		original,
		Wire,
		"response type",
		original.output,
		updated.output,
		location.MethodResponse,
	)
	c.checkUpdatedAttribute(
		original,
		getStreamSeverity(original.clientStreaming),
		"client streaming",
		strconv.FormatBool(original.clientStreaming),
		strconv.FormatBool(updated.clientStreaming),
		location.MethodRequest,
	)
	c.checkUpdatedAttribute(
		original,
		getStreamSeverity(original.serverStreaming),
		"server streaming",
		strconv.FormatBool(original.serverStreaming),
		strconv.FormatBool(updated.serverStreaming),
		location.MethodResponse,
	)
}

// getStreamSeverity determines the severity of a client/server stream
// update. A method without client/server streaming can update to
// enable either (or both) and continue to be wire-compatible.
// However, an update in the reverse direction, i.e. from streaming to
// not streaming is NOT wire-compatible.
func getStreamSeverity(original bool) Severity {
	if original {
		return Wire
	}
	return Source
}
