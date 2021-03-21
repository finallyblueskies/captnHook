package broker

import (
	"github.com/bareish/captnHook/pkg/services"
)

// ManageBrokerService ...
type ManageBrokerService struct {
	CryptoBroker services.BrokerService
	StockBroker services.BrokerService
	ForexBroker services.BrokerService
}

// Setup will initialize each broker service
func (m *ManageBrokerService) Setup() {
	// todo: maybe implement loop
	//service := reflect.ValueOf(m).Elem()
	//for i := 0; i < service.NumField(); i++ {
	//	if !service.IsNil() {
	//		service.Field(i).Elem().
	//	}
	//}
	nilCheck(m.ForexBroker)
	nilCheck(m.StockBroker)
	nilCheck(m.CryptoBroker)
}

func(m *ManageBrokerService) Get() services.BrokerManager {
	return services.BrokerManager{
		StockService: m.StockBroker,
		CryptoService: m.CryptoBroker,
		ForexService: m.ForexBroker,
	}

}

// nilCheck will setup a service on the Manager if the field is not empty
func nilCheck(service services.BrokerService)  {
	if service != nil {
		service.Setup()
	}
}
