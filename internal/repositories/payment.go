package repositories

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/domains"
)

type paymentRepository struct {
	db *db.PostgresDB
}

func NewPaymentRepository(db *db.PostgresDB) (*paymentRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}
	return &paymentRepository{db}, nil
}

func (p *paymentRepository) Create(ctx context.Context, payment *domains.Payment) (int, error) {
	row := p.db.Pool.QueryRow(ctx, findBankByName, payment.Bank.Name)
	var bank bankModel
	if err := row.Scan(&bank.ID, &bank.Name); err != nil {
		return 0, err
	}

	row = p.db.Pool.QueryRow(ctx, findCurrencyByName, payment.Currency.Name)
	var cur currencyModel
	if err := row.Scan(&cur.ID, &cur.Name); err != nil {
		return 0, err
	}

	row = p.db.Pool.QueryRow(ctx, findProviderByName, payment.Provider.Name)
	var provider providerModel
	if err := row.Scan(&provider.ID, &provider.Name); err != nil {
		return 0, err
	}

	row = p.db.Pool.QueryRow(ctx, payment.Transaction, payment.RequestID, cur.ID, provider.ID, payment.Amount, payment.PaymentDt, bank.ID, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *paymentRepository) FindById(ctx context.Context, id int) (*domains.Payment, error) {
	row := p.db.Pool.QueryRow(ctx, findPaymentById, id)
	var payment paymentModel
	if err := row.Scan(&payment.ID, &payment.Transaction, &payment.RequestID, &payment.CurrencyID, &payment.ProviderID, &payment.Amount, &payment.PaymentDt, &payment.BankID, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee); err != nil {
		return nil, err
	}

	row = p.db.Pool.QueryRow(ctx, findCurrencyById, id)
	var cur currencyModel
	if err := row.Scan(&cur.ID, &cur.Name); err != nil {
		return nil, err
	}

	row = p.db.Pool.QueryRow(ctx, findBankById, id)
	var bank bankModel
	if err := row.Scan(&bank.ID, &bank.Name); err != nil {
		return nil, err
	}

	row = p.db.Pool.QueryRow(ctx, findProviderById, id)
	var provider providerModel
	if err := row.Scan(&provider.ID, &provider.Name); err != nil {
		return nil, err
	}

	return payment.toDomain(cur, bank, provider), nil
}
