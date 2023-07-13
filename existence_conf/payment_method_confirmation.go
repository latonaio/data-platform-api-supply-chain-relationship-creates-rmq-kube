package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-supply-chain-relationship-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) paymentMethodExistenceConf(input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0
	transactions := input.General.Transaction
	for _, transaction := range transactions {
		if transaction.PaymentMethod != nil {
			if len(*transaction.PaymentMethod) == 0 {
				*exconfErrMsg = "cannot specify null keys"
				return
			}
		}
		wg2.Add(1)
		exReqTimes++
		go func(paymentMethod string) {
			res, err := c.paymentMethodExistenceConfRequest(paymentMethod, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}(*transaction.PaymentMethod)
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) paymentMethodExistenceConfRequest(paymentMethod string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	key := "PaymentMethod"
	keys := newResult(map[string]interface{}{
		"PaymentMethod": paymentMethod,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.PaymentMethodReturn.PaymentMethod = paymentMethod

	exist, err = c.exconfRequest(req, key, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}
