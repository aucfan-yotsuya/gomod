package redis

var (
	err error
	r   *Redis
	ep  *Endpoint
)

func New() *Endpoint {
	ep = new(Endpoint)
	return ep
}
