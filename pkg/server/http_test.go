package server

import (
	"encoding/json"
	"fmt"
	"testing"

	"chatserver/pkg/domain"
	"chatserver/pkg/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tokopedia/tdk/go/app/http"
)

func TestHttp_NewOrder(t *testing.T) {
	type testcase struct {
		mock        func()
		errorExpect error
	}

	// initialize gomock controller
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	// this is mock for http.TdkContext
	// with this object we can mock anything related with http handler
	ctxMock := http.NewMockTdkContext(ctl)

	// notice here we're not using real object of resource
	// we use mocks instead
	ordResource := domain.NewMockOrderResourceItf(ctl)

	userResource := domain.NewMockUserResourceItf(ctl)

	cases := []testcase{
		// normal case
		{
			mock: func() {

				invoice := "INV/123"
				reqBody := usecase.Order{
					UserID:    1001,
					Quantity:  5,
					ProductID: 1002,
				}
				byt, _ := json.Marshal(reqBody)
				// first we are expecting Body() function to be called
				// and then we returns the mocked value
				ctxMock.EXPECT().Body().Return(byt)

				// keep in mind the domain is called inside the module
				// so we are mocking the flow one by one

				order := domain.Order{
					Invoice:   invoice,
					Quantity:  reqBody.Quantity,
					ProductID: reqBody.ProductID,
				}

				userResource.EXPECT().FindUser(reqBody.UserID).Return(nil)

				ordResource.EXPECT().GetStock(reqBody.ProductID).Return(10)

				ordResource.EXPECT().InsertOrder(&order).Return(nil)

				txt := fmt.Sprintf("invoice created: %s", invoice)
				ctxMock.EXPECT().Write([]byte(txt))
			},
			// we dont expect any error
			errorExpect: nil,
		},
		// ...
		// you can add any case you want
	}

	// we inject mocked domain into order usecase
	ordDomain := domain.InitOrderDomain(ordResource)
	userDomain := domain.InitUserDomain(userResource)
	orderUsecase = usecase.InitOrderUsecase(cfg, ordDomain, userDomain)

	for _, cs := range cases {
		cs.mock()
		err := handleNewOrder(ctxMock)
		assert.Equal(t, err, cs.errorExpect)
	}
}
