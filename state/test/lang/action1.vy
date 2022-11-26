action action1 on {a: string}: a = "test"

action action2 on {a: string}: [
    a = "test1"
    a = "test2"
]

action action3 on {a: string}: {
    a = "test1"
    a = "test2"
}

action action4 on {a: string}:
    if a == "test":
        a = "test2"

action action5 on {a: string}:
    if a == "test":
        a = "test2"
    else:
        a = "test3"
