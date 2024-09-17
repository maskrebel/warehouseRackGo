package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Slot struct {
	SKU        string
	ExpireDate time.Time
}

type WarehouseRack struct {
	total int
	slots []*Slot
}

func NewWarehouseRack(total int) *WarehouseRack {
	fmt.Printf("Created a warehouse rack with %d slots.\n", total)
	return &WarehouseRack{
		total: total,
		slots: make([]*Slot, total),
	}
}

func (wr WarehouseRack) Rack(sku string, expireDate time.Time) {
	for i, slot := range wr.slots {
		if slot == nil {
			wr.slots[i] = &Slot{
				SKU:        sku,
				ExpireDate: expireDate,
			}
			fmt.Println("Allocated slot number: " + strconv.Itoa(i+1))
			return
		}
	}
	fmt.Println("Sorry, rack is full")
	return
}

func (wr WarehouseRack) rackOut(slotNumber int) {
	wr.slots[slotNumber-1] = nil
	fmt.Printf("Slot number %d is free\n", slotNumber)
}

func (wr WarehouseRack) status() {
	status := "Slot No.\tSKU No.\t\tExp Date\n"
	for idx, slot := range wr.slots {
		if slot != nil {
			status += fmt.Sprintf("%d\t\t\t%s\t\t\t%s\n", idx+1, slot.SKU, slot.ExpireDate)
		}
	}
	fmt.Println(status)
}

func (wr WarehouseRack) skuNumbersForProductWithExpDate(expDate time.Time) {
	result := ""
	for _, slot := range wr.slots {
		if slot != nil && slot.ExpireDate == expDate {
			if result == "" {
				result += slot.SKU
			} else {
				result += ", " + slot.SKU
			}
		}
	}
	fmt.Println(result)
}

func (wr WarehouseRack) slotNumbersForProductWithExpDate(expDate time.Time) {
	result := ""
	for idx, slot := range wr.slots {
		if slot != nil && slot.ExpireDate == expDate {
			if result == "" {
				result += strconv.Itoa(idx + 1)
			} else {
				result += ", " + strconv.Itoa(idx+1)
			}
		}
	}
	fmt.Println(result)
}

func (wr WarehouseRack) slotNumberForSKUNumber(sku string) {
	resString := "Not Found"
	for idx, slot := range wr.slots {
		if slot != nil && slot.SKU == sku {
			slotNumber := idx + 1
			fmt.Println(slotNumber)
			return
		}
	}
	fmt.Println(resString)
}

func preprocessing(commands string, rack WarehouseRack) WarehouseRack {
	action := strings.Split(commands, " ")[0]
	if strings.Compare("create_rack", action) == 0 || strings.Compare("create_warehouse_rack", action) == 0 {
		order := strings.Split(commands, " ")[1]
		total, _ := strconv.Atoi(order)
		return *NewWarehouseRack(total)
	}

	if rack.total == 0 {
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
		rack.rackOut(slotNumber)
	}

	if strings.Compare("status", action) == 0 {
		rack.status()
	}

	if strings.Compare("sku_numbers_for_product_with_exp_date", action) == 0 {
		expDate, err := time.Parse("2006-01-02", strings.Split(commands, " ")[1])
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rack.skuNumbersForProductWithExpDate(expDate)
	}

	if strings.Compare("slot_numbers_for_product_with_exp_date", action) == 0 {
		expDate, err := time.Parse("2006-01-02", strings.Split(commands, " ")[1])
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rack.slotNumbersForProductWithExpDate(expDate)
	}

	if strings.Compare("slot_number_for_sku_number", action) == 0 {
		sku := strings.Split(commands, " ")[1]
		rack.slotNumberForSKUNumber(sku)
	}

	return rack
}

func main() {
	rack := WarehouseRack{}
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
