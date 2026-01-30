#include "WSReader.h"
#include <docLogger>
#include "../monitoring_data.h"

Message::Message()
{
    log = new doc::Logger();
}

void Message::readText()
{
}

void Message::readClose()
{
}

void Message::readExit()
{
}

void Message::readMonitoring()
{
    MonitoringData md = json::parse(data).get<MonitoringData>();

    std::cout << std::format("Download {}%% at speed {} kB/s, Writing {}%% at speed {} kB/s\n", md.dlPrc, md.dlSpeed, md.wrPrc, md.wrSpeed);
}

void Message::read()
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

void to_json(json &j, const Message &m)
{
    j = json{{"type", m.type}, {"Data", m.data}};
}

void from_json(const json &j, Message &m)
{
    j.at("type").get_to(m.type);
    j.at("Data").get_to(m.data);
}