# View

if(BUILD_SHARED)
    add_library(View SHARED
        "${CMAKE_SOURCE_DIR}/Sources/View/src/New/WebView.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/View/src/New/AetovaWindow.cpp"
    )
else()
    add_library(View OBJECT
        "${CMAKE_SOURCE_DIR}/Sources/View/src/New/WebView.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/View/src/New/AetovaWindow.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/View/include/New/WebView.h"
        "${CMAKE_SOURCE_DIR}/Sources/View/include/New/AetovaWindow.h"
    )
endif()

target_link_libraries(View PUBLIC common_deps)

target_include_directories(View PUBLIC 
    "${CMAKE_SOURCE_DIR}/Sources/View/include"
)

set_target_properties(View PROPERTIES
    AUTOMOC ON
    AUTOUIC ON
    AUTORCC ON
)

# for dll export

if(BUILD_SHARED)
    target_compile_definitions(View PRIVATE BUILDING SHARED)
    target_compile_definitions(View INTERFACE SHARED)
endif()