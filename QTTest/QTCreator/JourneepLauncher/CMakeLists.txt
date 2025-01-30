cmake_minimum_required(VERSION 3.16)

project(JourneepLauncher VERSION 0.1 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(Qt6 REQUIRED COMPONENTS Quick)

qt_standard_project_setup(REQUIRES 6.5)

qt_add_executable(appJourneepLauncher
    main.cpp
)

qt_add_qml_module(appJourneepLauncher
    URI JourneepLauncher
    VERSION 1.0
    QML_FILES
    main.qml
RESOURCES UntitledProject1/designer/plugin.metainfo UntitledProject1/CMakeLists.txt UntitledProject1/qmldir UntitledProject1Content/fonts/fonts.txt UntitledProject1Content/CMakeLists.txt qtquickcontrols2.conf UntitledProject1.qrc
QML_FILES UntitledProject1/Constants.qml UntitledProject1/EventListModel.qml UntitledProject1/EventListSimulator.qml UntitledProject1Content/App.qml UntitledProject1Content/Screen01.ui.qml UntitledProject1.qmlproject
)

qt_add_resources(my_project_resources "UntitledProject1.qrc")

# Qt for iOS sets MACOSX_BUNDLE_GUI_IDENTIFIER automatically since Qt 6.1.
# If you are developing for iOS or macOS you should consider setting an
# explicit, fixed bundle identifier manually though.
set_target_properties(appJourneepLauncher PROPERTIES
#    MACOSX_BUNDLE_GUI_IDENTIFIER com.example.appJourneepLauncher
    MACOSX_BUNDLE_BUNDLE_VERSION ${PROJECT_VERSION}
    MACOSX_BUNDLE_SHORT_VERSION_STRING ${PROJECT_VERSION_MAJOR}.${PROJECT_VERSION_MINOR}
    MACOSX_BUNDLE TRUE
    WIN32_EXECUTABLE TRUE
)

target_link_libraries(appJourneepLauncher
    PRIVATE Qt6::Quick
)

include(GNUInstallDirs)
install(TARGETS appJourneepLauncher
    BUNDLE DESTINATION .
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
    RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
)
