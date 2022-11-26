type map1: {
    a: string
    b: boolean
} = {
    a: "test"
    b: true
} in [{a: "test", b: true}, {a: "test2", b: false}]

type video: {
     file: string
     name: string
     isDark: boolean
} in [
    { file: "ambient_01.mp4", name: "Bright Nebula", isDark: true }
    { file: "ambient_02.mp4", name: "Processor", isDark: true }
    { file: "ambient_03.mp4", name: "Purple Nebula", isDark: true }
    { file: "ambient_04.mp4", name: "Snow Lantern", isDark: false }
    { file: "ambient_05.mp4", name: "Glowing Particles", isDark: true }
    { file: "ambient_06.mp4", name: "Planet Earth", isDark: true }
    { file: "ambient_07.mp4", name: "Forest Drive", isDark: false }
    { file: "ambient_08.mp4", name: "Ocean Sunset", isDark: false }
    { file: "ambient_09.mp4", name: "AI Tube", isDark: true }
    { file: "ambient_10.mp4", name: "Forest Sunset", isDark: false }
]
