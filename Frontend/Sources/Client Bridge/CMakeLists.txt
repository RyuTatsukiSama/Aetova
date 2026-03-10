# Client Bridge

if(BUILD_SHARED)
    add_library(ClientBridge SHARED
        "${CMAKE_SOURCE_DIR}/Sources/Client Bridge/src/WSReader.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/Client Bridge/src/client.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/Client Bridge/src/client_bridge.cpp"
    )
else()
    add_library(ClientBridge OBJECT
        "${CMAKE_SOURCE_DIR}/Sources/Client Bridge/src/WSReader.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/Client Bridge/src/client.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/Client Bridge/src/client_bridge.cpp"
        "${CMAKE_SOURCE_DIR}/Sources/Client Bridge/include/WSReader.h"
        "${CMAKE_SOURCE_DIR}/Sources/Client Bridge/include/client_bridge.h"
        "${CMAKE_SOURCE_DIR}/Sources/Client Bridge/include/client.h"
    )
endif()

target_link_libraries(ClientBridge PUBLIC common_deps)

target_include_directories(ClientBridge PUBLIC 
    "${CMAKE_SOURCE_DIR}/Sources/Client Bridge/include"
)

set_target_properties(ClientBridge PROPERTIES
    AUTOMOC ON
    AUTOUIC ON
    AUTORCC ON
)

# for dll export

if(BUILD_SHARED)
    target_compile_definitions(ClientBridge PRIVATE BUILDING SHARED)
    target_compile_definitions(ClientBridge INTERFACE SHARED)
endif()