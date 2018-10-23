package receive

import (
	"fmt"
	"testing"
)

func TestParseWriteRequest(t *testing.T) {
	timeseries, err := parseWriteRequest(b)
	if err != nil {
		t.Fatal(err)
	}

	for _, ts := range timeseries {
		lset, err := parseLabels(ts.LabelSetBytes)
		if err != nil {
			t.Fatal(err)
		}

		for _, s := range ts.Samples {
			fmt.Printf("%s %f %d\n", lset.String(), s.Value, s.Timestamp)
		}
	}
}

var (
	b = []byte{ // WriteRequest
		0xa,                                            // 10 <- start of time-series
		0x60,                                           // 96 <- bytes from here until end of time-series?
		0xa,                                            // 10 <- label start?
		0x13,                                           // 19 <- bytes from here until end of label
		0xa,                                            // 10 <- label key next
		0x8,                                            // 8 <- size of label key
		0x5f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, // __name__
		0x12,                                     // 18 <- label value next
		0x7,                                      // 7 <- size of label value
		0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, // "version"
		0xa,                                            // 10
		0x1a,                                           // 26 <- bytes from here until end of label
		0xa,                                            // 10 <- label key next
		0x8,                                            // 8 <- size of label key
		0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, // instance
		0x12,                                                                               // 18 <- label value next
		0xe,                                                                                // 14 <- size of label value
		0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, // localhost:8080
		0xa,              // 10
		0x11,             // 17 <- bytes from here until end of label
		0xa,              // 10 <- label key next
		0x3,              // 3 <- size of label key
		0x6a, 0x6f, 0x62, // job
		0x12,                                                       // 18 <- label value next
		0xa,                                                        // 10 <- size of label value
		0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, // prometheus
		0xa,                                      // 10
		0x11,                                     // 17 <- bytes from here until end of label
		0xa,                                      // 10 <- label key next
		0x7,                                      // 7 <- size of label key
		0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, // version
		0x12,                               // 18 <- label value next
		0x6,                                // 6 <- size of label value
		0x76, 0x30, 0x2e, 0x31, 0x2e, 0x30, // v0.1.0
		0x12, // 18 <- sample next
		0x7,  // 7  <- bytes from here until end of sample
		0x10, // 16 <- time-stamp next
		0xa7, //
		0xa0, //
		0xdf, //
		0xde, //
		0xd6, //
		0x2c, //

		// --- end of time-series ---

		0xa,                                            // 10 <- start of time-series
		0x51,                                           // 81 <- bytes from here until end of time-series?
		0xa,                                            // 10
		0xe,                                            // 14
		0xa,                                            // 10
		0x8,                                            // 8 <- size
		0x5f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, // __name__
		0x12,
		0x2,
		0x75,
		0x70,
		0xa,
		0x1a,
		0xa,
		0x8,
		0x69,
		0x6e,
		0x73,
		0x74,
		0x61,
		0x6e,
		0x63,
		0x65,
		0x12,
		0xe,
		0x6c,
		0x6f,
		0x63,
		0x61,
		0x6c,
		0x68,
		0x6f,
		0x73,
		0x74,
		0x3a,
		0x38,
		0x30,
		0x38,
		0x30,
		0xa,
		0x11,
		0xa,
		0x3,
		0x6a,
		0x6f,
		0x62,
		0x12,
		0xa,
		0x70,
		0x72,
		0x6f,
		0x6d,
		0x65,
		0x74,
		0x68,
		0x65,
		0x75,
		0x73,
		0x12,
		0x10,
		0x9,
		0x0,
		0x0,
		0x0,
		0x0,
		0x0,
		0x0,
		0xf0,
		0x3f,
		0x10,
		0xa7,
		0xa0,
		0xdf,
		0xde,
		0xd6,
		0x2c,
		0xa,
		0x66,
		0xa,
		0x23,
		0xa,
		0x8,
		0x5f,
		0x5f,
		0x6e,
		0x61,
		0x6d,
		0x65,
		0x5f,
		0x5f,
		0x12,
		0x17,
		0x73,
		0x63,
		0x72,
		0x61,
		0x70,
		0x65,
		0x5f,
		0x64,
		0x75,
		0x72,
		0x61,
		0x74,
		0x69,
		0x6f,
		0x6e,
		0x5f,
		0x73,
		0x65,
		0x63,
		0x6f,
		0x6e,
		0x64,
		0x73,
		0xa,
		0x1a,
		0xa,
		0x8,
		0x69,
		0x6e,
		0x73,
		0x74,
		0x61,
		0x6e,
		0x63,
		0x65,
		0x12,
		0xe,
		0x6c,
		0x6f,
		0x63,
		0x61,
		0x6c,
		0x68,
		0x6f,
		0x73,
		0x74,
		0x3a,
		0x38,
		0x30,
		0x38,
		0x30,
		0xa,
		0x11,
		0xa,
		0x3,
		0x6a,
		0x6f,
		0x62,
		0x12,
		0xa,
		0x70,
		0x72,
		0x6f,
		0x6d,
		0x65,
		0x74,
		0x68,
		0x65,
		0x75,
		0x73,
		0x12,
		0x10,
		0x9,
		0x80,
		0xae,
		0x5f,
		0xab,
		0x62,
		0xfc,
		0x92,
		0x3f,
		0x10,
		0xa7,
		0xa0,
		0xdf,
		0xde,
		0xd6,
		0x2c,
		0xa,
		0x65,
		0xa,
		0x22,
		0xa,
		0x8,
		0x5f,
		0x5f,
		0x6e,
		0x61,
		0x6d,
		0x65,
		0x5f,
		0x5f,
		0x12,
		0x16,
		0x73,
		0x63,
		0x72,
		0x61,
		0x70,
		0x65,
		0x5f,
		0x73,
		0x61,
		0x6d,
		0x70,
		0x6c,
		0x65,
		0x73,
		0x5f,
		0x73,
		0x63,
		0x72,
		0x61,
		0x70,
		0x65,
		0x64,
		0xa,
		0x1a,
		0xa,
		0x8,
		0x69,
		0x6e,
		0x73,
		0x74,
		0x61,
		0x6e,
		0x63,
		0x65,
		0x12,
		0xe,
		0x6c,
		0x6f,
		0x63,
		0x61,
		0x6c,
		0x68,
		0x6f,
		0x73,
		0x74,
		0x3a,
		0x38,
		0x30,
		0x38,
		0x30,
		0xa,
		0x11,
		0xa,
		0x3,
		0x6a,
		0x6f,
		0x62,
		0x12,
		0xa,
		0x70,
		0x72,
		0x6f,
		0x6d,
		0x65,
		0x74,
		0x68,
		0x65,
		0x75,
		0x73,
		0x12,
		0x10,
		0x9,
		0x0,
		0x0,
		0x0,
		0x0,
		0x0,
		0x0,
		0xf0,
		0x3f,
		0x10,
		0xa7,
		0xa0,
		0xdf,
		0xde,
		0xd6,
		0x2c,
		0xa,
		0x74,
		0xa,
		0x31,
		0xa,
		0x8,
		0x5f,
		0x5f,
		0x6e,
		0x61,
		0x6d,
		0x65,
		0x5f,
		0x5f,
		0x12,
		0x25,
		0x73,
		0x63,
		0x72,
		0x61,
		0x70,
		0x65,
		0x5f,
		0x73,
		0x61,
		0x6d,
		0x70,
		0x6c,
		0x65,
		0x73,
		0x5f,
		0x70,
		0x6f,
		0x73,
		0x74,
		0x5f,
		0x6d,
		0x65,
		0x74,
		0x72,
		0x69,
		0x63,
		0x5f,
		0x72,
		0x65,
		0x6c,
		0x61,
		0x62,
		0x65,
		0x6c,
		0x69,
		0x6e,
		0x67,
		0xa,
		0x1a,
		0xa,
		0x8,
		0x69,
		0x6e,
		0x73,
		0x74,
		0x61,
		0x6e,
		0x63,
		0x65,
		0x12,
		0xe,
		0x6c,
		0x6f,
		0x63,
		0x61,
		0x6c,
		0x68,
		0x6f,
		0x73,
		0x74,
		0x3a,
		0x38,
		0x30,
		0x38,
		0x30,
		0xa,
		0x11,
		0xa,
		0x3,
		0x6a,
		0x6f,
		0x62,
		0x12,
		0xa,
		0x70,
		0x72,
		0x6f,
		0x6d,
		0x65,
		0x74,
		0x68,
		0x65,
		0x75,
		0x73,
		0x12,
		0x10,
		0x9,
		0x0,
		0x0,
		0x0,
		0x0,
		0x0,
		0x0,
		0xf0,
		0x3f,
		0x10,
		0xa7,
		0xa0,
		0xdf,
		0xde,
		0xd6,
		0x2c,
	}
)
