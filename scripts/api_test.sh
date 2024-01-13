#!/bin/sh

AUTH_SERVER="127.0.0.1:9999"
API_SERVER="127.0.0.1:7777"
COOKIE_FILE="/tmp/$(date +'%Y%m%d%H%M%S').cookie"

curl -c "${COOKIE_FILE}" --silent -X POST http://${AUTH_SERVER}/api/login --data '{"username": "test", "password": "passwd"}'
echo ""
cat "${COOKIE_FILE}"
echo ""

echo "NG patter test:"
curl -v http://${API_SERVER}/api/users
echo ""

echo ""
echo "OK pattern test:"
curl -v http://${API_SERVER}/api/users -b "${COOKIE_FILE}"
echo ""

rm -f "${COOKIE_FILE}"
