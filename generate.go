package main

import "math/rand"

func messageGenerator(list []string) string {
	return list[rand.Intn(len(list))]
}
