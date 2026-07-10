package ports

import messainfromation "triggo/pkg/github/model/messa_infromation"

type GithubServices interface {
	ValidatedHash(signature string, payload []byte) error
	DecodeMessage(event string, body []byte) (messainfromation.MessaInformation, string)
}
