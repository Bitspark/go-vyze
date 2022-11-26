type complex1: {
    a: []string
    b: {
        c: raw
    }
}

type complex2: []{
    a: []$complex1
    b: []{
        u: boolean
        d: string
    }
}
