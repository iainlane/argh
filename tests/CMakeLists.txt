include(CTest)
find_package(Check REQUIRED)

include_directories(${CHECK_INCLUDE_DIRS})
add_executable(c_api_tests c_api_tests.c)
target_link_libraries(c_api_tests ${CHECK_LIBRARIES} myc)

add_test(tests_c_api tests_c_api)
