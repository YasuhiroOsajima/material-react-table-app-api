#!/bin/sh

AUTH_SERVER="127.0.0.1:9999"
API_SERVER="127.0.0.1:7777"
COOKIE_FILE="/tmp/$(date +'%Y%m%d%H%M%S').cookie"

curl -c "${COOKIE_FILE}" --silent -X POST http://${AUTH_SERVER}/api/login --data '{"username": "test", "password": "passwd"}'
echo ""
cat "${COOKIE_FILE}"
echo ""

echo "Not found patter test:"
curl -v -X DELETE http://${API_SERVER}/api/users/test -b "${COOKIE_FILE}"
echo ""

echo ""
echo "OK pattern test:"
curl -v -X DELETE http://${API_SERVER}/api/users/Alice -b "${COOKIE_FILE}"
echo ""

rm -f "${COOKIE_FILE}"
