package xerrors

import "github.com/lib/pq"

const uniquenessViolation = pq.ErrorCode("23505")

func IsPqUniquenessViolation(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		if pgerr.Code == uniquenessViolation {
			return true
		}
	}
	return false
}
