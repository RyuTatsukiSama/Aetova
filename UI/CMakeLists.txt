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
    IMPORTED_IMPLIB_DEBUG "${CMAKE_SOURCE_DIR}/Dependencies/docLogger_v1-1-1/lib/Debug/docLogger.lib"
    IMPORTED_IMPLIB_RELEASE "${CMAKE_SOURCE_DIR}/Dependencies/docLogger_v1-1-1/lib/Release/docLogger.lib"
    IMPORTED_IMPLIB "${CMAKE_SOURCE_DIR}/Dependencies/docLogger_v1-1-1/lib/Debug/docLogger.lib"
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

set(SOURCES
    "${CMAKE_SOURCE_DIR}/Scripts/main.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/Common/common.h"
    "${CMAKE_SOURCE_DIR}/Scripts/Common/common.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/GameProcess/GameLauncher/gamelauncher.h"
    "${CMAKE_SOURCE_DIR}/Scripts/GameProcess/GameLauncher/gamelauncher.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/GameProcess/GameThread/gamethread.h"
    "${CMAKE_SOURCE_DIR}/Scripts/GameProcess/GameThread/gamethread.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/gamewindow.h"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/gamewindow.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/BackgroundWidget/backgroundwidget.h"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/BackgroundWidget/backgroundwidget.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/ButtonGame/buttongame.h"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/ButtonGame/buttongame.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/HTMLLabel/htmllabel.h"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/HTMLLabel/htmllabel.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/HTMLLabel/LabelTeam/labelteam.h"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/HTMLLabel/LabelTeam/labelteam.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/SplashLabel/splashlabel.h"
    "${CMAKE_SOURCE_DIR}/Scripts/GameWindow/SplashLabel/splashlabel.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/ClientBridge/client.h"
    "${CMAKE_SOURCE_DIR}/Scripts/ClientBridge/monitoring_data.h"
    "${CMAKE_SOURCE_DIR}/Scripts/ClientBridge/client.cpp"
    "${CMAKE_SOURCE_DIR}/Scripts/ClientBridge/WSReader/WSReader.h"
    "${CMAKE_SOURCE_DIR}/Scripts/ClientBridge/WSReader/WSReader.cpp"
)

set(RESSOURCES
    "${CMAKE_SOURCE_DIR}/Resources/html.qrc"
    "${CMAKE_SOURCE_DIR}/Resources/resources.qrc"
    "${CMAKE_SOURCE_DIR}/Resources/ui.qrc"
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
    Qt6::WebEngineWidgets
    nlohmann_json::nlohmann_json
    docLogger
)

set_target_properties(Aetova PROPERTIES
    WIN32_EXECUTABLE OFF
)

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