This repo:

  * Builds a static library (gofunction.a) from am Go package (`gofunction/`)
    which exports an empty Go function `GoFunction`
  * Links that into a C library (`c/`) which exports a function that simply
    calls the empty Go funtion.
  * Then has a `main` package which has a function
    `callCFunctionWhichCallsGoFunction` that calls the C function using Cgo. But
    that function is not actually called from `main()`.

it crashes with a segfault.

So we have Go -> C -> Go. The real version of this problem is me trying to pass
a Handle pointing to a logrus.Entry into C, so that the C part of my program can
call the same loggers as the Go part. This is a simplified version of that which
breaks in the same way.

Here's the backtrace. It's also running in GH Actions in this repo.

```
Thread 3 "main" received signal SIGSEGV, Segmentation fault.
[Switching to Thread 0x7fffa3fff640 (LWP 4122)]
0x0000000000439af3 in runtime.mstart0 () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/proc.go:1431
1431		gp := getg()

Thread 5 (Thread 0x7fffa1ffd640 (LWP 4124) "main"):
#0  clone3 () at ../sysdeps/unix/sysv/linux/x86_64/clone3.S:62
No locals.
#1  0x0000000000000000 in ?? ()
No symbol table info available.

Thread 4 (Thread 0x7fffa2ffe640 (LWP 4123) "main"):
#0  runtime.usleep () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/sys_linux_amd64.s:135
No locals.
#1  0x00007ffff7ef63e5 in runtime.sysmon () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/proc.go:5298
        now = <optimized out>
        lastpoll = <optimized out>
        ~R0 = <optimized out>
        lasttrace = 0
        idle = 59
        delay = 10000
#2  0x00007ffff7eedd95 in runtime.mstart1 () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/proc.go:1499
        fn = {void (void)} 0x0
        gp = <optimized out>
#3  0x00007ffff7eedcda in runtime.mstart0 () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/proc.go:1456
        gp = 0x1c0000069c0
        osStack = 252
#4  0x00007ffff7f116e5 in runtime.mstart () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/asm_amd64.s:394
No locals.
#5  0x000000000046583c in crosscall_amd64 () at gcc_amd64.S:42
No locals.
#6  0x00007ffff7c70a60 in ?? ()
No symbol table info available.
#7  0x00007ffff7d09850 in ?? () at ./nptl/pthread_create.c:321 from /lib/x86_64-linux-gnu/libc.so.6
        RSEQ_CS_FLAG_NO_RESTART_ON_PREEMPT_BIT = RSEQ_CS_FLAG_NO_RESTART_ON_PREEMPT_BIT
        arch_kind_other = arch_kind_other
        ns_s_zn = ns_s_qd
        _bitindex_arch_Fast_Unaligned_Load = _bitindex_arch_Fast_Unaligned_Load
        _bitindex_arch_AVX_Fast_Unaligned_Load = _bitindex_arch_AVX_Fast_Unaligned_Load
        PTHREAD_MUTEX_TIMED_NP = PTHREAD_MUTEX_TIMED_NP
        PTHREAD_MUTEX_RECURSIVE_NP = PTHREAD_MUTEX_RECURSIVE_NP
        PTHREAD_MUTEX_ERRORCHECK_NP = PTHREAD_MUTEX_ERRORCHECK_NP
        PTHREAD_MUTEX_ADAPTIVE_NP = PTHREAD_MUTEX_ADAPTIVE_NP
        PTHREAD_MUTEX_NORMAL = PTHREAD_MUTEX_TIMED_NP
        PTHREAD_MUTEX_RECURSIVE = PTHREAD_MUTEX_RECURSIVE_NP
        PTHREAD_MUTEX_ERRORCHECK = PTHREAD_MUTEX_ERRORCHECK_NP
        PTHREAD_MUTEX_DEFAULT = PTHREAD_MUTEX_TIMED_NP
        PTHREAD_MUTEX_FAST_NP = PTHREAD_MUTEX_TIMED_NP
        PTHREAD_CANCEL_ENABLE = PTHREAD_CANCEL_ENABLE
        PTHREAD_CANCEL_DEFERRED = PTHREAD_CANCEL_DEFERRED
        _URC_NORMAL_STOP = _URC_NORMAL_STOP
        cpuid_register_index_ebx = cpuid_register_index_ebx
        RT_CONSISTENT = RT_CONSISTENT
        ns_s_an = ns_s_an
        TD_EVENT_NONE = TD_ALL_EVENTS
        RSEQ_CPU_ID_REGISTRATION_FAILED = RSEQ_CPU_ID_REGISTRATION_FAILED
        nonexisting = nonexisting
        cet_permissive = cet_permissive
        lt_executable = lt_executable
        cet_always_on = cet_always_on
        _bitindex_arch_MathVec_Prefer_No_AVX512 = _bitindex_arch_MathVec_Prefer_No_AVX512
        dso_sort_algorithm_original = dso_sort_algorithm_original
        ns_s_ar = ns_s_ar
        _URC_END_OF_STACK = _URC_END_OF_STACK
        _URC_INSTALL_CONTEXT = _URC_INSTALL_CONTEXT
        dso_sort_algorithm_dfs = dso_sort_algorithm_dfs
        cpuid_register_index_ecx = cpuid_register_index_ecx
        lc_property_none = lc_property_none
        TD_READY = TD_READY
        ns_s_ud = ns_s_ns
        PREFERRED_FEATURE_INDEX_1 = PREFERRED_FEATURE_INDEX_1
        PREFERRED_FEATURE_INDEX_MAX = PREFERRED_FEATURE_INDEX_MAX
        lc_property_valid = lc_property_valid
        TD_CREATE = TD_CREATE
        TD_TIMEOUT = TD_TIMEOUT
        TD_MAX_EVENT_NUM = TD_TIMEOUT
        RSEQ_CPU_ID_UNINITIALIZED = RSEQ_CPU_ID_UNINITIALIZED
        arch_kind_unknown = arch_kind_unknown
        arch_kind_intel = arch_kind_intel
        lc_property_unknown = lc_property_unknown
        cpuid_register_index_edx = cpuid_register_index_edx
        _bitindex_arch_Prefer_No_VZEROUPPER = _bitindex_arch_Prefer_No_VZEROUPPER
        _bitindex_arch_Prefer_ERMS = _bitindex_arch_Prefer_ERMS
        lt_library = lt_library
        _bitindex_arch_Prefer_FSRM = _bitindex_arch_Prefer_FSRM
        _bitindex_arch_Avoid_Short_Distance_REP_MOVSB = _bitindex_arch_Avoid_Short_Distance_REP_MOVSB
        PTHREAD_CANCEL_ASYNCHRONOUS = PTHREAD_CANCEL_ASYNCHRONOUS
        _bitindex_arch_Fast_Copy_Backward = _bitindex_arch_Fast_Copy_Backward
        existing = existing
        _bitindex_arch_Fast_Rep_String = _bitindex_arch_Fast_Rep_String
        TD_EVENTS_ENABLE = TD_EVENTS_ENABLE
        _bitindex_arch_Slow_SSE4_2 = _bitindex_arch_Slow_SSE4_2
        RT_ADD = RT_ADD
        arch_kind_amd = arch_kind_amd
        CPUID_INDEX_1 = CPUID_INDEX_1
        CPUID_INDEX_7 = CPUID_INDEX_7
        CPUID_INDEX_80000001 = CPUID_INDEX_80000001
        CPUID_INDEX_D_ECX_1 = CPUID_INDEX_D_ECX_1
        CPUID_INDEX_80000007 = CPUID_INDEX_80000007
        CPUID_INDEX_80000008 = CPUID_INDEX_80000008
        CPUID_INDEX_7_ECX_1 = CPUID_INDEX_7_ECX_1
        CPUID_INDEX_19 = CPUID_INDEX_19
        CPUID_INDEX_14_ECX_0 = CPUID_INDEX_14_ECX_0
        CPUID_INDEX_MAX = CPUID_INDEX_MAX
        arch_kind_zhaoxin = arch_kind_zhaoxin
        TD_SLEEP = TD_SLEEP
        _bitindex_arch_I686 = _bitindex_arch_I686
        TD_MIN_EVENT_NUM = TD_READY
        unknown = unknown
        TD_SWITCHFROM = TD_SWITCHFROM
        TD_PREEMPT = TD_PREEMPT
        TD_CONCURRENCY = TD_CONCURRENCY
        TD_DEATH = TD_DEATH
        _URC_NO_REASON = _URC_NO_REASON
        _URC_FATAL_PHASE2_ERROR = _URC_FATAL_PHASE2_ERROR
        TD_SWITCHTO = TD_SWITCHTO
        _URC_FATAL_PHASE1_ERROR = _URC_FATAL_PHASE1_ERROR
        RSEQ_CS_FLAG_NO_RESTART_ON_MIGRATE_BIT = RSEQ_CS_FLAG_NO_RESTART_ON_MIGRATE_BIT
        _URC_FOREIGN_EXCEPTION_CAUGHT = _URC_FOREIGN_EXCEPTION_CAUGHT
        _bitindex_arch_Fast_Unaligned_Copy = _bitindex_arch_Fast_Unaligned_Copy
        ns_s_ns = ns_s_ns
        ns_s_max = ns_s_max
        _bitindex_arch_I586 = _bitindex_arch_I586
        _bitindex_arch_Prefer_No_AVX512 = _bitindex_arch_Prefer_No_AVX512
        TD_ALL_EVENTS = TD_ALL_EVENTS
        _URC_HANDLER_FOUND = _URC_HANDLER_FOUND
        _URC_CONTINUE_UNWIND = _URC_CONTINUE_UNWIND
        TD_IDLE = TD_IDLE
        ns_s_qd = ns_s_qd
        cet_elf_property = cet_elf_property
        RT_DELETE = RT_DELETE
        lt_loaded = lt_loaded
        PTHREAD_CANCEL_DISABLE = PTHREAD_CANCEL_DISABLE
        TD_CATCHSIG = TD_CATCHSIG
        TD_PRI_INHERIT = TD_PRI_INHERIT
        cet_always_off = cet_always_off
        TD_LOCK_TRY = TD_LOCK_TRY
        sigall_set = {__val = {18446744073709551615 <repeats 16 times>}}
        ns_s_pr = ns_s_an
        TD_REAP = TD_REAP
        _bitindex_arch_Slow_BSF = _bitindex_arch_Slow_BSF
        RSEQ_CS_FLAG_NO_RESTART_ON_SIGNAL_BIT = RSEQ_CS_FLAG_NO_RESTART_ON_SIGNAL_BIT
        _bitindex_arch_Prefer_PMINUB_for_stringop = _bitindex_arch_Prefer_PMINUB_for_stringop
        cpuid_register_index_eax = cpuid_register_index_eax
        _thread_db_pthread_eventbuf_eventmask_event_bits = {32, 2, 1616}
        _thread_db_rtld_global__dl_tls_dtv_slotinfo_list = {64, 1, 4176}
        _thread_db_pthread_schedpolicy = {32, 1, 1596}
        _thread_db___pthread_keys = {128, 1024, 0}
        _thread_db_const_thread_area = 25
        _thread_db_pthread_eventbuf = {192, 1, 1616}
        _thread_db___nptl_initial_report_events = {8, 1, 0}
        _thread_db_sizeof_list_t = 16
        _thread_db_link_map_l_tls_offset = {64, 1, 1112}
        _thread_db_sizeof_td_eventbuf_t = 24
        _thread_db_link_map_l_tls_modid = {64, 1, 1120}
        __nptl_rtld_global = 0x7ffff7ffd040 <_rtld_global>
        _thread_db_pthread_tid = {32, 1, 720}
        _thread_db_pthread_eventbuf_eventmask = {64, 1, 1616}
        _thread_db_pthread_nextevent = {64, 1, 1640}
        _thread_db_td_eventbuf_t_eventdata = {64, 1, 16}
        _thread_db___nptl_nthreads = {32, 1, 0}
        _thread_db_dtv_slotinfo_list_slotinfo = {128, 0, 16}
        _thread_db_pthread_list = {128, 1, 704}
        _thread_db_sizeof_pthread_key_data_level2 = 512
        _thread_db_list_t_prev = {64, 1, 8}
        _thread_db_td_thr_events_t_event_bits = {32, 2, 0}
        __GI___nptl_threads_events = {event_bits = {0, 0}}
        _thread_db_pthread_start_routine = {64, 1, 1600}
        _thread_db_pthread_schedparam_sched_priority = {32, 1, 1592}
        _thread_db_pthread_specific = {2048, 1, 1296}
        _thread_db_td_eventbuf_t_eventnum = {32, 1, 8}
        _thread_db_dtv_dtv = {128, 134217727, 0}
        _thread_db_sizeof_pthread = 2496
        _thread_db_sizeof_td_thr_events_t = 8
        _thread_db_pthread_key_data_level2_data = {128, 32, 0}
        _thread_db_pthread_report_events = {8, 1, 1553}
        _thread_db_pthread_cancelhandling = {32, 1, 776}
        _thread_db_list_t_next = {64, 1, 0}
        __GI___nptl_last_event = 0x0
        __nptl_version = "2.35"
        _thread_db_rtld_global__dl_stack_user = {128, 1, 4248}
        _thread_db_rtld_global__dl_stack_used = {128, 1, 4232}
#8  0x0000000000000000 in ?? ()
No symbol table info available.

Thread 3 (Thread 0x7fffa3fff640 (LWP 4122) "main"):
#0  0x0000000000439af3 in runtime.mstart0 () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/proc.go:1431
        gp = 0x0
        osStack = <optimized out>
#1  0x000000000045f6e5 in runtime.mstart () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/asm_amd64.s:394
No locals.
#2  0x000000000046583c in crosscall_amd64 () at gcc_amd64.S:42
No locals.
#3  0x00007fffffffd400 in ?? ()
No symbol table info available.
#4  0x00007ffff7d09850 in ?? () at ./nptl/pthread_create.c:321 from /lib/x86_64-linux-gnu/libc.so.6
        RSEQ_CS_FLAG_NO_RESTART_ON_PREEMPT_BIT = RSEQ_CS_FLAG_NO_RESTART_ON_PREEMPT_BIT
        arch_kind_other = arch_kind_other
        ns_s_zn = ns_s_qd
        _bitindex_arch_Fast_Unaligned_Load = _bitindex_arch_Fast_Unaligned_Load
        _bitindex_arch_AVX_Fast_Unaligned_Load = _bitindex_arch_AVX_Fast_Unaligned_Load
        PTHREAD_MUTEX_TIMED_NP = PTHREAD_MUTEX_TIMED_NP
        PTHREAD_MUTEX_RECURSIVE_NP = PTHREAD_MUTEX_RECURSIVE_NP
        PTHREAD_MUTEX_ERRORCHECK_NP = PTHREAD_MUTEX_ERRORCHECK_NP
        PTHREAD_MUTEX_ADAPTIVE_NP = PTHREAD_MUTEX_ADAPTIVE_NP
        PTHREAD_MUTEX_NORMAL = PTHREAD_MUTEX_TIMED_NP
        PTHREAD_MUTEX_RECURSIVE = PTHREAD_MUTEX_RECURSIVE_NP
        PTHREAD_MUTEX_ERRORCHECK = PTHREAD_MUTEX_ERRORCHECK_NP
        PTHREAD_MUTEX_DEFAULT = PTHREAD_MUTEX_TIMED_NP
        PTHREAD_MUTEX_FAST_NP = PTHREAD_MUTEX_TIMED_NP
        PTHREAD_CANCEL_ENABLE = PTHREAD_CANCEL_ENABLE
        PTHREAD_CANCEL_DEFERRED = PTHREAD_CANCEL_DEFERRED
        _URC_NORMAL_STOP = _URC_NORMAL_STOP
        cpuid_register_index_ebx = cpuid_register_index_ebx
        RT_CONSISTENT = RT_CONSISTENT
        ns_s_an = ns_s_an
        TD_EVENT_NONE = TD_ALL_EVENTS
        RSEQ_CPU_ID_REGISTRATION_FAILED = RSEQ_CPU_ID_REGISTRATION_FAILED
        nonexisting = nonexisting
        cet_permissive = cet_permissive
        lt_executable = lt_executable
        cet_always_on = cet_always_on
        _bitindex_arch_MathVec_Prefer_No_AVX512 = _bitindex_arch_MathVec_Prefer_No_AVX512
        dso_sort_algorithm_original = dso_sort_algorithm_original
        ns_s_ar = ns_s_ar
        _URC_END_OF_STACK = _URC_END_OF_STACK
        _URC_INSTALL_CONTEXT = _URC_INSTALL_CONTEXT
        dso_sort_algorithm_dfs = dso_sort_algorithm_dfs
        cpuid_register_index_ecx = cpuid_register_index_ecx
        lc_property_none = lc_property_none
        TD_READY = TD_READY
        ns_s_ud = ns_s_ns
        PREFERRED_FEATURE_INDEX_1 = PREFERRED_FEATURE_INDEX_1
        PREFERRED_FEATURE_INDEX_MAX = PREFERRED_FEATURE_INDEX_MAX
        lc_property_valid = lc_property_valid
        TD_CREATE = TD_CREATE
        TD_TIMEOUT = TD_TIMEOUT
        TD_MAX_EVENT_NUM = TD_TIMEOUT
        RSEQ_CPU_ID_UNINITIALIZED = RSEQ_CPU_ID_UNINITIALIZED
        arch_kind_unknown = arch_kind_unknown
        arch_kind_intel = arch_kind_intel
        lc_property_unknown = lc_property_unknown
        cpuid_register_index_edx = cpuid_register_index_edx
        _bitindex_arch_Prefer_No_VZEROUPPER = _bitindex_arch_Prefer_No_VZEROUPPER
        _bitindex_arch_Prefer_ERMS = _bitindex_arch_Prefer_ERMS
        lt_library = lt_library
        _bitindex_arch_Prefer_FSRM = _bitindex_arch_Prefer_FSRM
        _bitindex_arch_Avoid_Short_Distance_REP_MOVSB = _bitindex_arch_Avoid_Short_Distance_REP_MOVSB
        PTHREAD_CANCEL_ASYNCHRONOUS = PTHREAD_CANCEL_ASYNCHRONOUS
        _bitindex_arch_Fast_Copy_Backward = _bitindex_arch_Fast_Copy_Backward
        existing = existing
        _bitindex_arch_Fast_Rep_String = _bitindex_arch_Fast_Rep_String
        TD_EVENTS_ENABLE = TD_EVENTS_ENABLE
        _bitindex_arch_Slow_SSE4_2 = _bitindex_arch_Slow_SSE4_2
        RT_ADD = RT_ADD
        arch_kind_amd = arch_kind_amd
        CPUID_INDEX_1 = CPUID_INDEX_1
        CPUID_INDEX_7 = CPUID_INDEX_7
        CPUID_INDEX_80000001 = CPUID_INDEX_80000001
        CPUID_INDEX_D_ECX_1 = CPUID_INDEX_D_ECX_1
        CPUID_INDEX_80000007 = CPUID_INDEX_80000007
        CPUID_INDEX_80000008 = CPUID_INDEX_80000008
        CPUID_INDEX_7_ECX_1 = CPUID_INDEX_7_ECX_1
        CPUID_INDEX_19 = CPUID_INDEX_19
        CPUID_INDEX_14_ECX_0 = CPUID_INDEX_14_ECX_0
        CPUID_INDEX_MAX = CPUID_INDEX_MAX
        arch_kind_zhaoxin = arch_kind_zhaoxin
        TD_SLEEP = TD_SLEEP
        _bitindex_arch_I686 = _bitindex_arch_I686
        TD_MIN_EVENT_NUM = TD_READY
        unknown = unknown
        TD_SWITCHFROM = TD_SWITCHFROM
        TD_PREEMPT = TD_PREEMPT
        TD_CONCURRENCY = TD_CONCURRENCY
        TD_DEATH = TD_DEATH
        _URC_NO_REASON = _URC_NO_REASON
        _URC_FATAL_PHASE2_ERROR = _URC_FATAL_PHASE2_ERROR
        TD_SWITCHTO = TD_SWITCHTO
        _URC_FATAL_PHASE1_ERROR = _URC_FATAL_PHASE1_ERROR
        RSEQ_CS_FLAG_NO_RESTART_ON_MIGRATE_BIT = RSEQ_CS_FLAG_NO_RESTART_ON_MIGRATE_BIT
        _URC_FOREIGN_EXCEPTION_CAUGHT = _URC_FOREIGN_EXCEPTION_CAUGHT
        _bitindex_arch_Fast_Unaligned_Copy = _bitindex_arch_Fast_Unaligned_Copy
        ns_s_ns = ns_s_ns
        ns_s_max = ns_s_max
        _bitindex_arch_I586 = _bitindex_arch_I586
        _bitindex_arch_Prefer_No_AVX512 = _bitindex_arch_Prefer_No_AVX512
        TD_ALL_EVENTS = TD_ALL_EVENTS
        _URC_HANDLER_FOUND = _URC_HANDLER_FOUND
        _URC_CONTINUE_UNWIND = _URC_CONTINUE_UNWIND
        TD_IDLE = TD_IDLE
        ns_s_qd = ns_s_qd
        cet_elf_property = cet_elf_property
        RT_DELETE = RT_DELETE
        lt_loaded = lt_loaded
        PTHREAD_CANCEL_DISABLE = PTHREAD_CANCEL_DISABLE
        TD_CATCHSIG = TD_CATCHSIG
        TD_PRI_INHERIT = TD_PRI_INHERIT
        cet_always_off = cet_always_off
        TD_LOCK_TRY = TD_LOCK_TRY
        sigall_set = {__val = {18446744073709551615 <repeats 16 times>}}
        ns_s_pr = ns_s_an
        TD_REAP = TD_REAP
        _bitindex_arch_Slow_BSF = _bitindex_arch_Slow_BSF
        RSEQ_CS_FLAG_NO_RESTART_ON_SIGNAL_BIT = RSEQ_CS_FLAG_NO_RESTART_ON_SIGNAL_BIT
        _bitindex_arch_Prefer_PMINUB_for_stringop = _bitindex_arch_Prefer_PMINUB_for_stringop
        cpuid_register_index_eax = cpuid_register_index_eax
        _thread_db_pthread_eventbuf_eventmask_event_bits = {32, 2, 1616}
        _thread_db_rtld_global__dl_tls_dtv_slotinfo_list = {64, 1, 4176}
        _thread_db_pthread_schedpolicy = {32, 1, 1596}
        _thread_db___pthread_keys = {128, 1024, 0}
        _thread_db_const_thread_area = 25
        _thread_db_pthread_eventbuf = {192, 1, 1616}
        _thread_db___nptl_initial_report_events = {8, 1, 0}
        _thread_db_sizeof_list_t = 16
        _thread_db_link_map_l_tls_offset = {64, 1, 1112}
        _thread_db_sizeof_td_eventbuf_t = 24
        _thread_db_link_map_l_tls_modid = {64, 1, 1120}
        __nptl_rtld_global = 0x7ffff7ffd040 <_rtld_global>
        _thread_db_pthread_tid = {32, 1, 720}
        _thread_db_pthread_eventbuf_eventmask = {64, 1, 1616}
        _thread_db_pthread_nextevent = {64, 1, 1640}
        _thread_db_td_eventbuf_t_eventdata = {64, 1, 16}
        _thread_db___nptl_nthreads = {32, 1, 0}
        _thread_db_dtv_slotinfo_list_slotinfo = {128, 0, 16}
        _thread_db_pthread_list = {128, 1, 704}
        _thread_db_sizeof_pthread_key_data_level2 = 512
        _thread_db_list_t_prev = {64, 1, 8}
        _thread_db_td_thr_events_t_event_bits = {32, 2, 0}
        __GI___nptl_threads_events = {event_bits = {0, 0}}
        _thread_db_pthread_start_routine = {64, 1, 1600}
        _thread_db_pthread_schedparam_sched_priority = {32, 1, 1592}
        _thread_db_pthread_specific = {2048, 1, 1296}
        _thread_db_td_eventbuf_t_eventnum = {32, 1, 8}
        _thread_db_dtv_dtv = {128, 134217727, 0}
        _thread_db_sizeof_pthread = 2496
        _thread_db_sizeof_td_thr_events_t = 8
        _thread_db_pthread_key_data_level2_data = {128, 32, 0}
        _thread_db_pthread_report_events = {8, 1, 1553}
        _thread_db_pthread_cancelhandling = {32, 1, 776}
        _thread_db_list_t_next = {64, 1, 0}
        __GI___nptl_last_event = 0x0
        __nptl_version = "2.35"
        _thread_db_rtld_global__dl_stack_user = {128, 1, 4248}
        _thread_db_rtld_global__dl_stack_used = {128, 1, 4232}
#5  0x0000000000000016 in ?? ()
No symbol table info available.
#6  0x000000c0000069c0 in ?? ()
No symbol table info available.
#7  0x000000000045f6e0 in ?? ()
No locals.
#8  0x00007fffa3fff640 in ?? ()
No symbol table info available.
#9  0x0000000000465229 in threadentry (v=<optimized out>) at gcc_linux_amd64.c:94
        ts = {g = 0xc0000069c0, tls = <optimized out>, fn = 0x7fffa3ffedf0}
#10 0x00007ffff7d09b43 in start_thread (arg=<optimized out>) at ./nptl/pthread_create.c:442
        ret = <optimized out>
        pd = <optimized out>
        out = <optimized out>
        unwind_buf = {cancel_jmp_buf = {{jmp_buf = {140737488343712, -5239290227621810058, 140735944848960, 22, 140737351030864, 140737488344064, 5239105509183533174, 5239272881144148086}, mask_was_saved = 0}}, priv = {pad = {0x0, 0x0, 0x0, 0x0}, data = {prev = 0x0, cleanup = 0x0, canceltype = 0}}}
        not_first_call = <optimized out>
#11 0x00007ffff7d9ba00 in clone3 () at ../sysdeps/unix/sysv/linux/x86_64/clone3.S:81
No locals.

Thread 2 (Thread 0x7ffff7c71640 (LWP 4121) "main"):
#0  runtime.cgoSigtramp () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/sys_linux_amd64.s:390
No locals.
#1  <signal handler called>
No locals.
#2  __GI___pthread_sigmask (how=how@entry=2, newmask=<optimized out>, newmask@entry=0x7ffff7c70c70, oldmask=oldmask@entry=0x0) at ./nptl/pthread_sigmask.c:43
        local_newmask = {__val = {140737350404920, 1000000, 140737350405024, 140737219922864, 4608512, 140737350405040, 1, 4608425, 140737352943333, 140737350404944, 0, 140737352826881, 140737350405144, 140737350405496, 140737350405040, 140737219922864}}
        result = 0
#3  0x0000000000465396 in _cgo_sys_thread_start (ts=<optimized out>) at gcc_linux_amd64.c:77
        attr = {__size = '\000' <repeats 17 times>, "\020", '\000' <repeats 37 times>, __align = 0}
        ign = {__val = {18446744067267100671, 140737350405232, 0, 32, 1924145375264, 416, 8589934984, 24, 0, 2, 140737315148040, 416, 0, 206158430210, 0, 0}}
        oset = {__val = {0, 511101108348, 390842024046, 140737353344672, 1099338457352, 18446744073709551528, 24, 140737219919920, 44, 140737353652384, 1, 140737351098809, 140737352856103, 392, 140737350405496, 140737350405520}}
        p = 140735928067648
        size = 16777216
        err = 0
#4  0x00007ffff7f119c0 in runtime.asmcgocall () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/asm_amd64.s:878
No locals.
#5  0x00007ffff7ec7227 in runtime.newobject (typ=<optimized out>, ~r0=<optimized out>) at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/malloc.go:1254
No locals.
#6  0x00007ffff7c70d58 in ?? ()
No symbol table info available.
#7  0x0000000000000000 in ?? ()
No symbol table info available.

Thread 1 (Thread 0x7ffff7c72740 (LWP 4118) "main"):
#0  clone3 () at ../sysdeps/unix/sysv/linux/x86_64/clone3.S:62
No locals.
#1  0x00007ffff7d9ba51 in __GI___clone_internal (cl_args=cl_args@entry=0x7fffffffd210, func=func@entry=0x7ffff7d09850 <start_thread>, arg=arg@entry=0x7fffa1ffd640) at ../sysdeps/unix/sysv/linux/clone-internal.c:54
        ret = <optimized out>
        saved_errno = 22
        flags = <optimized out>
        stack = <optimized out>
#2  0x00007ffff7d09759 in create_thread (pd=pd@entry=0x7fffa1ffd640, attr=attr@entry=0x7fffffffd4c0, stopped_start=stopped_start@entry=0x7fffffffd32e, stackaddr=stackaddr@entry=0x7fffa0ffd000, stacksize=16776960, thread_ran=thread_ran@entry=0x7fffffffd32f) at ./nptl/pthread_create.c:295
        need_setaffinity = <optimized out>
        clone_flags = 4001536
        tp = 0x7fffa1ffd640
        args = {flags = 4001536, pidfd = 140735911287056, child_tid = 140735911287056, parent_tid = 140735911287056, exit_signal = 0, stack = 140735894507520, stack_size = 16776960, tls = 140735911286336, set_tid = 0, set_tid_size = 0, cgroup = 0}
        ret = <optimized out>
        __PRETTY_FUNCTION__ = "create_thread"
#3  0x00007ffff7d0a280 in __pthread_create_2_1 (newthread=newthread@entry=0x7fffffffd4b0, attr=attr@entry=0x7fffffffd4c0, start_routine=start_routine@entry=0x465200 <threadentry>, arg=arg@entry=0x4ff580) at ./nptl/pthread_create.c:828
        stackaddr = 0x7fffa0ffd000
        stacksize = <optimized out>
        iattr = 0x7fffffffd4c0
        default_attr = {external = {__size = "\234VA\000\000\000\000\000\257\376B\000\000\000\000\000H\317\303\366\377\177\000\000\000 \000\000\000\000\000\000\000\200\003\000\000\000\000\000\b\324\377\377\377\177\000\000\323eB\000\000\000\000", __align = 4282012}, internal = {schedparam = {sched_priority = 4282012}, schedpolicy = 0, flags = 4390575, guardsize = 140737333415752, stackaddr = 0x2000, stacksize = 229376, extension = 0x7fffffffd408, unused = 0x4265d3 <runtime.(*mheap).allocSpan+1395>}}
        destroy_default_attr = <optimized out>
        c11 = <optimized out>
        pd = 0x7fffa1ffd640
        err = 0
        retval = 0
        self = <optimized out>
        stopped_start = false
        thread_ran = false
        original_sigmask = {__val = {18446744067266838271, 2048, 16781312, 1, 576460752320204800, 3, 4283704, 0, 4348621, 4344149, 0, 824635817984, 140737488344064, 4283463, 824634114048, 32768}}
        __PRETTY_FUNCTION__ = "__pthread_create_2_1"
#4  0x0000000000465151 in _cgo_try_pthread_create (thread=thread@entry=0x7fffffffd4b0, attr=attr@entry=0x7fffffffd4c0, pfn=pfn@entry=0x465200 <threadentry>, arg=arg@entry=0x4ff580) at gcc_libinit.c:100
        tries = 0
        err = <optimized out>
        ts = {tv_sec = 6, tv_nsec = 0}
#5  0x0000000000465384 in _cgo_sys_thread_start (ts=0x4ff580) at gcc_linux_amd64.c:75
        attr = {__size = '\000' <repeats 17 times>, "\020", '\000' <repeats 37 times>, __align = 0}
        ign = {__val = {18446744067267100671, 72198331526272288, 0, 32, 140737488344472, 47, 8589934612, 24, 0, 824633748992, 197568776192, 4248989, 0, 206158430210, 0, 0}}
        oset = {__val = {0, 511101108348, 390842024046, 4668928, 197572948165, 18446744073709551528, 24, 140737352625280, 0, 5035840, 8240, 140737351098809, 140737488344584, 4251847, 140737488344712, 140737488344736}}
        p = 140735911286336
        size = 16777216
        err = <optimized out>
#6  0x000000000045f9e1 in runtime.asmcgocall () at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/asm_amd64.s:878
No locals.
#7  0x000000c000044800 in ?? ()
No symbol table info available.
#8  0x000000c000044800 in ?? ()
No symbol table info available.
#9  0x0000000000446748 in runtime.(*rwmutex).runlock (rw=0x4fb280 <runtime[allocmLock]>) at /opt/hostedtoolcache/go/1.20.5/x64/src/runtime/rwmutex.go:78
No locals.
#10 0x0000000000000000 in ?? ()
No symbol table info available.
```
