package domain

import (
	"fmt"

	"github.com/tokopedia/tdk/go/app/resource"
)

type Order struct {
	OrderID   int
	ProductID int
	Quantity  int
	Invoice   string
}

type OrderDomain struct {
	resource OrderResourceItf
}

func InitOrderDomain(rsc OrderResourceItf) OrderDomain {
	return OrderDomain{
		resource: rsc,
	}
}

func (d OrderDomain) IsValidStock(order Order) bool {
	// first get stock from resource
	stock := d.resource.GetStock(order.ProductID)
	// and return the bool
	return stock >= order.Quantity
}

func (d OrderDomain) CreateOrder(order *Order) error {
	// lets generate invoice before we apply it to DB
	idinvoice := "123"
	order.Invoice = fmt.Sprintf("INV/%s", idinvoice)

	// apply it to database
	d.resource.InsertOrder(order)
	return nil
}

type OrderResourceItf interface {
	GetStock(productID int) int
	InsertOrder(*Order) error
}

type OrderResource struct {
	DB    resource.SQLDB
	Redis resource.Redis
}

func (rsc OrderResource) GetStock(productID int) int {
	// we query Redis to get stock
	// example:
	// key := fmt.Sprintf("stock_%v", order.ProductID)
	// stock, _ = Redis.Int(rsc.Redis.Do("GET", key ))

	// lets return dummy stock
	return 10
}

func (rsc OrderResource) InsertOrder(order *Order) error {
	// apply order to database
	// example:
	// rsc.DB.Exec("INSERT INTO tbl_order values(?,?,?)", order)
	return nil
}
