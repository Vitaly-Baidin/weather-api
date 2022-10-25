package main

import (
	"context"
	"fmt"
	"github.com/Vitaly-Baidin/weather-api/internal/app"
	"github.com/Vitaly-Baidin/weather-api/internal/service/webapi"
	"log"
)

func main() {

	app.Run()

	//p, err := postgres.New("postgres://root:rootroot@localhost:5432/weather")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//city := repo.NewCity(p)

	city := webapi.NewCity()

	c, err := city.FindByFullAddress(context.Background(), "", "", "Irkutsk")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(c)
	//city := service.NewCity(
	//	repo.NewCity(p),
	//	webapi.NewCity(),
	//)
	//
	//temp := repo.NewTemperatureRepo(p)
	//
	////_, err = city.SetCity(context.Background(), "питер")
	////if err != nil {
	////	log.Fatalln(err)
	////}
	////
	////_, err = city.SetCity(context.Background(), "Moscow")
	////if err != nil {
	////	log.Fatalln(err)
	////}
	//
	////cities, err := city.FindAllCities(context.Background())
	////if err != nil {
	////	log.Fatalln(err)
	////}
	////
	////fmt.Println(cities[0])
	////
	////temperature, err := webapi.NewTemperature().FindByCoord(cities[1].Latitude, cities[1].Longitude)
	////if err != nil {
	////	log.Fatalln(err)
	////}
	//
	////for _, e := range temperature {
	////	err := temp.Save(context.Background(), e)
	////	//fmt.Println("save data", string(e.Data))
	////	if err != nil {
	////		log.Fatalln(err)
	////	}
	////}
	//
	//byCity, err := temp.FindAllByCity(context.Background())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(byCity)
	//
	//fmt.Println(time.Unix(1666526400, 0))
	//fmt.Println(time.Unix(1666537200, 0))
	//fmt.Println(time.Unix(1666548000, 0))
	//fmt.Println(time.Unix(1666558800, 0))
	//fmt.Println(time.Unix(1666569600, 0))
}
