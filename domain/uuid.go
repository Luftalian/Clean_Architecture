package domain

type UUID [16]byte

var (
	Nil UUID // empty UUID, all zeros
)
