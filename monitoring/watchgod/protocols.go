package watchgod

type Protocol int64

const (
	TCP Protocol = iota
	ICMP
	HTTP
	UDP
)

func (p Protocol) String() string {
	switch p {
	case TCP:
		return "TCP"
	case HTTP:
		return "http"
	case ICMP:
		return "icmp"
	case UDP:
		return "udp"
	}
	return "unknown"
}
