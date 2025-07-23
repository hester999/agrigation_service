package apperr

import "errors"

var (
	ErrNotFound              = errors.New("not found")
	ErrBadRequest            = errors.New("bad request")
	ErrNameIsRequired        = errors.New("name is required")
	ErrUserIsRequired        = errors.New("user is required")
	ErrInvalidPrice          = errors.New("invalid price")
	ErrInvalidDuration       = errors.New("invalid duration")
	ErrDataIsRequired        = errors.New("data is required")
	ErrNameTooLong           = errors.New("name too long")
	ErrStartDateDurationPair = errors.New("start_date and duration must be provided together")
	ErrInvalidLimit          = errors.New("invalid limit")
	ErrInvalidOffset         = errors.New("invalid offset")
)
