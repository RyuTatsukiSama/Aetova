#pragma once

#ifdef SHARED
#ifdef BUILDING
#define VIEW_EXPORT __decmspec(dllexport)
#else
#define VIEW_EXPORT __decmspec(dllimpoty)
#endif
#else
#define VIEW_EXPORT
#endif