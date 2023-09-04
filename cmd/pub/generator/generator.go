package generator

import (
	models "OrderService/internal/models"
	"bufio"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func NewOrder() *models.Order {
	uid := uuid.NewV4().String()
	trackNumber := randValueString("C:\\dev\\Wildberries\\OrderService\\data\\track_number.txt")
	return &models.Order{
		OrderUid:    uid,
		TrackNumber: trackNumber,
		Entry:       "WBILAAAA",
		Delivery: models.Delivery{
			Name:    randValueString("C:\\dev\\Wildberries\\OrderService\\data\\name.txt") + strconv.Itoa(rand.Intn(100000)+1000),
			Phone:   randValueString("C:\\dev\\Wildberries\\OrderService\\data\\phone.txt"),
			Zip:     randIntValue(),
			City:    randValueString("C:\\dev\\Wildberries\\OrderService\\data\\city.txt"),
			Address: randValueString("C:\\dev\\Wildberries\\OrderService\\data\\address.txt"),
			Region:  randValueString("C:\\dev\\Wildberries\\OrderService\\data\\address.txt"),
			Email:   randValueString("C:\\dev\\Wildberries\\OrderService\\data\\email.txt"),
		},
		Payment: models.Payment{
			Transaction:  uid,
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       uint(rand.Intn(100000)) + 1000,
			PaymentDT:    0,
			Bank:         "Sber",
			DeliveryCost: rand.Intn(10000) + 1000,
			GoodsTotal:   rand.Intn(1000),
			CustomFee:    0,
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmID:              99,
		DateCreated:       randValueString("C:\\dev\\Wildberries\\OrderService\\data\\data.txt"),
		OofShard:          "1",
		Items: []models.Item{
			{
				ChrtID:      uint64(rand.Intn(10000000)) + 1000000,
				TrackNumber: trackNumber,
				Price:       uint64(rand.Intn(100000) + 100),
				Rid:         uid, // rofl
				Name:        randValueString("C:\\dev\\Wildberries\\OrderService\\data\\items.txt"),
				Sale:        uint64(rand.Intn(100)),
				Size:        "0",
				TotalPrice:  uint64(rand.Intn(100000) + 60),
				NmId:        2389215,
				Brand:       randValueString("C:\\dev\\Wildberries\\OrderService\\data\\brand.txt"),
				Status:      202,
				OrderUid:    uid,
			},
		},
	}
}

func randValueString(file string) string {
	data, err := os.Open(file)
	if err != nil {
		logrus.Fatalf("failed to read file: %s", err.Error())
	}
	defer func(data *os.File) {
		err := data.Close()
		if err != nil {
			logrus.Fatalf("file connection close error: %s", err.Error())
		}
	}(data)

	fileScanner := bufio.NewScanner(data)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	rand.NewSource(time.Now().Unix())

	randomIndex := rand.Intn(len(lines))
	randomValue := lines[randomIndex]

	return randomValue
}

func randIntValue() string {
	rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.Intn(10000000)
	formattedNumber := fmt.Sprintf("%07d", randomNumber)
	return formattedNumber
}

func parseDate(date string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		fmt.Println("time conversion error::", err.Error())
	}
	return parsedTime
}
