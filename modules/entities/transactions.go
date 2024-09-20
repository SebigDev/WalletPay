package entities

import (
	"fmt"
	"time"

	"github.com/sebigdev/walletpay/modules/daos"
	"github.com/sebigdev/walletpay/modules/responses"
	"github.com/sebigdev/walletpay/modules/vo"

	"github.com/google/uuid"
)

type TransactionType string
type TransactionDirection string

const (
	Internal TransactionType = "Internal"
	External TransactionType = "External"
)
const (
	Credit TransactionDirection = "Credit"
	Debit  TransactionDirection = "Debit"
)

type ErrorTransaction struct {
	Reason string
}

func (e ErrorTransaction) Error() string {
	return fmt.Sprintf(e.Reason)
}

type Debitor struct {
	Wallet    string
	PayAmount float64
	Currency  string
	Name      string
	UserId    string
}

type Beneficiary struct {
	Wallet    string
	PayAmount float64
	Currency  string
	Name      string
	UserId    string
}

type LegalName struct {
	FirstName string
	LastName  string
}

type Transaction struct {
	ID          string
	To          Beneficiary
	From        Debitor
	Description string
	CreatedAt   time.Time
	PaymentRef  string
	TxCurrency  string
}

func (ln *LegalName) String() string {
	return fmt.Sprintf("%s %s", ln.FirstName, ln.LastName)
}

func NewBeneficiary(userid string, wallet vo.WalletNumber, money vo.Money, name vo.CreditorName) *Beneficiary {
	return &Beneficiary{
		Wallet:    wallet.String(),
		PayAmount: float64(money.Amount),
		Name:      name.String(),
		Currency:  money.Currency.String(),
		UserId:    userid,
	}
}

func NewDebitor(person *Person, wallet vo.WalletNumber, money vo.Money) *Debitor {
	p := LegalName{person.user.FirstName, person.user.LastName}
	return &Debitor{
		Wallet:    wallet.String(),
		PayAmount: float64(money.Amount),
		Name:      p.String(),
		Currency:  money.Currency.String(),
		UserId:    person.GetUserID(),
	}
}

func NewTransaction(from Debitor, to Beneficiary, purpose string) *Transaction {
	return &Transaction{
		To:          to,
		From:        from,
		Description: purpose,
		CreatedAt:   time.Now().UTC(),
		ID:          uuid.NewString(),
		PaymentRef:  fmt.Sprintf("TXRef%d", time.Now().UTC().UnixMilli()),
		TxCurrency:  to.Currency,
	}
}

func (trans *Transaction) Create() *daos.TransactionDao {
	return &daos.TransactionDao{
		ReceiverAccount: daos.ToDao{
			ToWalletNumber: trans.To.Wallet,
			ToName:         trans.To.Name,
			ToCurrency:     trans.To.Currency,
		},
		SenderAccount: daos.FromDao{
			FromWalletNumber: trans.From.Wallet,
			FromName:         trans.From.Name,
			FromCurrency:     trans.From.Currency,
		},
		Amount:               trans.From.PayAmount,
		Description:          trans.Description,
		CreatedAt:            trans.CreatedAt,
		Sender:               trans.From.UserId,
		Receiver:             trans.To.UserId,
		TransactionReference: trans.PaymentRef,
		ID:                   trans.ID,
		Currency:             trans.TxCurrency,
	}
}

func ToResponse(t daos.TransactionDao, userId string) *responses.TransactionFullResponse {
	return &responses.TransactionFullResponse{
		ID:                   t.ID,
		TransactionReference: t.TransactionReference,
		BeneficiaryCurrency:  t.ReceiverAccount.ToCurrency,
		BeneficiaryAccount:   t.ReceiverAccount.ToWalletNumber,
		BeneficiaryName:      t.ReceiverAccount.ToName,
		Amount:               t.Amount,
		OriginatorAccount:    t.SenderAccount.FromWalletNumber,
		OriginatorCurrency:   t.SenderAccount.FromCurrency,
		DebitorName:          t.SenderAccount.FromName,
		Sender:               t.Sender,
		Receiver:             t.Receiver,
		Description:          t.Description,
		CreatedAt:            t.CreatedAt,
		Direction:            getDirection(userId, t.Sender),
	}
}

func getDirection(userId, sender string) string {
	if userId == sender {
		return "Debit"
	}
	return "Credit"
}
