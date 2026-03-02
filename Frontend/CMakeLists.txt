cmake_minimum_required(VERSION 3.31)
project(Aetova LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# nlohmann/json add
set(FETCHCONTENT_BASE_DIR "${CMAKE_SOURCE_DIR}/.cache" CACHE PATH "FetchContent base dir")
 
include(FetchContent)

set(JSON_SRC_DIR "${CMAKE_SOURCE_DIR}/.cache/json-src")
if(EXISTS "${JSON_SRC_DIR}/CMakeLists.txt")
    set(FETCHCONTENT_FULLY_DISCONNECTED ON)
else()
    set(FETCHCONTENT_FULLY_DISCONNECTED OFF)
endif()

FetchContent_Declare(JSON
  GIT_REPOSITORY https://github.com/nlohmann/json.git
  GIT_TAG v3.11.3 
  GIT_SHALLOW ON
  EXCLUDE_FROM_ALL
  SYSTEM
  SOURCE_DIR "${JSON_SRC_DIR}"
)
FetchContent_MakeAvailable(JSON)

# doc Logger
add_library(docLogger SHARED IMPORTED)

set_target_properties(docLogger PROPERTIES
    IMPORTED_IMPLIB "${CMAKE_SOURCE_DIR}/Dependencies/docLogger_v1-1-1/lib/docLogger.lib"
)

target_include_directories(docLogger INTERFACE
    "${CMAKE_SOURCE_DIR}/Dependencies/docLogger_v1-1-1/include"
)

# Qt setup
find_package(Qt6 REQUIRED COMPONENTS Widgets)
find_package(Qt6 REQUIRED COMPONENTS Network)
find_package(Qt6 REQUIRED COMPONENTS WebSockets)
find_package(Qt6 REQUIRED COMPONENTS WebEngineWidgets)

qt_standard_project_setup()

set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTORCC ON)

set(QT_FORCE_CMP0156_TO_VALUE NEW)

set(app_icon_resource_windows "${CMAKE_SOURCE_DIR}/Resources/Assets/Icons/Logo.rc")

# Common deps

add_library(common_deps INTERFACE)

target_link_libraries(common_deps INTERFACE
    Qt6::Widgets
    Qt6::Core
    Qt6::Gui
    Qt6::Network
    Qt6::WebSockets
    Qt6::WebEngineWidgets
    nlohmann_json::nlohmann_json
    docLogger
)

# View

add_library(View STATIC
    "${CMAKE_SOURCE_DIR}/Sources/View/src/New/WebView.cpp"
    "${CMAKE_SOURCE_DIR}/Sources/View/src/New/AetovaWindow.cpp"
)

target_link_libraries(View PUBLIC common_deps)

target_include_directories(View PUBLIC 
    "${CMAKE_SOURCE_DIR}/Sources/View/include"
)

# App

set(SOURCES 
    "${CMAKE_SOURCE_DIR}/Sources/App/src/main.cpp"
    "${CMAKE_SOURCE_DIR}/Sources/App/src/Common/common.h"
    "${CMAKE_SOURCE_DIR}/Sources/App/src/Common/common.cpp"
)

set(RESSOURCES
    "${CMAKE_SOURCE_DIR}/Resources/Web/DGP/DGP.qrc"
)

# Exe
qt_add_executable(Aetova
    ${SOURCES}
    ${RESSOURCES}
    ${app_icon_resource_windows}
)

target_link_libraries(Aetova PRIVATE 
    View
)

set_target_properties(Aetova PROPERTIES
    WIN32_EXECUTABLE OFF
)

# ---- COMMANDS ----

add_custom_command(TARGET Aetova POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E make_directory
        "$<TARGET_FILE_DIR:Aetova>/wd"
)

add_custom_command(TARGET Aetova POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${CMAKE_SOURCE_DIR}/Dependencies/docLogger_v1-1-1/dll/$<CONFIG>/docLogger.dll"
        "$<TARGET_FILE_DIR:Aetova>/docLogger.dll"    
)

add_custom_command(TARGET Aetova POST_BUILD
    COMMAND ${Qt6_DIR}/../../../bin/windeployqt.exe
        --debug
        $<TARGET_FILE:Aetova>
)