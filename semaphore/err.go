package semaphore

type (
	ErrWeighted struct {
		Message string
	}
	ErrErrGroup struct {
		Message string
	}
)

func (e *ErrWeighted) Error() string {
	return e.Message
}
func (e *ErrErrGroup) Error() string {
	return e.Message
}
