package service

import "github.com/mariotj/olshop-app/api/entity/model"

type OrderService interface {
	DisplayOrderDetails(orderId uint32) (model.DisplayOrderDetail, error)
}

type orderService struct {

}

func (o orderService) DisplayOrderDetails(orderId uint32) (model.DisplayOrderDetail, error) {
	var order model.Order
	var customer model.Customer
	var employee model.Employee
	var shopMethod model.ShippingMethod
	var orderDetails []model.OrderDetail
	var displayOrderDetail model.DisplayOrderDetail

	databaseService.DB.Model(&model.Order{}).First(&order, "id = ?", orderId)
	databaseService.DB.Model(&model.Customer{}).First(&customer, "id = ?", order.CustomerId)
	databaseService.DB.Model(&model.Employee{}).First(&employee, "id = ?", order.EmployeeId)
	databaseService.DB.Model(&model.ShippingMethod{}).First(&shopMethod, "id = ?", order.ShippingMethodId)
	databaseService.DB.Find(&orderDetails, "order_id = ?", order.ID)

	displayOrderDetail.OrderId = order.ID
	displayOrderDetail.CustomerName = customer.FirstName + customer.LastName
	displayOrderDetail.EmployeeName = employee.FirstName + employee.LastName
	displayOrderDetail.ShippingMethod = shopMethod.ShippingMethod

	var displayProducts []model.DisplayProduct
	var totalPrice float64

	for i:= range orderDetails {
		var product  model.Product
		databaseService.DB.Model(&model.Product{}).First(&product, "id = ?", orderDetails[i].ProductId)

		displayProduct := model.DisplayProduct{
			ProductId:     product.ID,
			ProductName:   product.ProductName,
			Quantity:      orderDetails[i].Quantity,
			UnitPrice:     orderDetails[i].UnitPrice,
			Discount:      orderDetails[i].Discount,
			SubTotalPrice: float64(orderDetails[i].Quantity) * orderDetails[i].UnitPrice - orderDetails[i].Discount,
		}
		displayProducts = append(displayProducts, displayProduct)

		totalPrice += displayProduct.SubTotalPrice
	}

	displayOrderDetail.ProductList = displayProducts
	displayOrderDetail.TotalPayment = totalPrice

	return displayOrderDetail, nil
}

func NewOrderService() OrderService  {
	return orderService{}
}

