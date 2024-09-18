package controllers

import (
	"fmt"
	"strconv"
	"time"
)

type Slot struct {
	SKU        string
	ExpireDate time.Time
}

type WarehouseRack struct {
	Total int
	slots []*Slot
}

func NewWarehouseRack(total int) WarehouseRack {
	fmt.Printf("Created a warehouse rack with %d slots.\n", total)
	return WarehouseRack{
		Total: total,
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

func (wr WarehouseRack) RackOut(slotNumber int) {
	wr.slots[slotNumber-1] = nil
	fmt.Printf("Slot number %d is free\n", slotNumber)
}

func (wr WarehouseRack) Status() {
	status := "Slot No.\tSKU No.\t\tExp Date\n"
	for idx, slot := range wr.slots {
		if slot != nil {
			status += fmt.Sprintf("%d\t\t\t%s\t\t\t%s\n", idx+1, slot.SKU, slot.ExpireDate)
		}
	}
	fmt.Println(status)
}

func (wr WarehouseRack) SkuNumbersForProductWithExpDate(expDate time.Time) {
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

func (wr WarehouseRack) SlotNumbersForProductWithExpDate(expDate time.Time) {
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

func (wr WarehouseRack) SlotNumberForSKUNumber(sku string) {
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
