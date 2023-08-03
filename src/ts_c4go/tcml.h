#ifndef TECO_SMI_H_  // NOLINT
#define TECO_SMI_H_  // NOLINT

#ifdef __cplusplus
extern "C" {
#include <cstdint>
#else
#include <stdint.h>
#endif  // __cplusplus

#include <stdbool.h>
#include <stddef.h>

/**
 * @brief Error codes retured by teco_smi_lib functions
 */
typedef enum {
  TCML_SUCCESS = 0x0,              //!< Operation was successful
  TCML_ERROR_INVALID_ARGS,         //!< Passed in arguments are not valid
  TCML_ERROR_NOT_SUPPORTED,        //!< The requested information or
                                   //!< action is not available for the
                                   //!< given input, on the given system
  TCML_ERROR_FILE_ERROR,           //!< Problem accessing a file. This
                                   //!< may because the operation is not
                                   //!< supported by the Linux kernel
                                   //!< version running on the executing
                                   //!< machine
  TCML_ERROR_PERMISSION,           //!< Permission denied/EACCESS file
                                   //!< error. Many functions require
                                   //!< root access to run.
  TCML_ERROR_OUT_OF_RESOURCES,     //!< Unable to acquire memory or other
                                   //!< resource
  TCML_ERROR_INTERNAL_EXCEPTION,   //!< An internal exception was caught
  TCML_ERROR_INPUT_OUT_OF_BOUNDS,  //!< The provided input is out of
                                   //!< allowable or safe range
  TCML_ERROR_INIT_ERROR,           //!< An error occurred when rsmi
                                   //!< initializing internal data
                                   //!< structures
  TCML_ERROR_NOT_YET_IMPLEMENTED,  //!< The requested function has not
                                   //!< yet been implemented in the
                                   //!< current system for the current
                                   //!< devices
  TCML_ERROR_NOT_FOUND,            //!< An item was searched for but not
                                   //!< found
  TCML_ERROR_INSUFFICIENT_SIZE,    //!< Not enough resources were
                                   //!< available for the operation
  TCML_ERROR_INTERRUPT,            //!< An interrupt occurred during
                                   //!< execution of function
  TCML_ERROR_UNEXPECTED_SIZE,      //!< An unexpected amount of data
                                   //!< was read
  TCML_ERROR_NO_DATA,              //!< No data was found for a given
                                   //!< input
  TCML_ERROR_UNEXPECTED_DATA,      //!< The data read or provided to
                                   //!< function is not what was expected
  TCML_ERROR_BUSY,                 //!< A resource or mutex could not be
                                   //!< acquired because it is already
                                   //!< being used
  TCML_ERROR_REFCOUNT_OVERFLOW,    //!< An internal reference counter
                                   //!< exceeded INT32_MAX

  TCML_ERROR_UNKNOWN_ERROR = 0xFFFFFFFF,  //!< An unknown error occurred
} tcmlReturn_t;

typedef enum {
  TCML_CLOCK_MPE = 0,
  TCML_CLOCK_SPE = 1,
  TCML_CLOCK_HBM = 2,
  TCML_CLOCK_GLB = 3,
  TCML_CLOCK_COUNT,
} tcmlClockType_t;

typedef struct tcmlUtilization_t {
  uint32_t gpu;
  uint32_t memory;
} tcmlUtilization_t;

typedef struct tcmlMemory_t {
  uint64_t total;
  uint64_t free;
  uint64_t used;
} tcmlMemory_t;

typedef struct tcmlClock_t {
  uint64_t current;
  uint64_t max;
} tcmlClock_t;

typedef enum {
  TCML_TOPOLOGY_INTERNAL = 0,
  TCML_TOPOLOGY_SINGLE =
      10,  // all devices that only need traverse a single PCIe switch
  TCML_TOPOLOGY_MULTIPLE =
      20,  // all devices that need not traverse a host bridge
  TCML_TOPOLOGY_HOSTBRIDGE =
      30,  // all devices that are connected to the same host bridge
  TCML_TOPOLOGY_NODE = 40,   // all devices that are connected to the same NUMA
                             // node but possibly multiple host bridges
  TCML_TOPOLOGY_SYSTEM = 50  // all devices in the system

  // there is purposefully no COUNT here because of the need for spacing above
} tcmlCardTopologyLevel_t;

/* P2P Capability Index Status*/
typedef enum {
  TCML_P2P_STATUS_OK = 0,
  TCML_P2P_STATUS_NOT_SUPPORTED,
  TCML_P2P_STATUS_UNKNOWN
} tcmlCardP2PStatus_t;

/* P2P Capability Index*/
typedef enum {
  TCML_P2P_CAPS_INDEX_READ = 0,
  TCML_P2P_CAPS_INDEX_WRITE,
  TCML_P2P_CAPS_INDEX_UNKNOWN
} tcmlCardP2PCapsIndex_t;

/**
 * Represents the queryable PCIe utilization counters
 */
typedef enum tcmlPcieUtilCounter_enum {
  TCML_PCIE_UTIL_TX_BYTES = 0,  // 1B granularity
  TCML_PCIE_UTIL_RX_BYTES = 1,  // 1B granularity

  // Keep this last
  TCML_PCIE_UTIL_COUNT
} tcmlPcieUtilCounter_t;

/**
 * Temperature thresholds.
 */
typedef enum tcmlTemperatureThresholds_enum {
  TCML_TEMPERATURE_THRESHOLD_SHUTDOWN = 0,  // Temperature at which the P-aicard
                                            // will shut down for HW protection
  TCML_TEMPERATURE_THRESHOLD_SLOWDOWN = 1,  // Temperature at which the P-aicard
                                            // will begin HW slowdown
  // Keep this last
  TCML_TEMPERATURE_THRESHOLD_COUNT
} tcmlTemperatureThresholds_t;

/**
 * Represents type of perf policy for which violation times can be queried
 */
typedef enum tcmlPerfPolicyType_enum {
  TCML_PERF_POLICY_THERMAL = 0,  //!< How long did thermal violations cause the
                                 //!< GPU to be below application clocks
  // Keep this last
  TCML_PERF_POLICY_COUNT
} tcmlPerfPolicyType_t;

/**
 * Buffer size guaranteed to be large enough for pci bus id
 */
#define TCML_DEVICE_PCI_BUS_ID_BUFFER_SIZE 32

/**
 * Buffer size guaranteed to be large enough for pci bus id for ::busIdLegacy
 */
#define TCML_DEVICE_PCI_BUS_ID_BUFFER_V2_SIZE 16
/**
 * PCI information about a device.
 */
typedef struct tcmlPciInfo_st {
  char busIdLegacy
      [TCML_DEVICE_PCI_BUS_ID_BUFFER_V2_SIZE];  //!< The legacy tuple
                                                //!< domain:bus:device.function
                                                //!< PCI identifier
                                                //!< (&amp; NULL
                                                //!< terminator)
  uint32_t domain;    //!< The PCI domain on which the device's bus resides, 0
                      //!< to 0xffffffff
  uint32_t bus;       //!< The bus on which the device resides, 0 to 0xff
  uint32_t device;    //!< The device's id on the bus, 0 to 31
  uint32_t function;  //!< The function's id on the bus, 0
  uint32_t pciDeviceId;  //!< The combined 16-bit device id and 16-bit vendor id
  uint32_t pciSubSystemId;  //!< The 32-bit Sub System Device ID

  char busId
      [TCML_DEVICE_PCI_BUS_ID_BUFFER_SIZE];  //!< The tuple
                                             //!< domain:bus:device.function
                                             //!< PCI identifier (&amp;
                                             //!< NULL terminator)
} tcmlPciInfo_t;

/**
 * PCI format string for ::busIdLegacy
 */
#define TCML_DEVICE_PCI_BUS_ID_LEGACY_FMT "%04X:%02X:%02X.%X"

/**
 * PCI format string for ::busId
 */
#define TCML_DEVICE_PCI_BUS_ID_FMT "%08X:%02X:%02X.%X"

/**
 * Utility macro for filling the pci bus id format from a tcmlPciInfo_t
 */
#define TCML_DEVICE_PCI_BUS_ID_FMT_ARGS(pciInfo) \
  (pciInfo)->domain, (pciInfo)->bus, (pciInfo)->device, (pciInfo)->function

// 初始化资源，调用其他接口函数前必须调用
tcmlReturn_t tcmlInitWithFlags(uint64_t flags);  // init_flags正常传入0

// 关闭资源，完成调用其他接口函数后必须调用
tcmlReturn_t tcmlShutdown(void);

// 通过返回值查询详细信息
const char *tcmlErrorString(tcmlReturn_t result);

// 查询设备总数
tcmlReturn_t tcmlDeviceGetCount(uint32_t *deviceCount);

// 查询SMI版本号
tcmlReturn_t tcmlSystemGetTCMLVersion(char *version, uint32_t length);

// 查询aicard-dev驱动版本号
tcmlReturn_t tcmlSystemGetSdaaDriverVersion(char *version, uint32_t length);

// 查询firmware版本号
tcmlReturn_t tcmlSystemGetFirmwareVersion(uint32_t dv_ind, char *version,
                                          uint32_t length);

// 查询sdaa版本号
tcmlReturn_t tcmlSystemGetSdaaRuntimeVersion(char *version, uint32_t length);

// 查询bios版本号
tcmlReturn_t tcmlDeviceGetBiosVersion(uint32_t dv_ind, char *version,
                                      uint32_t length);

// 获取PCIe版本
tcmlReturn_t tcmlDeviceGetCurrPcieLinkGeneration(uint32_t device,
                                                 uint32_t *currLinkGen);
tcmlReturn_t tcmlDeviceGetMaxPcieLinkGeneration(uint32_t device,
                                                uint32_t *maxLinkGen);
tcmlReturn_t tcmlDeviceGetCurrPcieLinkWidth(uint32_t device,
                                            uint32_t *currLinkWidth);
tcmlReturn_t tcmlDeviceGetMaxPcieLinkWidth(uint32_t device,
                                           uint32_t *maxLinkWidth);

/*
 * 查询pcie总线地址信息
 *  The format of bdfid will be as follows:
 *
 *      BDFID = ((DOMAIN & 0xffffffff) << 32) | ((BUS & 0xff) << 8) |
 *                                   ((DEVICE & 0x1f) <<3 ) | (FUNCTION & 0x7)
 *
 *  | Name     | Field   |
 *  ---------- | ------- |
 *  | Domain   | [64:32] |
 *  | Reserved | [31:16] |
 *  | Bus      | [15: 8] |
 *  | Device   | [ 7: 3] |
 *  | Function | [ 2: 0] |
 */
tcmlReturn_t tcmlDeviceGetPciInfo(uint32_t dv_ind, tcmlPciInfo_t *pci);

// 获取PCIe吞吐量 value KB/s
tcmlReturn_t tcmlDeviceGetPcieThroughput(uint32_t dv_ind,
                                         tcmlPcieUtilCounter_t counter,
                                         uint32_t *value);

// 查询加速卡序列号
tcmlReturn_t tcmlDeviceGetSerialNum(uint32_t dv_ind, char *serialNum,
                                    uint32_t length);

// 获取SPE可支持频率
tcmlReturn_t tcmlDeviceGetSupportedSpeClocks(uint32_t dv_ind,
                                             unsigned int *count,
                                             unsigned int *clocksMHz);

// 设置从核调频频率
tcmlReturn_t tcmlDeviceSetSpeLockedClocks(uint32_t dv_ind,
                                          uint32_t speClockMHz);

// 重置spe频率为默认值（最大支持频率）
tcmlReturn_t tcmlDeviceResetSpeLockedClocks(uint32_t dv_ind);

// 获取时钟频率
tcmlReturn_t tcmlDeviceGetClock(uint32_t dv_ind, tcmlClockType_t type,
                                tcmlClock_t *clock);

// 查询加速卡uuid
tcmlReturn_t tcmlDeviceGetUuid(uint32_t dv_ind, char *uuid, uint32_t length);

// 查询产品名称
tcmlReturn_t tcmlDeviceGetProductName(uint32_t dv_ind, char *productname,
                                      uint32_t length);

// 查询产品架构
tcmlReturn_t tcmlDeviceGetProductArchitecture(uint32_t dv_ind,
                                              char *architecture,
                                              uint32_t length);

// 查询加速卡利用率
tcmlReturn_t tcmlDeviceGetUtilizationRates(uint32_t dv_ind,
                                           tcmlUtilization_t *utilization);

// 查询加速卡内存信息
tcmlReturn_t tcmlDeviceGetMemoryInfo(uint32_t dv_ind, tcmlMemory_t *memory);

// 获取加速卡上运行的进程信息
tcmlReturn_t tcmlDeviceGetProcessInfo(uint32_t dv_ind, char *processInfo,
                                      uint32_t length);

// 获取健康状态
tcmlReturn_t tcmlDeviceGetHealthByID(uint32_t dv_ind, uint64_t *error_id);

// 硬件信息
tcmlReturn_t tcmlDeviceGetTemperature(uint32_t dv_ind, float *temperature);
tcmlReturn_t tcmlDeviceGetTemperatureThreshold(
    uint32_t dv_ind, tcmlTemperatureThresholds_t thresholdType,
    unsigned int *temp);
tcmlReturn_t tcmlDeviceGetChipVoltage(uint32_t dv_ind, float *voltage);
tcmlReturn_t tcmlDeviceGetChipCurrent(uint32_t dv_ind, float *current);
tcmlReturn_t tcmlDeviceGetChipPower(uint32_t dv_ind, float *power);
tcmlReturn_t tcmlDeviceGetHbmVoltage(uint32_t dv_ind, float *voltage);
tcmlReturn_t tcmlDeviceGetHbmCurrent(uint32_t dv_ind, float *current);
tcmlReturn_t tcmlDeviceGetHbmPower(uint32_t dv_ind, float *power);
tcmlReturn_t tcmlDeviceGetMcuSoftwareVersion(uint32_t dv_ind, char *mcuVersion,
                                             uint32_t length);
tcmlReturn_t tcmlDeviceGetPcbVersion(uint32_t dv_ind, char *pcbVersion,
                                     uint32_t length);

// tcml topo info
tcmlReturn_t tcmlDeviceGetTopologyCommonAncestor(
    uint32_t device1, uint32_t device2, tcmlCardTopologyLevel_t *pathInfo);
tcmlReturn_t tcmlDeviceGetCpuAffinity(uint32_t dv_ind, char *cpu_list,
                                      uint32_t length);
tcmlReturn_t tcmlDeviceGetNodeAffinity(uint32_t dv_ind, uint32_t *nodeSet);
tcmlReturn_t tcmlDeviceGetP2PStatus(uint32_t device1, uint32_t device2,
                                    tcmlCardP2PCapsIndex_t p2pIndex,
                                    tcmlCardP2PStatus_t *p2pStatus);

tcmlReturn_t tcmlDeviceGetPcieReplayCounter(uint32_t dv_ind, uint32_t *value);

tcmlReturn_t tcmlDeviceGetViolationStatus(uint32_t dv_ind,
                                          tcmlPerfPolicyType_t perfPolicyType,
                                          char *violTime, uint32_t length);

tcmlReturn_t tcmlDeviceReset(uint32_t dv_ind);
#ifdef __cplusplus
}
#endif  // __cplusplus

#endif  // NOLINT
