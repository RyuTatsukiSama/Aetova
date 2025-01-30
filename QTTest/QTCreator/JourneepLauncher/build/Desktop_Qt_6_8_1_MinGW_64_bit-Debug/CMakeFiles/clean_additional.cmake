# Additional clean files
cmake_minimum_required(VERSION 3.16)

if("${CONFIG}" STREQUAL "" OR "${CONFIG}" STREQUAL "Debug")
  file(REMOVE_RECURSE
  "CMakeFiles\\appJourneepLauncher_autogen.dir\\AutogenUsed.txt"
  "CMakeFiles\\appJourneepLauncher_autogen.dir\\ParseCache.txt"
  "appJourneepLauncher_autogen"
  )
endif()
