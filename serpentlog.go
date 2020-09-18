package cobra

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sort"
	"time"
)

type LogFields struct {
	Fields *log.Fields
	ok     bool
}

func WithField(k string, v interface{}) *LogFields {
	return NewLogFields(k, v)
}

func (f *LogFields) WithField(k string, v interface{}) *LogFields {
	nf := NewLogFields(k, v)
	for fk, fv := range *f.Fields {
		nf.Add(fk, fv)
	}
	nf.ok = f.ok
	return nf
}

func WithFields(f *log.Fields) *LogFields {
	return &LogFields{f, true}
}

func NewLogFields(k string, v interface{}) *LogFields {
	m := &LogFields{&log.Fields{}, true}
	m.Add(k, v)
	return m
}

func (f *LogFields) Tag(t string) *LogFields {
	f.ok = false
	reqTags := GetStringMap("serpenttags")
	switch {
	case GetBool("logalltags"):
		f.ok = true
	case len(reqTags) > 0:
		_, found := reqTags[t]
		f.ok = found
	}
	if f.ok {
		f.Add("serpenttag", t)
	}
	return f
}

func Tag(t string) *LogFields {
	f := &LogFields{&log.Fields{}, true}
	return f.Tag(t)
}

func (f *LogFields) Copy() *LogFields {
	nf := &LogFields{&log.Fields{}, true}
	for fk, fv := range *f.Fields {
		nf.Add(fk, fv)
	}
	nf.ok = f.ok
	return nf
}

func (f *LogFields) Add(k string, v interface{}) *LogFields {
	(*f.Fields)[k] = v
	return f
}

func (f *LogFields) Type(v interface{}) *LogFields {
	(*f.Fields)["type"] = fmt.Sprintf("%T", v)
	return f
}

// Strunc truncates string values greater than 20 character and treats
// all other arguments as Add.
func (f *LogFields) Strunc(k string, v interface{}) *LogFields {
	x := v
	switch v.(type) {
	case string:
		if len(v.(string)) > 20 {
			x = fmt.Sprintf("%q", v.(string)[:20]+"...")
		} else {
			x = fmt.Sprintf("%q", v.(string))
		}
	}
	(*f.Fields)[k] = x
	return f
}

func Out(args ...interface{}) {
	log.Info(args...)
}

func Outf(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Outln(args ...interface{}) {
	log.Infoln(args...)
}

func Log(args ...interface{}) {
	log.Debug(args...)
}

func Logf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Logln(args ...interface{}) {
	log.Debugln(args...)
}

func OutV(args ...interface{}) {
	if GetBool("verbose") {
		log.Info(args...)
	}
}

func OutfV(format string, args ...interface{}) {
	if GetBool("verbose") {
		log.Infof(format, args...)
	}
}

func OutlnV(args ...interface{}) {
	if GetBool("verbose") {
		log.Infoln(args...)
	}
}

func LogV(args ...interface{}) {
	if GetBool("verbose") {
		log.Debug(args...)
	}
}

func LogfV(format string, args ...interface{}) {
	if GetBool("verbose") {
		log.Debugf(format, args...)
	}
}

func LoglnV(args ...interface{}) {
	if GetBool("verbose") {
		log.Debugln(args...)
	}
}

// LogFields

func (f *LogFields) Out(args ...interface{}) {
	if f.ok {
		log.WithFields(*f.Fields).Info(args...)
	}
}

func (f *LogFields) Outf(format string, args ...interface{}) {
	if f.ok {
		log.WithFields(*f.Fields).Infof(format, args...)
	}
}

func (f *LogFields) Outln(args ...interface{}) {
	if f.ok {
		log.WithFields(*f.Fields).Infoln(args...)
	}
}

func (f *LogFields) Log(args ...interface{}) {
	if f.ok {
		log.WithFields(*f.Fields).Debug(args...)
	}
}

func (f *LogFields) Logf(format string, args ...interface{}) {
	if f.ok {
		log.WithFields(*f.Fields).Debugf(format, args...)
	}
}

func (f *LogFields) Logln(args ...interface{}) {
	if f.ok {
		log.WithFields(*f.Fields).Debugln(args...)
	}
}

func (f *LogFields) OutV(args ...interface{}) {
	if f.ok && GetBool("verbose") {
		log.WithFields(*f.Fields).Info(args...)
	}
}

func (f *LogFields) OutfV(format string, args ...interface{}) {
	if f.ok && GetBool("verbose") {
		log.WithFields(*f.Fields).Infof(format, args...)
	}
}

func (f *LogFields) OutlnV(args ...interface{}) {
	if f.ok && GetBool("verbose") {
		log.WithFields(*f.Fields).Infoln(args...)
	}
}

func (f *LogFields) LogV(args ...interface{}) {
	if f.ok && GetBool("verbose") {
		log.WithFields(*f.Fields).Debug(args...)
	}
}

func (f *LogFields) LogfV(format string, args ...interface{}) {
	if f.ok && GetBool("verbose") {
		log.WithFields(*f.Fields).Debugf(format, args...)
	}
}

func (f *LogFields) LoglnV(args ...interface{}) {
	if f.ok && GetBool("verbose") {
		log.WithFields(*f.Fields).Debugln(args...)
	}
}

// Log Formatting -------------------------------------------------------------

var (
	baseTimestamp time.Time
	emptyFieldMap log.FieldMap
)

type SerpentFormatter struct {
	Debug            bool // if false, debug logs are ignored.
	HideTags         bool
	DisableSorting   bool
	QuoteEmptyFields bool
	// sync.Once
}

func (f *SerpentFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	prefixFieldClashes(entry.Data)

	switch {
	case entry.Level == log.DebugLevel && !f.Debug:
	default:
		t, tagged := entry.Data["serpenttag"]
		if tagged {
			if !f.HideTags {
				fmt.Fprintf(b, "(%s) ", t)
			}
			delete(entry.Data, "serpenttag")
		}

		keys := make([]string, 0, len(entry.Data))
		for k := range entry.Data {
			keys = append(keys, k)
		}
		if !f.DisableSorting {
			sort.Strings(keys)
		}

		b.WriteString(entry.Message)
		// fmt.Printf("keys: %v\n", keys)
		if len(keys) > 0 {
			b.WriteString(" [")
			for i, k := range keys {
				v := entry.Data[k]
				fmt.Fprintf(b, "%s=%+v", k, v)
				if i < len(keys)-1 {
					b.WriteString(", ")
					// f.appendKeyValue(b, k, entry.Data[k])
				}
			}
			b.WriteString("]")
		}
		b.WriteByte('\n')
	}
	return b.Bytes(), nil
}

func (f *SerpentFormatter) needsQuoting(text string) bool {
	if f.QuoteEmptyFields && len(text) == 0 {
		return true
	}
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
			return true
		}
	}
	return false
}

func (f *SerpentFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	f.appendValue(b, value)
}

func (f *SerpentFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	if !f.needsQuoting(stringVal) {
		b.WriteString(stringVal)
	} else {
		b.WriteString(fmt.Sprintf("%q", stringVal))
	}
}

// This is to not silently overwrite `time`, `msg` and `level` fields when
// dumping it. If this code wasn't there doing:
//
//  logrus.WithField("level", 1).Info("hello")
//
// Would just silently drop the user provided level. Instead with this code
// it'll logged as:
//
//  {"level": "info", "fields.level": 1, "msg": "hello", "time": "..."}
//
// It's not exported because it's still using Data in an opinionated way. It's to
// avoid code duplication between the two default formatters.
func prefixFieldClashes(data log.Fields) {
	if t, ok := data["time"]; ok {
		data["fields.time"] = t
	}

	if m, ok := data["msg"]; ok {
		data["fields.msg"] = m
	}

	if l, ok := data["level"]; ok {
		data["fields.level"] = l
	}
}
