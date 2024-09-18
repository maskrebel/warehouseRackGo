package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"warehouse/controllers"
)

func preprocessing(commands string, rack controllers.WarehouseRack) controllers.WarehouseRack {
	action := strings.Split(commands, " ")[0]
	if strings.Compare("create_rack", action) == 0 || strings.Compare("create_warehouse_rack", action) == 0 {
		order := strings.Split(commands, " ")[1]
		total, _ := strconv.Atoi(order)
		return controllers.NewWarehouseRack(total)
	}

	if rack.Total == 0 {
		fmt.Println("Please create a rack first using 'create_rack'.")
		return rack
	}

	if strings.Compare("rack", action) == 0 {
		sku := strings.Split(commands, " ")[1]
		expDate, err := time.Parse("2006-01-02", strings.Split(commands, " ")[2])
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rack.Rack(sku, expDate)
	}

	if strings.Compare("rack_out", action) == 0 {
		slotNumber, _ := strconv.Atoi(strings.Split(commands, " ")[1])
		rack.RackOut(slotNumber)
	}

	if strings.Compare("status", action) == 0 {
		rack.Status()
	}

	if strings.Compare("sku_numbers_for_product_with_exp_date", action) == 0 {
		expDate, err := time.Parse("2006-01-02", strings.Split(commands, " ")[1])
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rack.SkuNumbersForProductWithExpDate(expDate)
	}

	if strings.Compare("slot_numbers_for_product_with_exp_date", action) == 0 {
		expDate, err := time.Parse("2006-01-02", strings.Split(commands, " ")[1])
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rack.SlotNumbersForProductWithExpDate(expDate)
	}

	if strings.Compare("slot_number_for_sku_number", action) == 0 {
		sku := strings.Split(commands, " ")[1]
		rack.SlotNumberForSKUNumber(sku)
	}

	return rack
}

func main() {
	rack := controllers.WarehouseRack{}
	if len(os.Args) == 2 {
		fileName := os.Args[1]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			rack = preprocessing(line, rack)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("-> ")
			commands, _ := reader.ReadString('\n')
			commands = strings.Replace(commands, "\n", "", -1)

			if strings.Compare("exit", commands) == 0 {
				break
			}

			rack = preprocessing(commands, rack)
		}
	}
}
