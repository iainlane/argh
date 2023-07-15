cmake_minimum_required(VERSION 3.10)

project(SimpleDanceSolver VERSION 0.0 LANGUAGES C)

list(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_SOURCE_DIR}/cmake")

find_program(GO_EXECUTABLE go)
if (GO_EXECUTABLE-NOTFOUND)
    message(FATAL_ERROR "Go not found! Please install Go.")
endif()

# verbose output
set(CMAKE_VERBOSE_MAKEFILE ON)

set(GO_OUTPUT_A ${CMAKE_BINARY_DIR}/gofunction.a)
set(GO_OUTPUT_H ${CMAKE_BINARY_DIR}/gofunction.h)

add_custom_command(
    OUTPUT ${GO_OUTPUT_A} ${GO_OUTPUT_H}
    COMMAND ${GO_EXECUTABLE} build -a -buildmode=c-archive -o ${GO_OUTPUT_A}
    WORKING_DIRECTORY ${CMAKE_SOURCE_DIR}/gofunction/
    COMMENT "Building gofunction binding"
    DEPENDS gofunction/gofunction.go
    VERBATIM
)

add_custom_target(go_package DEPENDS ${GO_OUTPUT_A} ${GO_OUTPUT_H})

include_directories(${CMAKE_SOURCE_DIR}/c)
include_directories(${CMAKE_CURRENT_BINARY_DIR})

# add -Wl,-Bsymbolic to avoid multiple Go runtime instances
set(CMAKE_SHARED_LINKER_FLAGS "${CMAKE_SHARED_LINKER_FLAGS} -Wl,-Bsymbolic-functions")

add_library(myc STATIC c/callgofunction.c)
add_dependencies(myc go_package)
target_link_libraries(myc ${GO_OUTPUT_A})

add_library(mycso SHARED c/callgofunction.c)
add_dependencies(mycso go_package)
target_link_libraries(mycso ${GO_OUTPUT_A})

add_subdirectory(tests)
