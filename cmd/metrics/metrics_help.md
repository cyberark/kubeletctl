A list of metrics taken from a tested kubelet
=============================================
```
apiserver_audit_event_total [ALPHA] Counter of audit events generated and sent to the audit backend.
apiserver_audit_requests_rejected_total [ALPHA] Counter of apiserver requests rejected due to an error in audit logging backend.
apiserver_client_certificate_expiration_seconds [ALPHA] Distribution of the remaining lifetime on the certificate used to authenticate a request.
apiserver_storage_data_key_generation_duration_seconds [ALPHA] Latencies in seconds of data encryption key(DEK) generation operations.
apiserver_storage_data_key_generation_failures_total [ALPHA] Total number of failed data encryption key(DEK) generation operations.
apiserver_storage_envelope_transformation_cache_misses_total [ALPHA] Total number of cache misses while accessing key decryption key(KEK).
get_token_count [ALPHA] Counter of total Token() requests to the alternate token source
get_token_fail_count [ALPHA] Counter of failed Token() requests to the alternate token source
go_gc_duration_seconds A summary of the GC invocation durations.
go_goroutines Number of goroutines that currently exist.
go_info Information about the Go environment.
go_memstats_alloc_bytes Number of bytes allocated and still in use.
go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
go_memstats_frees_total Total number of frees.
go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
go_memstats_heap_objects Number of allocated objects.
go_memstats_heap_released_bytes Number of heap bytes released to OS.
go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
go_memstats_lookups_total Total number of pointer lookups.
go_memstats_mallocs_total Total number of mallocs.
go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
go_memstats_other_sys_bytes Number of bytes used for other system allocations.
go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
go_memstats_sys_bytes Number of bytes obtained from system.
go_threads Number of OS threads created.
kubelet_certificate_manager_client_expiration_renew_errors [ALPHA] Counter of certificate renewal errors.
kubelet_certificate_manager_client_expiration_seconds [ALPHA] Gauge of the lifetime of a certificate. The value is the date the certificate will expire in seconds since January 1, 1970 UTC.
kubelet_cgroup_manager_duration_seconds [ALPHA] Duration in seconds for cgroup manager operations. Broken down by method.
kubelet_container_log_filesystem_used_bytes [ALPHA] Bytes used by the container's logs on the filesystem.
kubelet_containers_per_pod_count [ALPHA] The number of containers per pod.
kubelet_docker_operations_duration_seconds [ALPHA] Latency in seconds of Docker operations. Broken down by operation type.
kubelet_docker_operations_total [ALPHA] Cumulative number of Docker operations by operation type.
kubelet_http_inflight_requests [ALPHA] Number of the inflight http requests
kubelet_http_requests_duration_seconds [ALPHA] Duration in seconds to serve http requests
kubelet_http_requests_total [ALPHA] Number of the http requests received since the server started
kubelet_network_plugin_operations_duration_seconds [ALPHA] Latency in seconds of network plugin operations. Broken down by operation type.
kubelet_node_config_error [ALPHA] This metric is true (1) if the node is experiencing a configuration-related error, false (0) otherwise.
kubelet_node_name [ALPHA] The node's name. The count is always 1.
kubelet_pleg_discard_events [ALPHA] The number of discard events in PLEG.
kubelet_pleg_last_seen_seconds [ALPHA] Timestamp in seconds when PLEG was last seen active.
kubelet_pleg_relist_duration_seconds [ALPHA] Duration in seconds for relisting pods in PLEG.
kubelet_pleg_relist_interval_seconds [ALPHA] Interval in seconds between relisting in PLEG.
kubelet_pod_start_duration_seconds [ALPHA] Duration in seconds for a single pod to go from pending to running.
kubelet_pod_worker_duration_seconds [ALPHA] Duration in seconds to sync a single pod. Broken down by operation type: create, update, or sync
kubelet_pod_worker_start_duration_seconds [ALPHA] Duration in seconds from seeing a pod to starting a worker.
kubelet_run_podsandbox_duration_seconds [ALPHA] Duration in seconds of the run_podsandbox operations. Broken down by RuntimeClass.Handler.
kubelet_running_container_count [ALPHA] Number of containers currently running
kubelet_running_pod_count [ALPHA] Number of pods currently running
kubelet_runtime_operations_duration_seconds [ALPHA] Duration in seconds of runtime operations. Broken down by operation type.
kubelet_runtime_operations_errors_total [ALPHA] Cumulative number of runtime operation errors by operation type.
kubelet_runtime_operations_total [ALPHA] Cumulative number of runtime operations by operation type.
kubernetes_build_info [ALPHA] A metric with a constant '1' value labeled by major, minor, git version, git commit, git tree state, build date, Go version, and compiler from which Kubernetes was built, and platform on which it is running.
process_cpu_seconds_total Total user and system CPU time spent in seconds.
process_max_fds Maximum number of open file descriptors.
process_open_fds Number of open file descriptors.
process_resident_memory_bytes Resident memory size in bytes.
process_start_time_seconds Start time of the process since unix epoch in seconds.
process_virtual_memory_bytes Virtual memory size in bytes.
process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
rest_client_exec_plugin_certificate_rotation_age [ALPHA] Histogram of the number of seconds the last auth exec plugin client certificate lived before being rotated. If auth exec plugin client certificates are unused, histogram will contain no data.
rest_client_exec_plugin_ttl_seconds [ALPHA] Gauge of the shortest TTL (time-to-live) of the client certificate(s) managed by the auth exec plugin. The value is in seconds until certificate expiry (negative if already expired). If auth exec plugins are unused or manage no TLS certificates, the value will be +INF.
rest_client_request_duration_seconds [ALPHA] Request latency in seconds. Broken down by verb and URL.
rest_client_requests_total [ALPHA] Number of HTTP requests, partitioned by status code, method, and host.
storage_operation_duration_seconds [ALPHA] Storage operation duration
storage_operation_status_count [ALPHA] Storage operation return statuses count
volume_manager_total_volumes [ALPHA] Number of volumes in Volume Manager
```