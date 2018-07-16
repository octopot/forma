package draft

type Signer interface {
	Token() string
}

type Encoder interface {
	Encode(string) string
	Decode(string) string
}
