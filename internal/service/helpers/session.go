package helpers

import (
	"net/http"

	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

func GetNonce(address string, r *http.Request) (*data.Nonce, *jsonapi.ErrorObject, error) {
	db := DB(r)
	nonce, err := db.Nonce().FilterByAddress(address).Get()
	if err != nil {
		return nil, problems.InternalError(), errors.New("failed to query db")
	}

	if nonce == nil {
		return nil, problems.BadRequest(errors.New("nonce does not exist"))[0], errors.New("nonce not found")
	}

	err = db.Nonce().FilterByAddress(address).Delete()
	if err != nil {
		return nil, problems.InternalError(), errors.New("failed to query db")
	}
	return nonce, nil, nil
}
func ValidateNonce(address string, signature string, r *http.Request) (*jsonapi.ErrorObject, error) {
	nonce, apierr, err := GetNonce(address, r)
	if err != nil {
		return apierr, err
	}
	err = VerifySignature(NonceToHash(nonce), signature, address)
	if err != nil {
		return problems.BadRequest(errors.New("signature verification failed"))[0], err
	}
	return nil, nil
}
