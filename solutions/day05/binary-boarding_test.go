package main

import "testing"

func TestFindRowColumn1(t *testing.T) {
	r, c := findRowColumn("BFFFBBFRRR")
	if r != 70 || c != 7 {
		t.Errorf("r, c is %d, %d; want 70, 7", r, c)
	}
}
func TestFindRowColumn2(t *testing.T) {
	r, c := findRowColumn("FFFBBBFRRR")
	if r != 14 || c != 7 {
		t.Errorf("r, c is %d, %d; want 14, 7", r, c)
	}
}

func TestFindRowColumn3(t *testing.T) {
	r, c := findRowColumn("BBBBBBBRRR")
	if r != 127 || c != 7 {
		t.Errorf("r, c is %d, %d; want 127, 7", r, c)
	}
}
func TestFindRowColumn4(t *testing.T) {
	r, c := findRowColumn("BBFFBBFRLL")
	if r != 102 || c != 4 {
		t.Errorf("r, c is %d, %d; want 102, 4", r, c)
	}
}

func TestSeatId(t *testing.T) {
	seatId := calculateSeatId(70, 7)
	if seatId != 567 {
		t.Errorf("seatId is %d; want 567", seatId)
	}
}
