#ifndef WSREADER_H
#define WSREADER_H

#include <string>
#include <nlohmann/json.hpp>
using json = nlohmann::json;
namespace doc {
	class Logger;
}

enum MessageType
{
    TEXT,
    CLOSE,
    EXIT,
    MONITORING
};

class WSReader
{
	doc::Logger* log;

    void readText();
    void readClose();
    void readExit();
    void readMonitoring();

public:
    WSReader();

    void read();

    bool operator==(const WSReader &other) const
    {
        return type == other.type && data == other.data;
    }

    MessageType type;
    std::string data;
};

void to_json(json &j, const WSReader &m);
void from_json(const json &j, WSReader &m);

#endif