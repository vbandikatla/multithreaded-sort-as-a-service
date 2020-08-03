package controllers

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"sync"
)

func (server *Server) Get() {
	list := server.Ctx.Input.Param(":list")
	
	server.Data["nums"] = list

	nums := []float64{}
	list_split := strings.Split(list, ",")

	for _, str := range list_split {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			server.Data["result"] = "no idea on how to sort !"
			return
		}
		nums = append(nums, num)
	}

	if (!isValid(nums)) {
		server.Data["result"] = "validation error!"
		return
	}

	nums = sort(nums)
	for _, num := range nums {
		fmt.Println(num)
	}

	var resultBuilder string

	for _, num := range nums {
		resultBuilder += ", " + fmt.Sprintf("%f", num)
	}

	server.Data["result"] = resultBuilder[2:]
	server.TplName = "result.html"
}

func sort(nums []float64) []float64 {
	var wg sync.WaitGroup
	ch := make(chan float64, len(nums))

	min, max := getMinMax(nums)
	for _, num := range nums {
		wg.Add(1)
		go func(n float64) {
			defer wg.Done()

			normalized_time := (n - min)/(max - min)
			time.Sleep(time.Duration(normalized_time * 1000) * time.Millisecond)
			ch <- n
		} (num)
	}
	wg.Wait()
	close(ch)

	res := make([]float64, 0)
	for num := range ch {
		res = append(res, num)
	}
	//fmt.Println("here" + min + max)
	for _, num := range res {
		fmt.Println(num)
	}
	return res
}

func getMinMax(nums []float64) (float64, float64) {
	var min, max float64 = nums[0], nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}

	return min, max
}

func isValid(nums []float64) bool {
	return true
}