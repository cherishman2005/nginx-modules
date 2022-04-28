package main

import (
    "fmt"
    "time"
)

const (
    CommonCart = "common"
    BuyNowCart = "buyNow"
)

type cartExts struct {
    CartType string
    TTL      time.Duration
}

type CartExt interface {
    apply(*cartExts)
}

// 这里新增了类型，标记这个函数。相关技巧后面介绍
type tempFunc func(*cartExts)

// 实现 CartExt 接口
type funcCartExt struct {
    f tempFunc
}

// 实现的接口
func (fdo *funcCartExt) apply(e *cartExts) {
    fdo.f(e)
}

func newFuncCartExt(f tempFunc) *funcCartExt {
    return &funcCartExt{f: f}
}

type DemoCart struct {
    UserID string
    ItemID string
    Sku    int64
    Ext    cartExts
}

var DefaultExt = cartExts{
    CartType: CommonCart,       // 默认是普通购物车类型
    TTL:      time.Minute * 60, // 默认 60min 过期
}

func NewCart(userID string, Sku int64, exts ...CartExt) *DemoCart {
    c := &DemoCart{
        UserID: userID,
        Sku:    Sku,
        Ext:    DefaultExt, // 设置默认值
    }

    // 遍历进行设置
    for _, ext := range exts {
        ext.apply(&c.Ext)
    }

    return c
}

func WithCartType(cartType string) CartExt {
    return newFuncCartExt(func(exts *cartExts) {
        exts.CartType = cartType
    })
}

func WithTTL(d time.Duration) CartExt {
    return newFuncCartExt(func(exts *cartExts) {
        exts.TTL = d
    })
}

func main() {
    exts := []CartExt{
        WithCartType(CommonCart),
        WithTTL(1000),
    }

    cart := NewCart("dayu", 888, exts...)
    fmt.Printf("cart:%v\n", cart)
}
