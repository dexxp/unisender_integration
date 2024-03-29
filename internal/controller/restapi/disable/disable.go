package disable

import (
	"context"
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/pkg/account_v1"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func (i *Impl) NewDisable(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_ = r.ParseForm()
	id := r.Form["account_id"]

	accountID, _ := strconv.Atoi(id[0])

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	defer conn.Close()

	client := account_v1.NewAccountServiceClient(conn)

	req := &account_v1.DeleteAccountRequest{AccountId: int32(accountID)}

	res, err := client.DeleteAccount(context.Background(), req)

	if err != nil {
		fmt.Printf("Failed to call gRPC method: %v\n", err)
	}

	w.Write([]byte(res.String()))
}
