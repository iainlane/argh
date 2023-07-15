#include <check.h>
#include <stdlib.h>
#include <stdio.h>

#include "callgofunction.h"

START_TEST(test_it_works)
{
    GoInt result = call_go_function();

    ck_assert_int_eq(result, 1337);
}
END_TEST

Suite *solver_suite(void)
{
    Suite *s;
    TCase *tc_core;

    s = suite_create("c_api");

    /* Core test case */
    tc_core = tcase_create("Core");

    tcase_add_test(tc_core, test_it_works);
    suite_add_tcase(s, tc_core);

    return s;
}

int main(void)
{
    int number_failed;
    Suite *s;
    SRunner *sr;

    s = solver_suite();
    sr = srunner_create(s);

    srunner_set_fork_status(sr, CK_NOFORK);
    srunner_run_all(sr, CK_NORMAL);
    number_failed = srunner_ntests_failed(sr);
    srunner_free(sr);

    return (number_failed == 0) ? EXIT_SUCCESS : EXIT_FAILURE;
}
