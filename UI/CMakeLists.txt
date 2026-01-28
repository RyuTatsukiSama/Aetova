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

FetchContent_Declare(
  JSON
  GIT_REPOSITORY https://github.com/nlohmann/json.git
  GIT_TAG v3.11.3 
)

FetchContent_MakeAvailable(JSON)

# Qt setup
find_package(Qt6 REQUIRED COMPONENTS Widgets)
find_package(Qt6 REQUIRED COMPONENTS Network)
find_package(Qt6 REQUIRED COMPONENTS WebSockets)

qt_standard_project_setup()

set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTORCC ON)

set(QT_FORCE_CMP0156_TO_VALUE NEW)

set(SOURCES
    "./Scripts/main.cpp"
    "./Scripts/Common/common.h"
    "./Scripts/Common/common.cpp"
    "./Scripts/GameProcess/GameLauncher/gamelauncher.h"
    "./Scripts/GameProcess/GameLauncher/gamelauncher.cpp"
    "./Scripts/GameProcess/GameThread/gamethread.h"
    "./Scripts/GameProcess/GameThread/gamethread.cpp"
    "./Scripts/GameWindow/gamewindow.h"
    "./Scripts/GameWindow/gamewindow.cpp"
    "./Scripts/GameWindow/BackgroundWidget/backgroundwidget.h"
    "./Scripts/GameWindow/BackgroundWidget/backgroundwidget.cpp"
    "./Scripts/GameWindow/ButtonGame/buttongame.h"
    "./Scripts/GameWindow/ButtonGame/buttongame.cpp"
    "./Scripts/GameWindow/HTMLLabel/htmllabel.h"
    "./Scripts/GameWindow/HTMLLabel/htmllabel.cpp"
    "./Scripts/GameWindow/HTMLLabel/LabelTeam/labelteam.h"
    "./Scripts/GameWindow/HTMLLabel/LabelTeam/labelteam.cpp"
    "./Scripts/GameWindow/SplashLabel/splashlabel.h"
    "./Scripts/GameWindow/SplashLabel/splashlabel.cpp"
    "./Scripts/ClientBridge/client.h"
    "./Scripts/ClientBridge/client_struct.h"
    "./Scripts/ClientBridge/client.cpp"
)

set(RESSOURCES
    "./Resources/html.qrc"
    "./Resources/resources.qrc"
    "./Resources/ui.qrc"
)

 set(app_icon_resource_windows "${CMAKE_CURRENT_SOURCE_DIR}/Resources/Logo/Logo.rc")

qt_add_executable(Aetova
    ${SOURCES}
    ${RESSOURCES}
    ${app_icon_resource_windows}
)

target_link_libraries(Aetova PRIVATE 
    Qt6::Widgets
    Qt6::Core
    Qt6::Gui
    Qt6::Network
    Qt6::WebSockets
    nlohmann_json::nlohmann_json
)

# add the docLogger
target_link_libraries(Aetova PRIVATE "./lib/docLogger")

target_include_directories(Aetova PRIVATE "./Dependencies/docLogger_v1-1-1/include")

set_target_properties(Aetova PROPERTIES
    WIN32_EXECUTABLE OFF
)

add_custom_command(TARGET Aetova POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E make_directory
        "$<TARGET_FILE_DIR:Aetova>/wd"
)

# Copy the lib for VSS
add_custom_command(TARGET Aetova PRE_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_directory_if_different
        "${CMAKE_SOURCE_DIR}/Dependencies/docLogger_v1-1-1/lib/$<CONFIG>"
        "$<TARGET_FILE_DIR:Aetova>/lib"
)

# Copy the lib for VS Code
add_custom_command(TARGET Aetova PRE_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_directory_if_different
        "${CMAKE_SOURCE_DIR}/Dependencies/docLogger_v1-1-1/lib/$<CONFIG>"
        "$<TARGET_FILE_DIR:Aetova>/../lib"
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