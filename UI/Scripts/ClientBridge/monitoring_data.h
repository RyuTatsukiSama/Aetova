#ifndef MONITORING_DATA_H
#define MONITORING_DATA_H

typedef struct MonitoringData
{
    float dlPrc;
    float dlSpeed;
    float wrPrc;
    float wrSpeed;
};

void to_json(json &j, const MonitoringData &md)
{
    j = json{{"dlPrc", md.dlPrc}, {"dlSpeed", md.dlSpeed},{"wrPrc", md.wrPrc},{"wrSpeed", md.wrSpeed}};
}

void from_json(const json &j, MonitoringData &md)
{
    j.at("dlPrc").get_to(md.dlPrc);
    j.at("dlSpeed").get_to(md.dlSpeed);
    j.at("wrPrc").get_to(md.wrPrc);
    j.at("wrSpeed").get_to(md.wrSpeed);
}


#endif