package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-supply-chain-relationship-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-supply-chain-relationship-creates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-supply-chain-relationship-creates-rmq-kube/config"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,

) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:  context.Background(),
		conf: conf,
		rmq:  rmq,
	}
}

func (c *DPFMAPICaller) AsyncSupplyChainRelationshipCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
) (interface{}, []error) {
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	var response interface{}
	if input.APIType == "creates" {
		response = c.createSqlProcess(nil, &mtx, input, output, accepter, &errs, log)
	} else if input.APIType == "updates" {
		response = c.updateSqlProcess(nil, &mtx, input, output, accepter, &errs, log)
	}

	return response, nil
}

func (c *DPFMAPICaller) finWait(
	mtx *sync.Mutex,
	finChan chan error,
	ticker *time.Ticker,
) error {
	select {
	case e := <-finChan:
		if e != nil {
			mtx.Lock()
			return e
		}
	case <-ticker.C:
		return xerrors.New("time out")
	}
	return nil
}

func (c *DPFMAPICaller) generalCreate(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	errFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	var err error = nil
	defer func() {
		errFin <- err
	}()
	defer wg.Done()
	return
}

func checkResult(msg rabbitmq.RabbitmqMessage) bool {
	data := msg.Data()
	d, ok := data["result"]
	if !ok {
		return false
	}
	result, ok := d.(string)
	if !ok {
		return false
	}
	return result == "success"
}

func getBoolPtr(b bool) *bool {
	return &b
}
