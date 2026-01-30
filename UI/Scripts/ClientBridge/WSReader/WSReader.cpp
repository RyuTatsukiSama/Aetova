#include "WSReader.h"
#include <docLogger>
#include "../monitoring_data.h"

WSReader::WSReader()
{
    log = new doc::Logger();
}

void WSReader::readText()
{
    log->Debug(data);
}

void WSReader::readClose()
{
}

void WSReader::readExit()
{
}

void WSReader::readMonitoring()
{
}

void WSReader::read()
{
    switch (type)
    {
    case TEXT:
        readText();
        break;
    case CLOSE:
        readClose();
        break;
    case EXIT:
        readExit();
        break;
    case MONITORING:
        readMonitoring();
        break;
    default:
        break;
    }
}

void to_json(json &j, const WSReader &m)
{
    j = json{{"type", m.type}, {"data", m.data}};
}

void from_json(const json &j, WSReader &m)
{
    j.at("type").get_to(m.type);
    j.at("data").get_to(m.data);
}