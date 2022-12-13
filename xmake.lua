target("libcode")
    set_kind("static")
    add_files("src/lib/*.cpp")

target("bincode")
    set_kind("binary")
    add_files("src/bin/*.cpp")
    add_deps("libcode")
