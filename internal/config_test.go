package python

import (
	"strings"
	"testing"
)

func TestParseConfigDisallowUnknownFields(t *testing.T) {
	_, err := parseConfig([]byte(`{"emit_sync_querier":true,"db_typ":"jsonb"}`))
	if err == nil {
		t.Fatal("expected unknown field error, got nil")
	}
	if !strings.Contains(err.Error(), "invalid plugin options") {
		t.Fatalf("expected error to reference plugin options, got: %v", err)
	}
	if !strings.Contains(err.Error(), `unknown field "db_typ"`) {
		t.Fatalf("expected unknown field in error, got: %v", err)
	}
}

func TestParseConfigValid(t *testing.T) {
	conf, err := parseConfig([]byte(`{"emit_sync_querier":true,"overrides":[{"db_type":"jsonb","py_type":"str"}]}`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !conf.EmitSyncQuerier {
		t.Fatal("expected emit_sync_querier to be true")
	}
	if len(conf.Overrides) != 1 || conf.Overrides[0].DbType != "jsonb" || conf.Overrides[0].PyType != "str" {
		t.Fatal("unexpected parsed overrides")
	}
}
