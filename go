#!/bin/sh

IMAGE="elixir:1.8.1-alpine"
PWD=$(pwd)
NAME=$(basename "${PWD}")

if [ -z "${PORT}" ]; then
	PORTSTR=""
else
	echo "--- using port: ${PORT}"
	PORTSTR="-p ${PORT}:${PORT}"
fi

ENVFILE="./.env"

touch "${ENVFILE}"

docker run \
	-it \
	--rm=true \
	--name "${NAME}_iex" \
	--env-file "${ENVFILE}" \
	${PORTSTR} \
	-v "${PWD}:/app" \
	-v "${PWD}/elixir-mix:/root/.mix" \
	-w /app \
	"${IMAGE}" $@
