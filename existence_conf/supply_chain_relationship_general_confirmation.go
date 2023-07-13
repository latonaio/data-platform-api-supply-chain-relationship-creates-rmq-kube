package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-supply-chain-relationship-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) supplyChainRelationshipGeneralExistenceConf(input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	if input.General.SupplyChainRelationshipID == nil ||
		input.General.Buyer == nil ||
		input.General.Seller == nil {
		*exconfErrMsg = "cannot specify null keys"
		return
	}
	res, err := c.supplyChainRelationshipGeneralExistenceConfRequest(*input.General.SupplyChainRelationshipID, *input.General.Buyer, *input.General.Seller, input, existenceMap, mtx, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
	}
	if res != "" {
		*exconfErrMsg = res
	}
}

func (c *ExistenceConf) supplyChainRelationshipGeneralExistenceConfRequest(supplyChainRelationshipID int, buyer int, seller int, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	key := "SupplyChainRelationshipGeneral"
	keys := newResult(map[string]interface{}{
		"SupplyChainRelationshipID": supplyChainRelationshipID,
		"Buyer":                     buyer,
		"Seller":                    seller,
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

	data := &SupplyChainRelationshipGeneralReturn{
		SupplyChainRelationshipID: supplyChainRelationshipID,
		Buyer:                     buyer,
		Seller:                    seller,
	}
	req.SupplyChainRelationshipGeneralReturn = data

	exist, err = c.exconfRequest(req, key, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}
