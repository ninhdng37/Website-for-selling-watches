package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"watchWebsite/application"
	"watchWebsite/domain/entity"
	"watchWebsite/infrastructure/auth"
	"watchWebsite/infrastructure/persistence"
	"watchWebsite/interfaces"
	"watchWebsite/notifications"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {
	host := os.Getenv("postgresql_host")
	user := os.Getenv("postgresql_user")
	password := os.Getenv("postgresql_password")
	dbname := os.Getenv("postgresql_dbname")
	port := os.Getenv("postgresql_port")

	services, err := persistence.NewRepositories(host, port, user, dbname, password)
	if err != nil {
		panic(err)
	}
	defer services.Close()

	customers := interfaces.NewCustomers(services.Customer)
	employees := interfaces.NewEmployee(services.Employee)
	orders := interfaces.NewOrder(services.Order)
	orderDetails := interfaces.NewOrderDetail(services.OrderDetail)
	brands := interfaces.NewBrands(services.Brand)
	provinces := interfaces.NewProvinces(services.Province)
	typesOfWatch := interfaces.NewTypeOfWatch(services.TypeOfWatch)
	watch := interfaces.NewWatchs(services.Watch)
	auth := auth.NewAuth(services.Customer, services.Employee)
	// inter := NewInterface(services.Watch)
	watches := interfaces.NewWatchs(services.Watch)
	s := NewServer(
		services.DB,
		services.Watch,
		services.Customer,
		services.Ward,
		services.District,
		services.Province,
		services.Order,
		services.OrderDetail,
		services.Brand,
		services.TypeOfWatch)

	server := echo.New()

	groupCustomer := server.Group("/customer")
	groupManager := server.Group("/manager")
	groupEmployee := server.Group("/employee")
	server.GET("/get-watch-relative", watches.GetWatchByNameRelative)
	GetGroupCustomer(groupCustomer, s, customers, watches, orders, auth)
	GetGroupManager(groupManager, employees, auth)
	GetGroupEmployee(groupEmployee, brands, provinces, typesOfWatch, orders, orderDetails, watch, s, auth)
	server.Logger.Fatal(server.Start(":8888"))
}

func GetGroupCustomer(groupCustomer *echo.Group,
	s *Server,
	customers *interfaces.Customers,
	watches *interfaces.Watchs,
	orders *interfaces.Order,
	auth *auth.Auth) {
	groupCustomer.Use(middleware.Static("D:/DoAnThucTap/Front-end"))
	groupCustomer.POST("/create", customers.CreateCustomer, auth.SendEmailToGmail)
	groupCustomer.POST("/login", auth.Login)
	groupCustomer.POST("/send-email", auth.SendEmailForForgotingPassword)
	groupCustomer.PUT("/reset-password", auth.ResetPassword)
	groupCustomer.POST("/add-order", s.AddOrder)
	groupCustomer.PUT("/cancel-orders", orders.UpdateOrderByCustomer)
	groupCustomer.GET("/verify-email/:tokenString", auth.VerifyEmail)
	groupCustomer.GET("/log-out", auth.Logout)
	groupCustomer.GET("/get-orders", s.GetOrderAndDetails)
	groupCustomer.GET("/verify-email-forgotting-password/:tokenString", auth.VerifyEmailForForgottingPassword)

	//presentation
	groupCustomer.GET("/login-form", GetLoginInterface)
	groupCustomer.GET("/home", GetHomeInterface)
	groupCustomer.GET("/registry", GetRegistryInterface)
	groupCustomer.GET("/expired-token", GetExpiredTokenInterface)
	groupCustomer.GET("/success-verification", GetSuccessVerificationInterface)
	groupCustomer.GET("/forgot-password", GetForgotingPasswordInterface)
	groupCustomer.GET("/reset-password-page", GetResettingPasswordInterface)
	groupCustomer.GET("/get-all", s.GetAll)
}

func GetGroupManager(groupManager *echo.Group,
	employees *interfaces.Employee,
	auth *auth.Auth) {
	groupManager.Use(middleware.Static("D:/DoAnThucTap/Front-end"))
	groupManager.POST("/create", employees.CreateEmployee)
	groupManager.POST("/login", auth.LoginForManager)
	groupManager.GET("/get-all-employees", employees.GetAllEmployee)
	groupManager.PUT("/update-employee", employees.UpdateEmployee)
	groupManager.DELETE("/delete-employee/:employeeID", employees.DeleteEmployee)
	//presentation
	groupManager.GET("/login-form", GetLoginInterfaceForManager)
	groupManager.GET("/home", GetHomeInterfaceForManager)
}

func GetGroupEmployee(groupEmployee *echo.Group,
	brand *interfaces.Brand,
	province *interfaces.Provinces,
	typeOfWatch *interfaces.TypeOfWatch,
	order *interfaces.Order,
	orderDetail *interfaces.OrderDetail,
	watch *interfaces.Watchs,
	s *Server,
	auth *auth.Auth,
) {
	groupEmployee.Use(middleware.Static("D:/DoAnThucTap/Front-end"))
	groupEmployee.POST("/login", auth.LoginForManager)

	groupEmployee.GET("/get-all-brands", brand.GetAllBrands)
	groupEmployee.POST("/create-brand", brand.CreateBrand)
	groupEmployee.PUT("/update-brand", brand.UpdateBrand)
	groupEmployee.DELETE("/delete-brand/:brandID", brand.DeleteBrand)

	groupEmployee.GET("/get-all-provinces", province.GetAllProvince)
	groupEmployee.POST("/create-province", province.CreateProvince)
	groupEmployee.PUT("/update-province", province.UpdateProvince)
	groupEmployee.DELETE("/delete-province/:provinceID", province.DeleteProvince)

	groupEmployee.GET("/get-all-types-of-watch", typeOfWatch.GetAllTypesOfWatch)
	groupEmployee.POST("/create-type-of-watch", typeOfWatch.CreateTypeOfWatch)
	groupEmployee.PUT("/update-type-of-watch", typeOfWatch.UpdateTypeOfWatch)
	groupEmployee.DELETE("/delete-type-of-watch/:typeOfWatchID", typeOfWatch.DeleteTypeOfWatch)

	groupEmployee.GET("/get-all-orders", order.GetAllOrders)
	groupEmployee.PUT("/update-order-status", order.UpdateOrderStatus)

	groupEmployee.GET("/get-order-details-by-order-id/:orderID", orderDetail.GetOrderDetailByOrderID)

	groupEmployee.GET("/get-all-watches", s.GetAllWatchs)
	groupEmployee.POST("/create-watch", watch.CreateWatch)
	groupEmployee.PUT("/update-watch", watch.UpdateWatchCus)
	groupEmployee.GET("/get-brans-and-types-of-watch", s.GetBrandsAndTypesOfWatch)
	groupEmployee.DELETE("/delete-watch/:watchID", watch.DeleteWatch)

	//presentation
	groupEmployee.GET("/login-form", GetLoginInterfaceForEmployee)
	groupEmployee.GET("/home", GetHomeInterfaceForEmployee)

}

type Server struct {
	db          *persistence.Database
	watch       application.WatchAppInterface
	cus         application.CustomerAppInterface
	ward        application.WardAppInterface
	district    application.DistrictAppInterface
	province    application.ProvinceAppInterface
	order       application.OrderAppInterface
	orderDetail application.OrderDetailAppInterface
	brand       application.BrandAppInterface
	typeOfWatch application.TypeOfWatchAppInterface
}

func NewServer(db *persistence.Database,
	watch application.WatchAppInterface,
	cus application.CustomerAppInterface,
	ward application.WardAppInterface,
	district application.DistrictAppInterface,
	province application.ProvinceAppInterface,
	order application.OrderAppInterface,
	orderDetail application.OrderDetailAppInterface,
	brand application.BrandAppInterface,
	typeOfWatch application.TypeOfWatchAppInterface) *Server {
	return &Server{
		db:          db,
		watch:       watch,
		cus:         cus,
		ward:        ward,
		district:    district,
		province:    province,
		order:       order,
		orderDetail: orderDetail,
		brand:       brand,
		typeOfWatch: typeOfWatch,
	}
}

func (s *Server) GetAll(c echo.Context) error {
	cookie, err := c.Cookie("token")
	notices := make(map[string]string)
	watchesCopy := []struct {
		WatchName string
		Price     uint32
		Image     string
		Quantity  uint32
	}{}
	verifications := struct {
		IsCustomer bool              `json:"isCustomer"`
		Wards      []entity.Ward     `json:"wards"`
		Districts  []entity.District `json:"districts"`
		Provinces  []entity.Province `json:"provinces"`
		Watches    []struct {
			WatchName string
			Price     uint32
			Image     string
			Quantity  uint32
		} `json:"watches"`
	}{}

	if err != nil {
		if err.Error() == "http: named cookie not present" {
			watches, err := s.watch.GetAllWatch()
			if err != nil {
				notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
				return c.JSON(http.StatusInternalServerError, notices)
			}

			wards, err := s.ward.GetAllWard()
			if err != nil {
				notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
				return c.JSON(http.StatusInternalServerError, notices)
			}

			districts, err := s.district.GetAllDistrict()
			if err != nil {
				notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
				return c.JSON(http.StatusInternalServerError, notices)
			}

			provinces, err := s.province.GetAllProvince()
			if err != nil {
				notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
				return c.JSON(http.StatusInternalServerError, notices)
			}

			for _, value := range watches {
				watch := struct {
					WatchName string
					Price     uint32
					Image     string
					Quantity  uint32
				}{}
				watch.WatchName = strings.TrimSpace(value.WatchName)
				watch.Price = *value.Price
				watch.Image = strings.TrimSpace(value.Image)
				watch.Quantity = *value.Quantity
				watchesCopy = append(watchesCopy, watch)
			}
			verifications.IsCustomer = false
			verifications.Watches = watchesCopy
			verifications.Wards = wards
			verifications.Districts = districts
			verifications.Provinces = provinces
			return c.JSON(http.StatusContinue, verifications)
		}
	}

	if cookie.Value == "" {
		watches, err := s.watch.GetAllWatch()
		if err != nil {
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}

		wards, err := s.ward.GetAllWard()
		if err != nil {
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}

		districts, err := s.district.GetAllDistrict()
		if err != nil {
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}

		provinces, err := s.province.GetAllProvince()
		if err != nil {
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}

		for _, value := range watches {
			watch := struct {
				WatchName string
				Price     uint32
				Image     string
				Quantity  uint32
			}{}
			watch.WatchName = strings.TrimSpace(value.WatchName)
			watch.Price = *value.Price
			watch.Image = strings.TrimSpace(value.Image)
			watch.Quantity = *value.Quantity
			watchesCopy = append(watchesCopy, watch)
		}
		verifications.IsCustomer = false
		verifications.Watches = watchesCopy
		verifications.Wards = wards
		verifications.Districts = districts
		verifications.Provinces = provinces
		return c.JSON(http.StatusContinue, verifications)
	}
	token, err := auth.IsTokenValid(cookie.Value)

	if err != nil {
		if err == echo.ErrUnauthorized {
			notices[notifications.UNAUTHORIZED_ACCESS_ERROR] = notifications.EXPIRED_TOKEN_ERROR
			return c.JSON(http.StatusUnauthorized, "notices")
		} else if err == echo.ErrBadRequest {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.EXPIRED_TOKEN_ERROR
			return c.JSON(http.StatusUnauthorized, "notices")
		}
	}

	claims := token.Claims.(jwt.MapClaims)

	isCustomer := claims["isCustomer"].(bool)
	if !isCustomer {
		notices[notifications.UNAUTHORIZED_ACCESS_ERROR] = notifications.EXPIRED_TOKEN_ERROR
		return c.JSON(http.StatusUnauthorized, "notices")
	}

	watches, err := s.watch.GetAllWatch()
	if err != nil {
		notices := make(map[string]string)
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	for _, value := range watches {
		watch := struct {
			WatchName string
			Price     uint32
			Image     string
			Quantity  uint32
		}{}
		watch.WatchName = strings.TrimSpace(value.WatchName)
		watch.Price = *value.Price
		watch.Image = strings.TrimSpace(value.Image)
		watch.Quantity = *value.Quantity
		watchesCopy = append(watchesCopy, watch)
	}
	wards, err := s.ward.GetAllWard()
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	districts, err := s.district.GetAllDistrict()
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	provinces, err := s.province.GetAllProvince()
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	c.SetCookie(cookie)
	verifications.IsCustomer = true
	verifications.Watches = watchesCopy
	verifications.Wards = wards
	verifications.Districts = districts
	verifications.Provinces = provinces
	return c.JSON(http.StatusOK, verifications)
}

func (s *Server) AddOrder(c echo.Context) error {
	notices := make(map[string]string)
	orderRequest := &entity.OrderRequest{}
	if err := c.Bind(&orderRequest); err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	cookie, err := c.Cookie("token")
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	tokenString := cookie.Value
	if tokenString == "" {
		notices[notifications.INTERNAL_SERVER_ERROR] = "Please login to buy"
		return c.JSON(http.StatusInternalServerError, notices)
	}
	token, err := auth.IsTokenValid(tokenString)
	if err != nil {
		if err == echo.ErrBadRequest {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.BAD_REQUEST_ERROR
			return c.JSON(http.StatusBadRequest, notices)
		} else if err == echo.ErrUnauthorized {
			notices[notifications.INTERNAL_SERVER_ERROR] = "Please login to buy"
			return c.JSON(http.StatusInternalServerError, notices)
		}
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	customer, err := s.cus.GetCustomerByEmail(email)
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	order := &entity.Order{}
	order.CustomerId = customer.CustomerId
	order.ProvinceId = orderRequest.ProvinceId
	order.DistrictId = orderRequest.DistrictId
	order.WardId = orderRequest.WardId
	order.ApartmentNumber = orderRequest.ApartmentNumber
	tx, err := s.db.Begin()
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	orderId, err := s.order.CreateOrder(tx, order)

	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	// orderDetails := make([]entity.OrderDetail)
	for _, value := range orderRequest.Items {
		orderDetail := &entity.OrderDetail{}
		orderDetail.Quantity = value.Quantity
		orderDetail.UnitPrice = value.UnitPrice
		watch, err := s.watch.GetWatchByName(value.Title)
		if err != nil {
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}
		orderDetail.WatchId = watch.WatchId
		orderDetail.OrderId = orderId

		if err != nil {
			fmt.Println(5)
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}

		defer tx.Rollback()
		err = s.orderDetail.CreateOrderDetail(tx, orderDetail)

		if err != nil {
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}
		*watch.Quantity = *watch.Quantity - *orderDetail.Quantity
		err = s.watch.UpdateWatch(tx, watch)
		if err != nil {
			fmt.Println(6)
			fmt.Println(err)
			notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
			return c.JSON(http.StatusInternalServerError, notices)
		}
	}

	err = tx.Commit()
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	notices[notifications.SUCCESS] = notifications.SUCCESS_ADDING_ORDER
	return c.JSON(http.StatusOK, notices)
}

func (s *Server) GetOrderAndDetails(c echo.Context) error {
	orders := []struct {
		OrderID      uint32
		Status       uint32
		OrderDate    string
		OrderDetails []*struct {
			WatchName  string
			WatchImage string
			Quantity   *uint32
			UnitPrice  *uint32
		}
	}{}
	notices := make(map[string]string)

	// email := c.QueryParam("email")
	cookie, err := c.Cookie("token")
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	tokenString := cookie.Value
	token, err := auth.IsTokenValid(tokenString)
	// customer, err := s.cus.GetCustomerByEmail(emailCustomer)
	if err != nil {
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	orderDetails, err := s.orderDetail.GetOrderDetailByCustomerEmail(email)
	if err != nil {
		fmt.Println(err)
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	ordersMap := make(map[uint32][]*entity.OrderDetailResponse)
	for _, item := range orderDetails {
		ordersMap[*item.OrderId] = append(ordersMap[*item.OrderId], item)
	}
	// fmt.Println(len(ordersMap))
	dateTimeFormat := "2006-01-02 15:04:05"
	for key, value := range ordersMap {

		order := struct {
			OrderID      uint32
			Status       uint32
			OrderDate    string
			OrderDetails []*struct {
				WatchName  string
				WatchImage string
				Quantity   *uint32
				UnitPrice  *uint32
			}
		}{}
		order.OrderID = key
		for _, v := range value {
			orderDetail := &struct {
				WatchName  string
				WatchImage string
				Quantity   *uint32
				UnitPrice  *uint32
			}{}
			orderDetail.WatchName = strings.TrimSpace(v.WatchName)
			orderDetail.Quantity = v.Quantity
			orderDetail.UnitPrice = v.UnitPrice
			order.Status = *v.Status
			order.OrderDate = (*v.OrderDate).Format(dateTimeFormat)
			orderDetail.WatchImage = strings.TrimSpace(v.WatchImage)
			order.OrderDetails = append(order.OrderDetails, orderDetail)
		}
		// order.OrderDetails = value
		orders = append(orders, order)
	}
	// fmt.Println(orders)
	// fmt.Println(ordersMap)
	return c.JSON(http.StatusOK, orders)
}

func (s *Server) GetAllWatchs(c echo.Context) error {
	tx, err := s.db.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	query := `SELECT 
					WB.WATCH_ID, 
					WB.WATCH_NAME, 
					WB.PRICE, 
					WB.IMAGE, 
					WB.STATUS, 
					WB.QUANTITY, 
					WB.BRAND_ID, 
					WB.BRAND_NAME, 
					WB.TYPE_OF_WATCH_ID, 
					TY.TYPE_OF_WATCH_NAME
				FROM 
					(SELECT 
						W.WATCH_ID, 
						W.WATCH_NAME,
						W.PRICE,
						W.IMAGE,
						W.STATUS,
						W.QUANTITY,
						W.BRAND_ID,
						W.TYPE_OF_WATCH_ID,
						B.BRAND_NAME
					FROM WATCHES W INNER JOIN BRANDS B ON W.BRAND_ID = B.BRAND_ID) WB
				INNER JOIN 
					TYPE_OF_WATCHS TY ON TY.TYPE_OF_WATCH_ID = WB.TYPE_OF_WATCH_ID
				ORDER BY WB.WATCH_ID`

	rows, err := tx.Query(query)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	defer rows.Close()

	watches := []*struct {
		WatchId     *uint32            `json:"watchId"`
		WatchName   string             `json:"watchName"`
		Price       *uint32            `json:"price"`
		Image       string             `json:"image"`
		Status      *uint32            `json:"status"`
		Quantity    *uint32            `json:"quantity"`
		Brand       entity.Brand       `json:"brand"`
		TypeOfWatch entity.TypeOfWatch `json:"typeOfWatch"`
	}{}

	for rows.Next() {
		watch := &struct {
			WatchId     *uint32            `json:"watchId"`
			WatchName   string             `json:"watchName"`
			Price       *uint32            `json:"price"`
			Image       string             `json:"image"`
			Status      *uint32            `json:"status"`
			Quantity    *uint32            `json:"quantity"`
			Brand       entity.Brand       `json:"brand"`
			TypeOfWatch entity.TypeOfWatch `json:"typeOfWatch"`
		}{}

		if err := rows.Scan(
			&watch.WatchId,
			&watch.WatchName,
			&watch.Price,
			&watch.Image,
			&watch.Status,
			&watch.Quantity,
			&watch.Brand.BrandID,
			&watch.Brand.BrandName,
			&watch.TypeOfWatch.TypeOfWatchID,
			&watch.TypeOfWatch.TypeOfWatchName); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		watch.Image = strings.TrimSpace(watch.Image)
		watch.Brand.TrimSpace()
		watch.TypeOfWatch.TrimSpace()

		watches = append(watches, watch)
	}

	brands, err := s.brand.GetAllBrands()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for _, brand := range brands {
		brand.TrimSpace()
	}

	typesOfWatch, err := s.typeOfWatch.GetAllTypesOfWatch()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for _, typeOfWatch := range typesOfWatch {
		typeOfWatch.TrimSpace()
	}

	// data := struct{

	// }{}

	data := struct {
		Watches []*struct {
			WatchId     *uint32            `json:"watchId"`
			WatchName   string             `json:"watchName"`
			Price       *uint32            `json:"price"`
			Image       string             `json:"image"`
			Status      *uint32            `json:"status"`
			Quantity    *uint32            `json:"quantity"`
			Brand       entity.Brand       `json:"brand"`
			TypeOfWatch entity.TypeOfWatch `json:"typeOfWatch"`
		} `json:"watches"`

		Brands       []*entity.Brand       `json:"brands"`
		TypesOfWatch []*entity.TypeOfWatch `json:"typesOfWatch"`
	}{}

	data.Watches = watches
	data.Brands = brands
	data.TypesOfWatch = typesOfWatch
	return c.JSON(http.StatusOK, data)
}

func (s *Server) GetBrandsAndTypesOfWatch(c echo.Context) error {
	// tx, err := s.db.Begin()

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	data := struct {
		Brands       []*entity.Brand       `json:"brands"`
		TypesOfWatch []*entity.TypeOfWatch `json:"typesOfWatch"`
	}{}

	brands, err := s.brand.GetAllBrands()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	typesOfWatch, err := s.typeOfWatch.GetAllTypesOfWatch()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data.Brands = brands
	data.TypesOfWatch = typesOfWatch

	return c.JSON(http.StatusOK, data)
}
