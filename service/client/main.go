package main

import (
	coffeecov1 "buf.build/gen/go/sideshow/coffeeco/protocolbuffers/go/sideshow/coffeeco/v1"
	"connectrpc.com/connect"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	coffeeservicev1 "github.com/datacircus/highwire/gen/coffeeservice/v1"
	"github.com/datacircus/highwire/gen/coffeeservice/v1/coffeeservicev1connect"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var verveCoffee = map[string]*coffeecov1.Coffee{
	"aster": &coffeecov1.Coffee{
		Name:   "ASTER",
		Vendor: coffeecov1.Coffee_COFFEE_VENDOR_VERVE,
		Roast:  coffeecov1.Coffee_ROAST_MEDIUM,
		FlavorProfile: []coffeecov1.Coffee_FlavorProfile{
			coffeecov1.Coffee_FLAVOR_PROFILE_FRUITY,
			coffeecov1.Coffee_FLAVOR_PROFILE_BROWN_SUGAR,
		},
	},
	"sermon": &coffeecov1.Coffee{
		Name:   "SERMON",
		Vendor: coffeecov1.Coffee_COFFEE_VENDOR_VERVE,
		Roast:  coffeecov1.Coffee_ROAST_MEDIUM,
		FlavorProfile: []coffeecov1.Coffee_FlavorProfile{
			coffeecov1.Coffee_FLAVOR_PROFILE_COCOA,
			coffeecov1.Coffee_FLAVOR_PROFILE_BLUEBERRY,
			coffeecov1.Coffee_FLAVOR_PROFILE_SWEET_AROMATICS,
		},
	},
	"streetlevel": &coffeecov1.Coffee{
		Name:   "STREETLEVEL",
		Vendor: coffeecov1.Coffee_COFFEE_VENDOR_VERVE,
		Roast:  coffeecov1.Coffee_ROAST_MEDIUM,
		FlavorProfile: []coffeecov1.Coffee_FlavorProfile{
			coffeecov1.Coffee_FLAVOR_PROFILE_APPLE,
			coffeecov1.Coffee_FLAVOR_PROFILE_HONEY,
			coffeecov1.Coffee_FLAVOR_PROFILE_SWEET,
		},
	},
}

type CoffeeInfo struct {
	Vendor    coffeecov1.Coffee_CoffeeVendor
	Coffee    *coffeecov1.Coffee
	UnitPrice *coffeecov1.Total
}

type CoffeeMenu struct {
	Menu  map[string]CoffeeInfo
	Index []string
}

func (cm *CoffeeMenu) AddItem(info *CoffeeInfo) *CoffeeMenu {
	index := info.Vendor.String() + "." + info.Coffee.Name
	// test the struct for the item
	if _, ok := cm.Menu[index]; !ok {
		// add the value
		(*cm).Menu[index] = CoffeeInfo{
			Vendor:    info.Vendor,
			Coffee:    info.Coffee,
			UnitPrice: info.UnitPrice,
		}
		// update the index literal
		(*cm).Index = append((*cm).Index, index)
	}
	return cm
}

func (cm *CoffeeMenu) RandomItem() *CoffeeInfo {
	k := rand.Intn(len((*cm).Index))
	item := (*cm).Menu[(*cm).Index[k]]
	return &item
	/*for _, info := range *cm {
		if k == 0 {
			return &info
		}
		k--
	}*/

}

func guid(customer *coffeecov1.Customer) string {
	hash := md5.New()
	hash.Write([]byte(
		customer.Name + customer.CustomerType.String() + customer.FirstSeen.String()))
	md5String := hex.EncodeToString(hash.Sum(nil))

	customerGUID, err := uuid.FromBytes([]byte(md5String[0:16]))
	if err != nil {
		log.Fatal(err)
	}
	return customerGUID.String()
}

func genSoldProduct(coffee *coffeecov1.Coffee) *coffeecov1.Product {

	// randomize the coffee size
	randSize := rand.Intn(len(coffeecov1.CoffeeProduct_Size_name))
	if randSize == 0 {
		randSize = 1
	}
	// randomize the coffee preparation
	randStyle := rand.Intn(len(coffeecov1.CoffeeProduct_Style_name))
	if randStyle == 0 {
		randStyle = 1
	}

	coffeeProduct := coffeecov1.CoffeeProduct{
		Coffee: coffee,
		Size:   coffeecov1.CoffeeProduct_Size(randSize),
		Style:  coffeecov1.CoffeeProduct_Style(randStyle),
	}

	return &coffeecov1.Product{ProductType: &coffeecov1.Product_Coffee{Coffee: &coffeeProduct}}
}

func genOrder(menu *CoffeeMenu, customer *coffeecov1.Customer, store *coffeecov1.Store) *coffeecov1.Order {
	// todo
	// we need to generate between 1 and 5 items
	// depending on the size and style, we can use the UnitPrice (price per pound of coffee)
	// to generate a "market" cost

	var items []*coffeecov1.Product
	var total = &coffeecov1.Total{
		Currency: coffeecov1.Total_CURRENCY_CODE_USD,
		Units:    0,
		Nanos:    0,
	}

	pricing := map[coffeecov1.CoffeeProduct_Size]float32{
		coffeecov1.CoffeeProduct_SIZE_SMALL:  0.24,
		coffeecov1.CoffeeProduct_SIZE_MEDIUM: 0.28,
		coffeecov1.CoffeeProduct_SIZE_LARGE:  0.36,
		coffeecov1.CoffeeProduct_SIZE_XLARGE: 0.40,
	}

	numItems := rand.Intn(5) + 1
	for i := 0; i < numItems; i++ {
		item := menu.RandomItem()
		product := genSoldProduct(item.Coffee)
		items = append(items, product)
		// depending on the UnitPrice and the Style and Size
		// we can take fraction for sale
		offset := pricing[product.GetCoffee().Size]
		cost := uint64(float32(item.UnitPrice.Units) * offset)
		total.Units += cost
		total.Nanos = total.Nanos + item.UnitPrice.Nanos
	}

	// trim the remainder for the Nanos in the total
	if total.Nanos > 100 {
		increaseUnitsBy := total.Nanos / 100
		remainder := total.Nanos % 100
		total.Units += uint64(increaseUnitsBy)
		total.Nanos = remainder
	}

	return &coffeecov1.Order{
		OrderCreated: timestamppb.New(time.Now()),
		PurchasedAt:  store,
		Customer:     customer,
		Items:        items,
		Total:        total,
	}
}

func addAllVerveCoffeeInfo(menu *CoffeeMenu) *CoffeeMenu {
	commonTotal := &coffeecov1.Total{
		Currency: coffeecov1.Total_CURRENCY_CODE_USD,
		Units:    19,
		Nanos:    75,
	}

	asterInfo := CoffeeInfo{
		Vendor:    coffeecov1.Coffee_COFFEE_VENDOR_VERVE,
		Coffee:    verveCoffee["aster"],
		UnitPrice: commonTotal,
	}
	sermonInfo := CoffeeInfo{
		Vendor:    coffeecov1.Coffee_COFFEE_VENDOR_VERVE,
		Coffee:    verveCoffee["sermon"],
		UnitPrice: commonTotal,
	}
	streetLevelInfo := CoffeeInfo{
		Vendor:    coffeecov1.Coffee_COFFEE_VENDOR_VERVE,
		Coffee:    verveCoffee["streetlevel"],
		UnitPrice: commonTotal,
	}
	menu.AddItem(&asterInfo)
	menu.AddItem(&sermonInfo)
	menu.AddItem(&streetLevelInfo)
	return menu
}

func main() {
	log.Println(fmt.Sprintf("Generating CoffeeService Orders"))

	cm := CoffeeMenu{
		Menu:  map[string]CoffeeInfo{},
		Index: []string{},
	}
	addAllVerveCoffeeInfo(&cm)

	item := cm.RandomItem()
	println(fmt.Sprintf("Random Coffee %s", item.Coffee.Name))

	customer := &coffeecov1.Customer{
		Name: "Scott Haines",
		Uuid: "",
		FirstSeen: timestamppb.New(
			time.Date(2024, 6, 7, 9, 44, 15, 0, time.UTC),
		),
		CustomerType:    coffeecov1.Customer_CUSTOMER_TYPE_MEMBER,
		LoyaltyMemberId: uuid.New().String(),
	}

	// using the Customer we can generate a reliable GUID
	customer.Uuid = guid(customer)

	coffeeLocation := &coffeecov1.Store{
		StoreId: uuid.New().String(),
		Created: timestamppb.New(
			time.Date(2023, 2, 4, 6, 30, 0, 0, time.UTC),
		),
		OpenedOn: timestamppb.New(
			time.Date(2024, 6, 7, 10, 24, 0, 0, time.UTC),
		),
		ClosedPermanentlyOn: nil,
		Status:              coffeecov1.Store_LOCATION_STATUS_OPEN_TO_PUBLIC,
	}

	coffeeOrder := genOrder(&cm, customer, coffeeLocation)

	orderRequest := &coffeeservicev1.CoffeeOrderRequest{
		Order: coffeeOrder,
	}

	client := coffeeservicev1connect.NewCoffeeServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	res, err := client.CoffeeOrder(
		context.Background(),
		connect.NewRequest(orderRequest),
	)

	if err != nil {
		log.Println(err)
	}
	log.Println(res.Msg.Response)

}
