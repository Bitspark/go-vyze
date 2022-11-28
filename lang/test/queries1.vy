pipe item1 on item -> id

pipe item2 on item -> {id, name}

pipe item3 on item -> {id, otherName: name}

pipe item4 on item -> {
    id
    name
    prices: <- distributorprice#price [] {
        price -> value
        distributor -> {id, name}
    }
}

pipe distributorPrice1 on distributorprice -> {
    price -> value
    p1: price -> value
}

pipe distributorPrice2 on distributorprice -> {
    price -> value
    distributor -> name
    item -> name
}

pipe distributorPrice3 on distributorprice -> {
    price -> value
    distributor -> {id, name}
    item -> name
}

pipe distributorPrice4 on distributorprice -> price -> value
