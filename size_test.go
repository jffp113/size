package size

import "testing"

func TestParseB(t *testing.T) {
	expected := "10B"
	expectedB := 10

	actual, err := Parse(expected)
	if err != nil {
		t.Errorf("should have parsed: %v", err)
		return
	}

	if actual.Bytes() != int64(expectedB) {
		t.Errorf("expected some size: %v", err)
	}
}

func TestParseKB(t *testing.T) {
	expected := "10KB"
	expectedKB := 10

	actual, err := Parse(expected)
	if err != nil {
		t.Errorf("should have parsed: %v", err)
		return
	}

	if actual.KBytes() != float64(expectedKB) {
		t.Errorf("expected some size: %v", err)
	}

	if actual.Bytes() != int64(expectedKB*1024) {
		t.Errorf("expected some size: %v", err)
	}
}

func TestParseMB(t *testing.T) {
	expected := "10MB"
	expectedMB := 10

	actual, err := Parse(expected)
	if err != nil {
		t.Errorf("should have parsed: %v", err)
		return
	}

	if actual.MBytes() != float64(expectedMB) {
		t.Errorf("expected some size: %v", err)
	}

	if actual.KBytes() != float64(expectedMB*1024) {
		t.Errorf("expected some size: %v", err)
	}

	if actual.Bytes() != int64(expectedMB*1024*1024) {
		t.Errorf("expected some size: %v", err)
	}
}

func TestParseGB(t *testing.T) {
	expected := "10GB"
	expectedGB := 10

	actual, err := Parse(expected)
	if err != nil {
		t.Errorf("should have parsed: %v", err)
		return
	}

	if actual.GBytes() != float64(expectedGB) {
		t.Errorf("expected some size: %v", err)
	}

	if actual.MBytes() != float64(expectedGB*1024) {
		t.Errorf("expected some size: %v", err)
	}
	//
	if actual.KBytes() != float64(expectedGB*1024*1024) {
		t.Errorf("expected some size: %v", err)
	}

	if actual.Bytes() != int64(expectedGB*1024*1024*1024) {
		t.Errorf("expected some size: %v", err)
	}
}

func TestParseTB(t *testing.T) {
	expected := "10TB"
	expectedTB := 10

	actual, err := Parse(expected)
	if err != nil {
		t.Errorf("should have parsed: %v", err)
		return
	}

	if actual.TBytes() != float64(expectedTB) {
		t.Errorf("expected some size: %v", err)
	}

	if actual.GBytes() != float64(expectedTB*1024) {
		t.Errorf("expected some size: %v", err)
	}

	if actual.MBytes() != float64(expectedTB*1024*1024) {
		t.Errorf("expected some size: %v", err)
	}

	if actual.KBytes() != float64(expectedTB*1024*1024*1024) {
		t.Errorf("expected some size: %v", err)
	}

	if actual.Bytes() != int64(expectedTB*1024*1024*1024*1024) {
		t.Errorf("expected some size: %v", err)
	}
}
