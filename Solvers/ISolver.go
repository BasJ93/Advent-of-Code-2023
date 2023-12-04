package Solvers

type ISolver interface {
	ChallengeOne(input IStringScanner) string
	ChallengeTwo(input IStringScanner) string
}
