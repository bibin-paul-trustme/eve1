package types

import (
  "testing"
  "errors"
  "fmt"
  log "github.com/sirupsen/logrus"
)
type TestTypesMatrixEntry struct {
  safename  string
  url       string
  sha       string
  filename  string
  err       error
  ts        TriState
  value     string
}

func TestParseTriState(t *testing.T) {
    log.Infof("TestLookupIoBundle: START\n")
  testMatrix := []TestTypesMatrixEntry{
        {err: nil, ts: TS_NONE, value: "none"},

        {err: nil, ts: TS_ENABLED, value: "enable"},
    {err: nil, ts: TS_ENABLED, value: "enabled"},
    {err: nil, ts: TS_ENABLED, value: "on"},

    {err: nil, ts: TS_DISABLED, value: "disabled"},
    {err: nil, ts: TS_DISABLED, value: "disable"},
    {err: nil, ts: TS_DISABLED, value: "off"},

    {err: errors.New(fmt.Sprintf("Bad value: bad-value")), ts: TS_NONE, value: "bad-value"},
    }

  for index := range testMatrix {
        entry := &testMatrix[index]
        ts, err := ParseTriState(entry.value)
    if ts != entry.ts {
            t.Errorf("Test Entry Index %d Failed: Expected TS: %s, Actual TS: %s\n",
                index, entry.ts, ts)
        } else if err == nil && entry.err == nil {

    } else if err.Error() != entry.err.Error() {
      t.Errorf("Test Entry Index %d Failed: Expected Error: %e, Actual Error: %e\n",
                index, entry.err, err)
    } else {}
  }
    log.Infof("TestLookupIoBundle: DONE\n")
}
