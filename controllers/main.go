package controllers

import (
	_ "github.com/gofiber/swagger"
	_ "github.com/swaggo/files" // swagger embed files
)

//@title go-auction
//@version 1.0
//description go-auction restapi
//@BasePath /api/v1

func AuthInit() *Auth {return &Auth{}}

func ProductInit() *Product {return &Product{}}

func AuctionInit() *Auction{ return &Auction{}}