#!/bin/bash

mkdir $1
cd $1
echo '''package main

import "fmt"

func main() {
    fmt.Println()
}''' > main.go