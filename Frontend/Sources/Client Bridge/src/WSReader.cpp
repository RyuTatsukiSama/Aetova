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
    MonitoringData md = std::get<json>(data).get<MonitoringData>();

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
    j["type"] = m.type;

    std::visit([&](const auto& value){
        using T = std::decay_t<decltype(value)>;

        if constexpr (std::is_same_v<T,json>)
        {
            j["data"] = value.dump();
        }
        else
        {
            j["data"] = value;
        }
    }, m.data);
}

void from_json(const json &j, Message &m)
{
    j.at("type").get_to(m.type);
    
    const auto& datafield = j.at("Data");

    if (datafield.is_object() || datafield.is_array())
    {
        m.data = datafield;
    }
    else if (datafield.is_string())
    {
        std::string dataStr = datafield.get<std::string>();

        try
        {
            m.data = json::parse(dataStr);
        }
        catch (...)
        {
            m.data = dataStr;
        }
    }
    else
    {
        m.data = datafield.dump();
    }
}