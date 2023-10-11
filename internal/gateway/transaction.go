package gateway

import "github.com.br/rcoelhorsc/fc-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
