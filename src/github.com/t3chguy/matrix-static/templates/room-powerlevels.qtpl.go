// This file is automatically generated by qtc from "room-powerlevels.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:1
package templates

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:1
import "github.com/t3chguy/matrix-static/mxclient"

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:5
type RoomPowerLevelsPage struct {
	RoomInfo    mxclient.RoomInfo
	PowerLevels mxclient.PowerLevels
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:12
func streamprintPLRow(qw422016 *qt422016.Writer, name string, pl mxclient.PowerLevel) {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:12
	qw422016.N().S(`
    <tr>
        <td>`)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:14
	qw422016.E().S(name)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:14
	qw422016.N().S(`</td>
        <td>`)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:15
	qw422016.N().D(pl.Int())
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:15
	qw422016.N().S(`</td>
    </tr>
`)
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
func writeprintPLRow(qq422016 qtio422016.Writer, name string, pl mxclient.PowerLevel) {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
	streamprintPLRow(qw422016, name, pl)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
func printPLRow(name string, pl mxclient.PowerLevel) string {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
	writeprintPLRow(qb422016, name, pl)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
	qs422016 := string(qb422016.B)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
	return qs422016
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:17
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:21
func (p *RoomPowerLevelsPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:21
	qw422016.N().S(`
    Public Room Servers - `)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:22
	qw422016.E().S(p.RoomInfo.Name)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:22
	qw422016.N().S(` - Riot Static
`)
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
func (p *RoomPowerLevelsPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
	p.StreamTitle(qw422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
func (p *RoomPowerLevelsPage) Title() string {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
	p.WriteTitle(qb422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
	qs422016 := string(qb422016.B)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
	return qs422016
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:23
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:25
func (p *RoomPowerLevelsPage) StreamHead(qw422016 *qt422016.Writer) {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:25
	qw422016.N().S(`
`)
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
func (p *RoomPowerLevelsPage) WriteHead(qq422016 qtio422016.Writer) {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
	p.StreamHead(qw422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
func (p *RoomPowerLevelsPage) Head() string {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
	p.WriteHead(qb422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
	qs422016 := string(qb422016.B)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
	return qs422016
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:26
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:28
func (p *RoomPowerLevelsPage) StreamHeader(qw422016 *qt422016.Writer) {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:28
	qw422016.N().S(`
    `)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:29
	StreamPrintRoomHeader(qw422016, p.RoomInfo)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:29
	qw422016.N().S(`
`)
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
func (p *RoomPowerLevelsPage) WriteHeader(qq422016 qtio422016.Writer) {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
	p.StreamHeader(qw422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
	qt422016.ReleaseWriter(qw422016)
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
func (p *RoomPowerLevelsPage) Header() string {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
	qb422016 := qt422016.AcquireByteBuffer()
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
	p.WriteHeader(qb422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
	qs422016 := string(qb422016.B)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
	qt422016.ReleaseByteBuffer(qb422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
	return qs422016
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:30
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:33
func (p *RoomPowerLevelsPage) StreamBody(qw422016 *qt422016.Writer) {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:33
	qw422016.N().S(`Room Power Level Requirements<table>`)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:37
	streamprintPLRow(qw422016, "Ban", p.PowerLevels.Ban)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:38
	streamprintPLRow(qw422016, "Kick", p.PowerLevels.Kick)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:39
	streamprintPLRow(qw422016, "Redact", p.PowerLevels.Redact)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:40
	streamprintPLRow(qw422016, "User Default", p.PowerLevels.UsersDefault)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:41
	streamprintPLRow(qw422016, "State Default", p.PowerLevels.StateDefault)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:42
	streamprintPLRow(qw422016, "Events Default", p.PowerLevels.EventsDefault)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:42
	qw422016.N().S(`<tr><td>Events</td><td><table>`)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:48
	for Type, pl := range p.PowerLevels.Events {
		//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:49
		streamprintPLRow(qw422016, Type, pl)
		//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:50
	}
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:50
	qw422016.N().S(`</table></td></tr><tr><td>Users (hides PL==UsersDefault)</td><td><table>`)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:59
	for mxid, pl := range p.PowerLevels.Users {
		//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:60
		if pl != p.PowerLevels.UsersDefault {
			//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:61
			streamprintPLRow(qw422016, mxid, pl)
			//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:62
		}
		//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:63
	}
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:63
	qw422016.N().S(`</table></td></tr></table><a href="./">Back to Room List</a>`)
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
func (p *RoomPowerLevelsPage) WriteBody(qq422016 qtio422016.Writer) {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
	p.StreamBody(qw422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
	qt422016.ReleaseWriter(qw422016)
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
}

//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
func (p *RoomPowerLevelsPage) Body() string {
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
	qb422016 := qt422016.AcquireByteBuffer()
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
	p.WriteBody(qb422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
	qs422016 := string(qb422016.B)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
	qt422016.ReleaseByteBuffer(qb422016)
	//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
	return qs422016
//line src\github.com\t3chguy\matrix-static\templates\room-powerlevels.qtpl:72
}
