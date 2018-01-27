package main

import (
    "testing"
)

func Test_IntEquality(t *testing.T) {
    z1, z2 := Int(0), Int(0)
    pz1, pz2 := &z1, &z2

    if !z1.Equals(z2) || !z1.Equals(pz2) || !pz1.Equals(z2) || !pz1.Equals(pz2) {
        t.Error("Expected zeroes to be equal")
    }

    if z1.Equals(nil) || pz1.Equals(nil) {
        t.Error("Expected zero not to be equal to nil")
    }

    x := Int(1)
    px := &x
    if x.Equals(z1) || x.Equals(pz1) || px.Equals(z1) || px.Equals(pz1) {
        t.Error("Expected one and zero not to be equal")
    }
}

func Test_PIntEquality(t *testing.T) {
    z1, z2 := PInt(0), PInt(0)
    pz1, pz2 := &z1, &z2

    if !z1.Equals(z2) || !z1.Equals(pz2) || !pz1.Equals(z2) || !pz1.Equals(pz2) {
        t.Error("Expected zeroes to be equal")
    }

    if z1.Equals(nil) || pz1.Equals(nil) {
        t.Error("Expected zero not to be equal to nil")
    }

    x := PInt(1)
    px := &x
    if x.Equals(z1) || x.Equals(pz1) || px.Equals(z1) || px.Equals(pz1) {
        t.Error("Expected one and zero not to be equal")
    }
}