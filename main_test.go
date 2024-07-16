package main

import "testing"

func TestSum(t *testing.T) {
    ans := handle("4 + 2")
    if ans != 6 {
        t.Errorf("Should be 6")
    }
}

func TestMutiply(t *testing.T) {
    ans := handle("4 * 2")
    if ans != 8 {
        t.Errorf("Should be 8")
    }
}


func TestDivide(t *testing.T) {
    ans := handle("4 / 2")
    if ans != 2 {
        t.Errorf("Should be 2")
    }
}

func TestReduce(t *testing.T) {
    ans := handle("0 - 2")
    if ans != -2 {
        t.Errorf("Should be -2")
    }
}

