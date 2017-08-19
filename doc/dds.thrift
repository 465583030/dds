/**
 * DDS Thrift TDL
 * @author ricl
 * @time 2017.08.19
 */

namespace go ddservice
namespace java ddservice

service ddservice
{
    /* ip address, timestamp, data exchange id, payload */
	string call(1:string ip,2:i64 time,3:i32 cid,4:string payload)
}