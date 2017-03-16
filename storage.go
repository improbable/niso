package niso

import "context"

type Storage interface {
	// Close the resources the Storage potentially holds. (Implements io.Closer)
	Close() error

	// GetClientData fetches the data for a ClientData by id
	GetClientData(ctx context.Context, id string) (*ClientData, error)

	// SaveAuthorize saves authorize data.
	SaveAuthorizeData(ctx context.Context, data *AuthorizeData) error

	// GetAuthorizeData looks up AuthorizeData by a code.
	//// ClientData information MUST be loaded together.
	// Optionally can return error if expired.
	GetAuthorizeData(ctx context.Context, code string) (*AuthorizeData, error)

	// RemoveAuthorize revokes or deletes the authorization code.
	DeleteAuthorizeData(ctx context.Context, code string) error
	//
	// SaveAccess writes AccessData.
	// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
	SaveAccessData(ctx context.Context, data *AccessData) error

	//
	//// LoadAccess retrieves access data by token. ClientData information MUST be loaded together.
	//// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
	//// Optionally can return error if expired.
	//LoadAccess(token string) (*AccessData, error)
	//
	//// RemoveAccess revokes or deletes an AccessData.
	//RemoveAccess(token string) error
	//

	// GetRefreshTokenData retrieves refresh token data from the token string.
	GetRefreshTokenData(ctx context.Context, token string) (*RefreshTokenData, error)

	// SaveRefreshTokenData
	SaveRefreshTokenData(ctx context.Context, data *RefreshTokenData) error

	// DeleteRefreshTokenData revokes or deletes a RefreshToken.
	DeleteRefreshTokenData(ctx context.Context, token string) error
}

type NotFoundError struct {
	err error
}

func (e *NotFoundError) Error() string {
	return e.err.Error()
}
