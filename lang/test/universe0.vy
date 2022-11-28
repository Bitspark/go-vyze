universe vergleichsportal {

    model object extends base.object {
    }

    model item extends object {
        item_id -> @data.string
        distributor_price
    }

    model distributor_price extends object {

    }

    model distributor extends object {

    }

    pipe distributor on distributor -> {
        id
        name
    }

}
