# gocpu

A simple cli tool to handle and watch your CPU

<br />

## Todo

- [x] Real Time CPU usage from frequency files in linux `/sys/devices/system/cpu/cpu[0-9]*/cpufreq/scaling_cur_freq`
- [x] CPU governor changing: all governors `/sys/devices/system/cpu/cpu0/cpufreq/scaling_available_governors`
- [x] disable/enable turbo boost through file control: `/sys/devices/system/cpu/intel_pstate/no_turbo`
