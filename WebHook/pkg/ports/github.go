package ports

type GithubServices interface {
	ValidatedHash(signature string, payload []byte) error
	DecodeMessage(event string, body []byte) string
}
