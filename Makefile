.PHONY: migrate_db prepopulate_db start_db start_prepopulated_db start_server unit_tests run_integration_tests gen_doc run_doc_server gen_doc_and_run_server

migrate_db:
	ENV_FILE="${ENV_FILE}" \
	./scripts/migrate_db.sh

start_db:
	ENV_FILE="${ENV_FILE}" \
	./scripts/start_db.sh

start_prepopulated_db: ENV_FILE := cli-env
start_prepopulated_db: | start_db migrate_db

stop_db:
	ENV_FILE="${ENV_FILE}" \
	./scripts/stop_db.sh

start_server: ENV_FILE := cli-env
start_server:
	ENV_FILE="${ENV_FILE}" \
	./scripts/start_server.sh

unit_tests:
	EXTRA_ARGS="" \
	./scripts/tests.sh

start_db_and_run_integration_tests: ENV_FILE := cli-test-env
start_db_and_run_integration_tests: | start_db migrate_db run_integration_tests

run_integration_tests: ENV_FILE := cli-test-env
run_integration_tests:
	ENV_FILE="${ENV_FILE}" \
    ./scripts/integration_tests.sh
