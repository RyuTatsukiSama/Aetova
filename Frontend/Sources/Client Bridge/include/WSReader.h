#ifndef WSREADER_H
#define WSREADER_H

#include <variant>
#include <string>
#include "monitoring_data.h"
#include <nlohmann/json.hpp>
using json = nlohmann::json;

namespace doc
{
    class Logger;
}

enum MessageType
{
    TEXT,
    CLOSE,
    EXIT,
    MONITORING
};

class Message
{
    doc::Logger *log;

public:
    void readText();
    void readClose();
    void readExit();
    MonitoringData readMonitoring();

    Message();

    void read();

    bool operator==(const Message &other) const
    {
        return type == other.type && data == other.data;
    }

    MessageType type;
    std::variant<std::string, json> data;
};

void to_json(json &j, const Message &m);
void from_json(const json &j, Message &m);

#endif