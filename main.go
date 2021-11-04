package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/victor0198utm/restaurant_hall/appData"
	"github.com/victor0198utm/restaurant_hall/models"
)

var w sync.WaitGroup

var m sync.Mutex

var clients int = 0
var client_id int = 1

func make_clients() {
	for {
		time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
		if clients < 2 {
			w.Add(1)
			go client()
		}
	}
	w.Done()
}

func client() {
	m.Lock()
	clients = clients + 1
	this_client_id := client_id
	client_id = client_id + 1
	m.Unlock()

	// get menu (request)
	resp, err := http.Get("http://network_food_ordering_1:8011/menu")
	if err != nil {
		log.Fatal(err)
	}

	menus := models.Menus{}
	err = json.NewDecoder(resp.Body).Decode(&menus)
	if err != nil {
		fmt.Println(err.Error(), http.StatusBadRequest)
		w.Done()
		return
	}

	fmt.Println("Menus:", menus)
	if menus.Restaurants == 0 {
		w.Done()
		return
	}

	// choose dishes from menus
	n_dishes := rand.Intn(3) + 1
	restaurant_ids := []int{}
	dishes_ids := []int{}
	for i := 0; i < n_dishes; i++ {
		r_id := rand.Intn(menus.Restaurants)
		restaurant_ids = append(restaurant_ids, r_id)

		// menus.Restaurants_data[r_id].Menu_items
		d_id := rand.Intn(menus.Restaurants_data[r_id].Menu_items) + 1
		dishes_ids = append(dishes_ids, d_id)
	}

	fmt.Println("restaurant_ids:", restaurant_ids, " dishes:", dishes_ids)

	// send order (request)
	orders := []models.OrderReq{}

	for i := 1; i <= menus.Restaurants; i++ {
		items := []int{}
		for j := 0; j < len(restaurant_ids); j++ {
			if restaurant_ids[j]+1 == i {
				items = append(items, dishes_ids[j])
			}
		}

		max_time := 0
		for _, dish_id := range items {

			//fmt.Println(i, the_dish)
			prepation_time := appData.GetDish(dish_id - 1).Preparation_time
			if max_time < prepation_time {
				max_time = prepation_time
			}
		}

		orders = append(orders, models.OrderReq{
			i,
			items,
			rand.Intn(4) + 1,
			int(float32(max_time) * 1.3),
			int(time.Now().Unix()),
		})
	}

	complete_order := models.ClientOrderReq{
		this_client_id,
		orders,
	}

	fmt.Println("COMPLETE ORDER:", complete_order)

	json_data, err_marshall := json.Marshal(complete_order)
	if err_marshall != nil {
		log.Fatal(err_marshall)
	}
	resp, err = http.Post("http://network_food_ordering_1:8011/order", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	// get response
	clientOrderResp := models.ClientOrderResp{}
	err = json.NewDecoder(resp.Body).Decode(&clientOrderResp)
	if err != nil {
		fmt.Println(err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("GET RESPOSE:", clientOrderResp)

	// wait for dishes

	orders_to_receive := clientOrderResp.Orders

	for {
		time.Sleep(1 * time.Second)
		// check if orders are ready (request)

		for i := 0; i < len(orders_to_receive); i++ {
			url := "http://" + orders_to_receive[i].Restaurant_address + "/v2/order/" + strconv.Itoa(orders_to_receive[i].Order_id)
			fmt.Println("Trying to get order:", url)
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			orderReceived := models.OrderRegistered{}
			err = json.NewDecoder(resp.Body).Decode(&orderReceived)
			if err != nil {
				fmt.Println(err.Error(), http.StatusBadRequest)
			}

			fmt.Println("ID:", orders_to_receive[i].Order_id, "|", orderReceived)

			if orderReceived.Is_ready == true {
				fmt.Println("Ready:", orderReceived)

				make_review(orders_to_receive[i].Restaurant_id, orderReceived)

				orders_to_receive = RemoveOrderResp(orders_to_receive, i)
				i = 0
			}
		}

		if len(orders_to_receive) == 0 {
			break
		}
	}

	time.Sleep(10 * time.Second)

	m.Lock()
	clients = clients - 1
	m.Unlock()
	w.Done()
}

func RemoveOrderResp(s []models.OrderResp, index int) []models.OrderResp {
	return append(s[:index], s[index+1:]...)
}

func make_review(restaurant_id int, receivedOrder models.OrderRegistered) {
	// compute ratings

	stars := 0
	if receivedOrder.Cooking_time < receivedOrder.Max_wait {
		stars = 5
	} else if float64(receivedOrder.Cooking_time) < float64(receivedOrder.Max_wait)*1.1 {
		stars = 4
	} else if float64(receivedOrder.Cooking_time) < float64(receivedOrder.Max_wait)*1.2 {
		stars = 3
	} else if float64(receivedOrder.Cooking_time) < float64(receivedOrder.Max_wait)*1.3 {
		stars = 2
	} else if float64(receivedOrder.Cooking_time) < float64(receivedOrder.Max_wait)*1.4 {
		stars = 1
	}

	// post rating to food ordering system
	review := models.Review{
		restaurant_id,
		stars,
	}
	json_data, err_marshall := json.Marshal(review)
	if err_marshall != nil {
		log.Fatal(err_marshall)
	}
	resp, err := http.Post("http://network_food_ordering_1:8011/rating", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Review sent. Stars: %d. Status: %d\n\n", stars, resp.StatusCode)
}

func display() {
	for {
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("Clients:", clients)
	}
	w.Done()
}

func main() {

	w.Add(1)
	go display()

	w.Add(1)
	go make_clients()

	w.Wait()
}
