package jslices_test

import (
	"runtime"
	"testing"

	"github.com/a10adotapp/jslices"
	"github.com/google/go-cmp/cmp"
)

func TestConcat(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("%+v", err)
			for depth := 0; ; depth++ {
				_, file, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}

				t.Logf("====> %d: %s:%d", depth, file, line)
			}
		}
	}()

	if diff := cmp.Diff(
		[]uint32{1, 2, 3, 4, 5},
		jslices.Concat([]uint32{1, 2}, uint32(3), uint32(4), uint32(5)),
	); diff != "" {
		t.Error(diff)
	}

	if diff := cmp.Diff(
		[]uint32{1, 2, 3, 4, 5},
		jslices.Concat([]uint32{1, 2}, []uint32{3, 4, 5}),
	); diff != "" {
		t.Error(diff)
	}

	if diff := cmp.Diff(
		[]uint32{1, 2, 3, 4, 5},
		jslices.Concat([]uint32{1, 2}, uint32(3), []uint32{4, 5}),
	); diff != "" {
		t.Error(diff)
	}

	if diff := cmp.Diff(
		[]string{"tokyo", "osaka", "okinawa", "chiba", "hyogo"},
		jslices.Concat([]string{"tokyo", "osaka"}, "okinawa", "chiba", "hyogo"),
	); diff != "" {
		t.Error(diff)
	}
}
