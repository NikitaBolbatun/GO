package shop

import (
	"bytes"
	"context"
	"encoding/csv"
	"github.com/pkg/errors"
	"strconv"
)

type ImportProductsError struct {
	product Product
	err     error
}
type ImportAccountError struct {
	Account Account
	err     error
}

func (productMutex ProductMutex) ImportProductsCSV(data []byte) (error []ImportProductsError) {
	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return append(error, ImportProductsError{Product{}, errors.New("import product error")})
	}
	if len(records) < 1 {
		return append(error, ImportProductsError{Product{}, errors.New("empty")})
	}
	sizeImport := 1000
	length := len(records)
	col := length % sizeImport
	count := length / sizeImport
	if col > 0 {
		count++
	}
	resChan := make(chan map[string]Product, 1)
	errChan := make(chan ImportProductsError, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < length; i += sizeImport {
		start := i
		end := start
		if length < sizeImport {
			end = col
		} else if i+sizeImport > length {
			end += col
		} else {
			end += sizeImport
		}
		go productMutex.ImportProductsCSVRecords(ctx, records[start:end], resChan, errChan)
	}
	products := make(map[string]Product)
	for i := 0; i < count; i++ {
		select {
		case result := <-resChan:
			for key := range result {
				products[key] = result[key]
			}
		case err := <-errChan:
			cancel()
			error = append(error, err)
		}
	}

	if len(error) > 0 {
		return error
	}

	productMutex.Lock()
	defer productMutex.Unlock()
	for prd := range products {
		productMutex.Product[prd] = products[prd]
	}

	return nil
}

func (productMutex ProductMutex) ImportProductsCSVRecords(ctx context.Context, records [][]string, resChan chan<- map[string]Product,
	errChan chan<- ImportProductsError) {

	products := make(map[string]Product)

	for _, record := range records {
		select {
		case <-ctx.Done():
			errChan <- ImportProductsError{Product{}, errors.New("cancel")}
			return
		default:
		}
		name := record[0]
		price, err := strconv.ParseFloat(record[1], 32)
		if err != nil {
			errChan <- ImportProductsError{Product{name, float32(price), 0}, errors.New("parse float error")}
			return
		}
		typ, err := strconv.Atoi(record[2])
		if err != nil {
			errChan <- ImportProductsError{Product{name, float32(price), ProductType(typ)}, errors.New("atom error")}
			return
		}
		product := NewProduct(name, float32(price), ProductType(typ))
		products[product.Name] = product
	}

	resChan <- products
}
func (accountMutex AccountMutex) ImportAccountsCSV(data []byte) (error []ImportAccountError) {
	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return append(error, ImportAccountError{Account{}, errors.Wrap(err, "import product error")})
	}
	if len(records) < 2 {
		return append(error, ImportAccountError{Account{}, errors.New("empty data")})
	}
	sizeImport := 1000
	length := len(records)
	col := length % sizeImport
	count := length / sizeImport
	if col > 0 {
		count++
	}
	resChan := make(chan map[string]Account, 1)
	errChan := make(chan ImportAccountError, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < length; i += sizeImport {

		start := i
		end := start

		if length < sizeImport {
			end = col
		} else if i+sizeImport > length {
			end += col
		} else {
			end += sizeImport
		}
		go accountMutex.ImportAccountsCSVRecords(ctx, records[start:end], resChan, errChan)
	}
	accounts := make(map[string]Account)
	for i := 0; i < count; i++ {
		select {
		case result := <-resChan:
			for acc := range result {
				accounts[acc] = result[acc]
			}
		case err := <-errChan:
			cancel()
			error = append(error, err)
		}
	}
	if len(error) > 0 {
		return error
	}
	accountMutex.Lock()
	defer accountMutex.Unlock()
	for acc := range accounts {
		accountMutex.Account[acc] = accounts[acc]
	}

	return nil
}

func (accountMutex AccountMutex) ImportAccountsCSVRecords(ctx context.Context, records [][]string, resChan chan<- map[string]Account, errChan chan<- ImportAccountError) {
	accounts := make(map[string]Account)
	for _, record := range records {
		select {
		case <-ctx.Done():
			errChan <- ImportAccountError{Account{}, errors.New("cancel")}
			return
		default:
		}
		name := record[0]
		balance, err := strconv.ParseFloat(record[1], 32)
		if err != nil {
			errChan <- ImportAccountError{Account{name, float32(balance), 0}, errors.New( "parse float err")}
			return
		}
		typ, err := strconv.Atoi(record[2])
		if err != nil {
			errChan <- ImportAccountError{Account{name, float32(balance), AccountType(typ)}, errors.New( "atom err")}
			return
		}
		account := Account{name, float32(balance), AccountType(typ)}
		accounts[name] = account
	}
	resChan <- accounts
}
