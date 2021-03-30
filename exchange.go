package polis

type Opts struct {
	MaxThreadDepth        uint
	MaxConnCurrentStreams uint

	PersistThreads    bool
	AllowMessageEdits bool
}

type Exchange struct{}
