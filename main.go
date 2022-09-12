package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/This-Is-Prince/agri-product-client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var searchClient pb.SearchServiceClient
var listProductClient pb.ListProductServiceClient
var listShopClient pb.ListShopServiceClient
var requestCtx context.Context
var cancel context.CancelFunc

func init() {
	fmt.Println("Starting shop service client")
	requestCtx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to establish client connection to localhost:50051: %v", err)
	}
	searchClient = pb.NewSearchServiceClient(conn)
	listProductClient = pb.NewListProductServiceClient(conn)
	listShopClient = pb.NewListShopServiceClient(conn)
}

type Point struct {
	Long float64
	Lat  float64
}

func main() {
	fmt.Println("Hi")

	// SearchByProduct()
	ListProduct()
	// ListShop()
}

func ListShop() {
	var wg sync.WaitGroup
	// -73.98513559999999, 40.7676919
	wg.Add(1)
	go func() {
		stream, err := listShopClient.ListShop(requestCtx, &pb.ListShopReq{
			Long:        -73.98513559999999,
			Lat:         40.7676919,
			MaxDistance: 15 * 1000,
		})
		if err != nil {
			log.Fatal(err)
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.Shop.Name)
		}
		wg.Done()
	}()

	wg.Wait()
	defer cancel()
}

func ListProduct() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		stream, err := listProductClient.ListProduct(requestCtx, &pb.ListProductReq{
			PriceGte:  200.55,
			PriceLte:  230.25,
			WeightGte: 40.0,
			Name:      "Potash",
		})
		if err != nil {
			log.Fatal(err)
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Hello")
				break
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.Product.Name)
		}
		wg.Done()
	}()

	wg.Wait()
	defer cancel()
}

func SearchByProduct() {
	ids := []string{
		"631dd394590a9ea1f4d9a733", "631dd3ca590a9ea1f4d9a734",
	}
	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			res, err := searchClient.SearchByProduct(requestCtx, &pb.SearchByProductReq{
				Id: id,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.Product.Name)
			wg.Done()
		}(id)
	}
	wg.Wait()
	defer cancel()
}

func SearchNearbyShop() {
	points := []Point{
		{-73.98513559999999, 40.7676919},
		{-73.871194, 40.6730975},
		{-73.9653967, 40.6064339},
		{-73.97822040000001, 40.6435254},
		{-73.7032601, 40.7386417},
		{-74.0259567, 40.6353674},
		{-73.9829239, 40.6580753},
		{-73.839297, 40.78147},
		{-73.95171, 40.767461},
		{-73.9925306, 40.7309346},
		{-73.9634876, 40.6940001},
	}
	var wg sync.WaitGroup

	for _, point := range points {
		wg.Add(1)
		go func(point Point) {
			res, err := searchClient.SearchNearbyShop(requestCtx, &pb.SearchNearbyShopReq{
				Long: point.Long,
				Lat:  point.Lat,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.Shop.Name)
			wg.Done()
		}(point)
	}
	wg.Wait()
	defer cancel()
}
