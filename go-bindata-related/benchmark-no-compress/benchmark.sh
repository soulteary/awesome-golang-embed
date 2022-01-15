#!/bin/bash

go test -run=TestSmallFileRepeatRequest -benchmem -memprofile mem-small.out -cpuprofile cpu-small.out -v
go test -run=TestLargeFileRepeatRequest -benchmem -memprofile mem-large.out -cpuprofile cpu-large.out -v
