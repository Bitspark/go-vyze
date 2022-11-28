pipe item1 on item -> id

pipe item2 on item -> {id, name}

pipe item3 on item -> {id, otherName: name}

pipe item4 on item -> {
    id
    name
    prices: price -> [] {
        price
        distributor -> {id, name}
    }
}

pipe distributorPrice1 on distributorPrice -> {
    price
    p1: price
    p2: price -> price
    p3: price -> price -> value
}

pipe distributorPrice2 on distributor_price -> {
    price
    distributor
    item
}

pipe distributorPrice3 on distributor_price -> {
    price
    distributor -> {id, name}
    item
}

pipe distributorPrice4 on distributor_price -> price


