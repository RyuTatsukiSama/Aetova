cmake_minimum_required(VERSION 3.31)

project(Aetova VERSION 1.0.0 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(Qt6 REQUIRED COMPONENTS Widgets)

qt_standard_project_setup()

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
    "./Resources/html.cpp"
    "./Resources/resources.cpp"
    "./Resources/ui.cpp"
)

qt_add_executable(Aetova
    ${SOURCES}
)

target_link_libraries(Aetova PRIVATE 
    Qt6::Widgets
    Qt6::Core
    Qt6::Gui
)

# add the docLogger
target_link_libraries(Aetova PRIVATE "./lib/docLogger")

target_include_directories(Aetova PRIVATE "./Dependencies/docLogger_v1-1-1/include")

set_target_properties(Aetova PROPERTIES
    WIN32_EXECUTABLE OFF
)

add_custom_command(TARGET Aetova POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_directory_if_different
        "${CMAKE_SOURCE_DIR}/apps"
        "$<TARGET_FILE_DIR:Aetova>/apps"
)

add_custom_command(TARGET Aetova PRE_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_directory_if_different
        "${CMAKE_SOURCE_DIR}/Dependencies/docLogger_v1-1-1/lib"
        "${CMAKE_SOURCE_DIR}/build/lib"
)

add_custom_command(TARGET Aetova POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${CMAKE_SOURCE_DIR}/Dependencies/docLogger_v1-1-1/dll/$<CONFIG>/docLogger.dll"
        "$<TARGET_FILE_DIR:Aetova>/docLogger.dll"    
)