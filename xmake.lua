--[[
    Build script for xmake.
]]--

-- SECTION CONFIGURABLE_PATHS

external_home = "/usr"

-- SECTION COMPILER_CONFIG

set_languages("cxx20")

-- SECTION TARGETS

target("bincode")
    set_kind("binary")
    add_files("src/*.cpp", "src/lib/*.cpp")
    add_links("CppUTest")
    add_linkdirs(external_home .. "/lib")
    add_includedirs(external_home .. "/include")
