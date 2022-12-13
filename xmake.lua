--[[
    Build script for xmake.
]]--

-- SECTION CONFIGURABLE_PATHS

external_home = "/usr"

-- SECTION TARGETS

target("libcode")
    set_kind("static")
    add_files("src/lib/*.cpp")
    set_languages("cxx20")

target("bincode_artifact")
    set_kind("binary")
    add_files("src/bin/*.cpp")
    add_deps("libcode")
    set_languages("cxx20")

target("test")
    set_kind("binary")
    add_files("src/run_tests.cpp")
    add_deps("libcode")
    set_languages("cxx20")
    add_links("CppUTest")
    add_linkdirs(external_home .. "/lib")
    add_includedirs(external_home .. "/include")
