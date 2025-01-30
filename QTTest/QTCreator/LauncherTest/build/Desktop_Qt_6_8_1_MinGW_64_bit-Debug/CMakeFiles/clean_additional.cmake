# Additional clean files
cmake_minimum_required(VERSION 3.16)

if("${CONFIG}" STREQUAL "" OR "${CONFIG}" STREQUAL "Debug")
  file(REMOVE_RECURSE
  "CMakeFiles\\LauncherTest_autogen.dir\\AutogenUsed.txt"
  "CMakeFiles\\LauncherTest_autogen.dir\\ParseCache.txt"
  "LauncherTest_autogen"
  )
endif()
