type state: {
    a: integer
    b: integer
    x: integer
} = {
    a: 0
    b: 0
    x: 0
}

type phrases: []{lang: string, phrases: []string}

action main on $state: [
    x = 3
    a = ($five + x) * 2
    if true:
        bind a->c, b<-c on {c: integer, d: integer} = $init:
            c = $multiply {a: c, b: d}
]

type phrases: {
    lang: string
    phrases: []string
}

action say on $phrases: []

action main2 on $state: [
    if !topic.fetchProfile && config.view && config.universe && config.profiles.fetched == "one": [
        topic.fetchProfile = true
        bind on $phrases = {lang: "en", phrases: [ "Found profile {{.config.profiles.selected.name}}." ] }: $say
    ]
]

constant init: {d: 2}
constant five: 5

function multiply on {a: integer, b: integer}: a * b
function fib on integer: (a > 0) ? (a + $fib (a - 1)) : 0
