#ifndef CLIENT_STRUCT_H
#define CLIENT_STRUCT_H

#include <string>
#include <nlohmann/json.hpp>
using json = nlohmann::json;

enum MessageType
{
    TEXT,
    CLOSE,
    EXIT,
    MONITORING
};

typedef struct Message 
{
    MessageType type;
    json data;
};


typedef struct MonitoringData
{
    float dlPrc;
    float dlSpeed;
    float wrPrc;
    float wrSpeed;
};


#endif