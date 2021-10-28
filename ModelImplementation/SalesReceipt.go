package modelimplementation

import "time"

type SalesReceipt struct {
	itsSaleDate 	time.Time
	itsAmount       float64
}

func (sa SalesReceipt) newSalesReceipt(saleDate time.Time,amount float64) SalesReceipt{
	sa.itsSaleDate = saleDate
	sa.itsAmount = amount
	return sa
} 

func (sa SalesReceipt) GetAmount() float64{
	return sa.itsAmount
}

func (sa SalesReceipt)GetSaleDate() time.Time{
	return sa.itsSaleDate
}