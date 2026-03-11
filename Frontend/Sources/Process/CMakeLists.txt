# Process

if(BUILD_SHARED)
    add_library(Process SHARED
        "${CMAKE_SOURCE_DIR}/Sources/Process/src/gamelauncher.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/Process/src/gamethread.cpp"
    )
else()
    add_library(Process OBJECT
        "${CMAKE_SOURCE_DIR}/Sources/Process/src/gamelauncher.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/Process/src/gamethread.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/Process/include/gamelauncher.h"
        "${CMAKE_SOURCE_DIR}/Sources/Process/include/gamethread.h"
    )
endif()

target_link_libraries(Process PUBLIC common_deps)

target_include_directories(Process PUBLIC 
    "${CMAKE_SOURCE_DIR}/Sources/Process/include"
)

set_target_properties(Process PROPERTIES
    AUTOMOC ON
    AUTOUIC ON
    AUTORCC ON
)

# for dll export

if(BUILD_SHARED)
    target_compile_definitions(Process PRIVATE BUILDING SHARED)
    target_compile_definitions(Process INTERFACE SHARED)
endif()