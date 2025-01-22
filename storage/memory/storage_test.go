package memory

import (
	"testing"

	tl "github.com/criteo/consul-timeline/timeline"
	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {

	s := New(Config{MaxSize: 3})

	_ = s.Store(tl.Event{ID: 1})
	_ = s.Store(tl.Event{ID: 2})
	_ = s.Store(tl.Event{ID: 3})
	_ = s.Store(tl.Event{ID: 4})

	require.Equal(t,
		[]tl.Event{
			{ID: 4},
			{ID: 2},
			{ID: 3},
		},
		s.events,
	)
}
