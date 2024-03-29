package grpc

import (
	"context"
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/pkg/account_v1"
)

func (i *Implementation) DeleteAccount(ctx context.Context, req *account_v1.DeleteAccountRequest) (*account_v1.DeleteAccountResponse, error) {

	err := i.disable.Disable(uint64(req.AccountId))

	fmt.Println(" IN RPC METHOD DELETE ACCOUNT")

	if err != nil {
		return &account_v1.DeleteAccountResponse{Success: false}, err
	}

	return &account_v1.DeleteAccountResponse{
		Success: true,
	}, nil
}
