#ifndef MONITORING_DATA_H
#define MONITORING_DATA_H

#include <nlohmann/json.hpp>
using json = nlohmann::json;

typedef struct MonitoringData
{
    float dlPrc;
    float dlSpeed;
    float wrPrc;
    float wrSpeed;

    bool operator==(const MonitoringData &other) const
    {
        return dlPrc == other.dlPrc && dlSpeed == other.dlSpeed && wrPrc == other.wrPrc && wrSpeed == other.wrSpeed;
    }
} MonitoringData;

inline void to_json(json &j, const MonitoringData &md)
{
    j = json{{"DlPrc", md.dlPrc}, {"DlSpeed", md.dlSpeed}, {"WrPrc", md.wrPrc}, {"WrSpeed", md.wrSpeed}};
}

inline void from_json(const json &j, MonitoringData &md)
{
    j.at("DlPrc").get_to(md.dlPrc);
    j.at("DlSpeed").get_to(md.dlSpeed);
    j.at("WrPrc").get_to(md.wrPrc);
    j.at("WrSpeed").get_to(md.wrSpeed);
}

#endif