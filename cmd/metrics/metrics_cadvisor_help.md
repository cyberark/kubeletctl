A list of metrics taken from a tested kubelet
=============================================

```
cadvisor_version_info A metric with a constant '1' value labeled by kernel version, OS version, docker version, cadvisor version & cadvisor revision.
container_cpu_load_average_10s Value of container cpu load average over the last 10 seconds.
container_cpu_system_seconds_total Cumulative system cpu time consumed in seconds.
container_cpu_usage_seconds_total Cumulative cpu time consumed in seconds.
container_cpu_user_seconds_total Cumulative user cpu time consumed in seconds.
container_file_descriptors Number of open file descriptors for the container.
container_fs_inodes_free Number of available Inodes
container_fs_inodes_total Number of Inodes
container_fs_io_current Number of I/Os currently in progress
container_fs_io_time_seconds_total Cumulative count of seconds spent doing I/Os
container_fs_io_time_weighted_seconds_total Cumulative weighted I/O time in seconds
container_fs_limit_bytes Number of bytes that can be consumed by the container on this filesystem.
container_fs_read_seconds_total Cumulative count of seconds spent reading
container_fs_reads_bytes_total Cumulative count of bytes read
container_fs_reads_merged_total Cumulative count of reads merged
container_fs_reads_total Cumulative count of reads completed
container_fs_sector_reads_total Cumulative count of sector reads completed
container_fs_sector_writes_total Cumulative count of sector writes completed
container_fs_usage_bytes Number of bytes that are consumed by the container on this filesystem.
container_fs_write_seconds_total Cumulative count of seconds spent writing
container_fs_writes_bytes_total Cumulative count of bytes written
container_fs_writes_merged_total Cumulative count of writes merged
container_fs_writes_total Cumulative count of writes completed
container_last_seen Last time a container was seen by the exporter
container_memory_cache Number of bytes of page cache memory.
container_memory_failcnt Number of memory usage hits limits
container_memory_failures_total Cumulative count of memory allocation failures.
container_memory_mapped_file Size of memory mapped files in bytes.
container_memory_max_usage_bytes Maximum memory usage recorded in bytes
container_memory_rss Size of RSS in bytes.
container_memory_swap Container swap usage in bytes.
container_memory_usage_bytes Current memory usage in bytes, including all memory regardless of when it was accessed
container_memory_working_set_bytes Current working set in bytes.
container_network_receive_bytes_total Cumulative count of bytes received
container_network_receive_errors_total Cumulative count of errors encountered while receiving
container_network_receive_packets_dropped_total Cumulative count of packets dropped while receiving
container_network_receive_packets_total Cumulative count of packets received
container_network_transmit_bytes_total Cumulative count of bytes transmitted
container_network_transmit_errors_total Cumulative count of errors encountered while transmitting
container_network_transmit_packets_dropped_total Cumulative count of packets dropped while transmitting
container_network_transmit_packets_total Cumulative count of packets transmitted
container_processes Number of processes running inside the container.
container_scrape_error 1 if there was an error while getting container metrics, 0 otherwise
container_sockets Number of open sockets for the container.
container_spec_cpu_period CPU period of the container.
container_spec_cpu_shares CPU share of the container.
container_spec_memory_limit_bytes Memory limit for the container.
container_spec_memory_reservation_limit_bytes Memory reservation limit for the container.
container_spec_memory_swap_limit_bytes Memory swap limit for the container.
container_start_time_seconds Start time of the container since unix epoch in seconds.
container_tasks_state Number of tasks in given state
container_threads Number of threads running inside the container
container_threads_max Maximum number of threads allowed inside the container, infinity if value is zero
```