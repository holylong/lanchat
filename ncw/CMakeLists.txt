cmake_minimum_required(VERSION 3.11)

project(netcat++)


file(GLOB ncpp_SRC ${CMAKE_CURRENT_SOURCE_DIR}/doexec.c
${CMAKE_CURRENT_SOURCE_DIR}/getopt.c
${CMAKE_CURRENT_SOURCE_DIR}/netcat.c
${CMAKE_CURRENT_SOURCE_DIR}/getopt.h
${CMAKE_CURRENT_SOURCE_DIR}/generic.h)

add_executable(nc++ ${ncpp_SRC})