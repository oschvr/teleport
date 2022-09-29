package auditevents

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

const emitted string = `
package desktopservice

func (s *WindowsService) onSessionEnd(ctx context.Context, emitter events.Emitter, id *tlsca.Identity, startedAt time.Time, recorded bool, windowsUser, sessionID string, desktop types.WindowsDesktop) {
  userMetadata := id.GetUserMetadata()
  userMetadata.Login = windowsUser

  event := &events.WindowsDesktopSessionEnd{
    Metadata: events.Metadata{
      Type:        libevents.WindowsDesktopSessionEndEvent,
      Code:        libevents.DesktopSessionEndCode,
      ClusterName: s.clusterName,
    },
    UserMetadata: userMetadata,
    SessionMetadata: events.SessionMetadata{
      SessionID: sessionID,
      WithMFA:   id.MFAVerified,
    },
    WindowsDesktopService: s.cfg.Heartbeat.HostUUID,
    DesktopAddr:           desktop.GetAddr(),
    Domain:                desktop.GetDomain(),
    WindowsUser:           windowsUser,
    DesktopLabels:         desktop.GetAllLabels(),
    StartTime:             startedAt,
    EndTime:               s.cfg.Clock.Now().UTC().Round(time.Millisecond),
    DesktopName:           desktop.GetName(),
    Recorded:              recorded,

    // There can only be 1 participant, desktop sessions are not join-able.
    Participants: []string{userMetadata.User},
  }
  s.emit(ctx, emitter, event)
}

func (s *WindowsService) onClipboardReceive(ctx context.Context, emitter events.Emitter, id *tlsca.Identity, sessionID string, desktopAddr string, length int32) {
  event := &events.DesktopClipboardReceive{
    Metadata: events.Metadata{
      Type:        libevents.DesktopClipboardReceiveEvent,
      Code:        libevents.DesktopClipboardReceiveCode,
      ClusterName: s.clusterName,
      Time:        s.cfg.Clock.Now().UTC(),
    },
    UserMetadata: id.GetUserMetadata(),
    SessionMetadata: events.SessionMetadata{
      SessionID: sessionID,
      WithMFA:   id.MFAVerified,
    },
    ConnectionMetadata: events.ConnectionMetadata{
      LocalAddr:  id.ClientIP,
      RemoteAddr: desktopAddr,
      Protocol:   libevents.EventProtocolTDP,
    },
    DesktopAddr: desktopAddr,
    Length:      length,
  }
  s.emit(ctx, emitter, event)
}
`

// A Metadata field that's a composite literal but has unexpected field names
const badMetadata string = `

package mypackage

func doStuffWithAnotherMetadata() otherpkg.Data{
  return otherpkg.Data{
    Metadata: otherpkg.Metadata{
      Type: types.MyCoolMetadataType,
      FavoriteNumber: 15,  
      AnimalName: "Dog",
    },
  }
}
`

const eventData string = `
  package events

  const(
    // WindowsDesktopSessionStartEvent is emitted when a user attempts
    // to connect to a desktop.
    WindowsDesktopSessionStartEvent = "windows.desktop.session.start"
    // WindowsDesktopSessionEndEvent is emitted when a user disconnects
    // from a desktop.
    WindowsDesktopSessionEndEvent = "windows.desktop.session.end"

    // CertificateCreateEvent is emitted when a certificate is issued.
    CertificateCreateEvent = "cert.create"

    // RenewableCertificateGenerationMismatchEvent is emitted when a renewable
    // certificate's generation counter is invalid.
    RenewableCertificateGenerationMismatchEvent = "cert.generation_mismatch"
  )
`

func TestGetEmittedAuditEventsFromFile(t *testing.T) {
	cases := []struct {
		desc       string
		fileString string
		expected   []string
	}{
		{
			desc:       "Happy path",
			fileString: emitted,
			expected: []string{
				"WindowsDesktopSessionEndEvent",
				"DesktopClipboardReceiveEvent",
			},
		},
		{
			desc:       "Metadata composite literal with unexpected fields",
			fileString: badMetadata,
			expected:   []string{},
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "myfile.go", c.fileString, parser.ParseComments)
			if err != nil {
				t.Fatal("unexpected error parsing the test fixture: ", err)
			}
			a := getEmittedAuditEventsFromFile(f)
			if !reflect.DeepEqual(c.expected, a) {
				t.Fatalf("expected %v but got %v", c.expected, a)
			}
		})
	}
}

func TestGetDataForEventTypes(t *testing.T) {
	cases := []struct {
		desc              string
		auditEventTypeMap auditEventTypeMap
		expected          []EventData
	}{
		{
			desc: "all event types found",
			auditEventTypeMap: auditEventTypeMap{
				"WindowsDesktopSessionStartEvent": struct{}{},
				"WindowsDesktopSessionEndEvent":   struct{}{},
			},
			expected: []EventData{
				EventData{
					Name:    "windows.desktop.session.start",
					Comment: "`windows.desktop.session.start` is emitted when a user attempts to connect to a desktop. ",
				},
				EventData{
					Name:    "windows.desktop.session.end",
					Comment: "`windows.desktop.session.end` is emitted when a user disconnects from a desktop. ",
				},
			},
		},
		{
			desc: "no event types found",
			auditEventTypeMap: auditEventTypeMap{
				"MyFakeEvent":      struct{}{},
				"AnotherFakeEvent": struct{}{},
			},
			expected: []EventData{},
		},
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "myfile.go", eventData, parser.ParseComments)
	if err != nil {
		t.Fatal("unexpected error parsing the test fixture: ", err)
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			if ed := getDataForEventTypes(f, c.auditEventTypeMap); !reflect.DeepEqual(ed, c.expected) {
				t.Errorf("expected %v but got %v", c.expected, ed)
			}

		})
	}
}

func TestGenerateAuditEventsTable(t *testing.T) {
	expected := "|Event Type|Description|\n" +
		"|---|---|\n" +
		"|windows.desktop.session.end|`windows.desktop.session.end` is emitted when a user disconnects from a desktop. |\n"
	fset := token.NewFileSet()
	f1, err := parser.ParseFile(fset, "myfile.go", emitted, parser.ParseComments)
	if err != nil {
		t.Fatal("unexpected error parsing the test fixture: ", err)
	}
	f2, err := parser.ParseFile(fset, "myfile.go", eventData, parser.ParseComments)
	if err != nil {
		t.Fatal("unexpected error parsing the test fixture: ", err)
	}
	var buf bytes.Buffer

	if err := GenerateAuditEventsTable(&buf, []*ast.File{f1, f2}); err != nil {
		t.Fatalf("unexpected error generating an audit events table: %v", err)
	}

	if expected != buf.String() {
		t.Fatalf("unexpected audit events table.\nWanted:\n%v\nGot:\n%v\n", expected, buf.String())
	}
}
