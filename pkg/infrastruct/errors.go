package infrastruct

import "net/http"

type CustomError struct {
	msg  string
	Code int
}

func NewError(msg string, code int) *CustomError {
	return &CustomError{
		msg:  msg,
		Code: code,
	}
}

func (c *CustomError) Error() string {
	return c.msg
}

var (
	ErrorInternalServerError = NewError("internal server error", http.StatusInternalServerError)
	ErrorBadRequest          = NewError("bad query input", http.StatusBadRequest)
	ErrorJWTIsBroken         = NewError("jwt spoiled", http.StatusForbidden)
	ErrorTokenIsEmpty        = NewError("token is empty", http.StatusForbidden)
	ErrorEmptyHeader         = NewError("empty auth header", http.StatusForbidden)
	ErrorInvalidHeader       = NewError("invalid auth header", http.StatusForbidden)
	ErrorBadNameDeck         = NewError("invalid deck name entry. The number of characters must be greater than 3", http.StatusBadRequest)
	ErrorDoubleNameDeck      = NewError("this deck name is already in use", http.StatusBadRequest)
	ErrorDoubleTitleCard     = NewError("deck card name already in use", http.StatusBadRequest)
	ErrorNotValidatedUser    = NewError("nickname or password not validated", http.StatusBadRequest)
	ErrorNicknameIsExist     = NewError("nickname already registered", http.StatusConflict)
	ErrorAuthIsIncorrect     = NewError("wrong login or password", http.StatusBadRequest)
	ErrorPermissionDenied    = NewError("you don't have enough rights", http.StatusForbidden)
	ErrorNoSuchDecks         = NewError("no such decks of cards", http.StatusBadRequest)
	ErrorNoSuchCard          = NewError("no such cards", http.StatusBadRequest)
	ErrorPermDeniedCrCard    = NewError("you are not the creator of this card", http.StatusForbidden)
	ErrorWrongFileType       = NewError("wrong file type", http.StatusBadRequest)
	ErrorWrongFileSize       = NewError("wrong file size", http.StatusBadRequest)
	ErrorFileHasBeenDeleted  = NewError("this file has been deleted", http.StatusBadRequest)
)
