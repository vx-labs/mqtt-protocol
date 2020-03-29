package decoder

type StatRecorder interface {
	Add(float64)
}

type noopStatRecorder struct{}

func (*noopStatRecorder) Add(float64) {}
