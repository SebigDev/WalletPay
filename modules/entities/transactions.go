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

const (
	Internal TransactionType = "Internal"
	External TransactionType = "External"
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

func NewBeneficiary(wallet vo.WalletNumber, money vo.Money, name vo.CreditorName) *Beneficiary {
	return &Beneficiary{
		Wallet:    wallet.String(),
		PayAmount: float64(money.Amount),
		Name:      name.String(),
		Currency:  money.Currency.String(),
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
		Beneficiary: daos.ToDao{
			ToWalletNumber: trans.To.Wallet,
			ToName:         trans.To.Name,
			ToCurrency:     trans.To.Currency,
		},
		Originator: daos.FromDao{
			FromWalletNumber: trans.From.Wallet,
			FromName:         trans.From.Name,
			FromCurrency:     trans.From.Currency,
		},
		Amount:               trans.From.PayAmount,
		Description:          trans.Description,
		CreatedAt:            trans.CreatedAt,
		UserId:               trans.From.UserId,
		TransactionReference: trans.PaymentRef,
		ID:                   trans.ID,
		Currency:             trans.TxCurrency,
	}
}

func ToResponse(t daos.TransactionDao) *responses.TransactionFullResponse {
	return &responses.TransactionFullResponse{
		ID:                   t.ID,
		TransactionReference: t.TransactionReference,
		BeneficiaryCurrency:  t.Beneficiary.ToCurrency,
		BeneficiaryAccount:   t.Beneficiary.ToWalletNumber,
		BeneficiaryName:      t.Beneficiary.ToName,
		Amount:               t.Amount,
		OriginatorAccount:    t.Originator.FromWalletNumber,
		OriginatorCurrency:   t.Originator.FromCurrency,
		DebitorName:          t.Originator.FromName,
		UserId:               t.UserId,
		Description:          t.Description,
		CreatedAt:            t.CreatedAt,
	}
}
