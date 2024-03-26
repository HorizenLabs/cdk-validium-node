#!/bin/bash
set -eEuo pipefail

# check required env vars are set
vars_to_check=(
  "NGINX_CERT"
  "NGINX_CERT_KEY"
  "NGINX_DOMAIN"
  "EXPLORER_NET_IP_ADDRESS"
  "EXPLORER_NET_UI_PORT"
)

for var in "${vars_to_check[@]}"; do
  if [ -z "${!var:-}" ]; then
    echo "Error: Environment variable ${var} is required. Exiting ..."
    sleep 5
    exit 1
  fi
done

# install requirements
command -v openssl > /dev/null 2>&1 || apk add --no-cache openssl
mkdir -p /etc/nginx/dynamic-conf/certificates /etc/nginx/dynamic-conf/conf.d /etc/nginx/dynamic-conf/snippets

# check validity of cloudflare authenticated origin pull cert, update if outdated
renew="false"
if [ -s "/etc/nginx/dynamic-conf/certificates/origin-pull-ca.pem" ]; then
  expires="$(openssl x509 -in /etc/nginx/dynamic-conf/certificates/origin-pull-ca.pem -noout -enddate | sed -e 's#notAfter=##' | rev | cut -d " " -f 2- | rev | xargs -I {} date --date="{}" +%s)"
  now="$(date +%s)"
  # valid for less than 90 days
  if [ "$((expires-now))" -lt 7776000 ]; then
    rm -f /etc/nginx/dynamic-conf/certificates/origin-pull-ca.pem
    renew="true"
  fi
else
  renew="true"
fi

if [ "${renew}" = "true" ]; then
  wget -q https://developers.cloudflare.com/ssl/static/authenticated_origin_pull_ca.pem -O /etc/nginx/dynamic-conf/certificates/origin-pull-ca.pem || { echo 'failed to get origin-pull-ca.pem' && false; }
  grep -q "END CERTIFICATE" /etc/nginx/dynamic-conf/certificates/origin-pull-ca.pem || { echo 'Fatal, does not look like a certificate from cloudflare, failed to get origin-pull-ca.pem!' && rm -f /etc/nginx/dynamic-conf/certificates/origin-pull-ca.pem && false; }
fi

# generate snakeoil certs if not exist
if ! [ -s "/etc/nginx/dynamic-conf/certificates/snakeoil.key" ] || ! [ -s "/etc/nginx/dynamic-conf/certificates/snakeoil.pem" ]; then
  openssl req -newkey rsa:4096 -new -x509 -days 3650 -nodes -out /etc/nginx/dynamic-conf/certificates/snakeoil.pem -keyout /etc/nginx/dynamic-conf/certificates/snakeoil.key -subj "/C=US/ST=South Carolina/L=Columbia/O=IT/CN=snakeoil.localdomain" -extensions v3_ca
  chmod 400 /etc/nginx/dynamic-conf/certificates/snakeoil.key
fi

# generate DH parameters if not exist
if ! grep -q "PARAMETERS" /etc/nginx/dynamic-conf/dhparam.pem > /dev/null 2>&1; then
  openssl dhparam 4096 > /etc/nginx/dynamic-conf/dhparam.pem
fi

# generate cloudflare-whitelist.conf and cloudflare-restore-origin-ip.conf
NGINX_WHITELIST="/etc/nginx/dynamic-conf/snippets/cloudflare-whitelist.conf"
NGINX_ORIGIN="/etc/nginx/dynamic-conf/snippets/cloudflare-restore-origin-ip.conf"
CFDLDIR="$(mktemp -d)"
cd "${CFDLDIR}"
wget -q https://www.cloudflare.com/ips-v4 || { echo 'failed to get cloudflare IPv4 Servers' && rm -rf "${CFDLDIR}" && false; }
wget -q https://www.cloudflare.com/ips-v6 || { echo 'failed to get cloudflare IPv6 Servers' && rm -rf "${CFDLDIR}" && false; }
grep -qv html ips-v4 || { echo 'Fatal, got html instead of IP list from cloudflare, failed to generate nginx/dynamic-conf/cloudflare-*.conf!' && rm -r "${CFDLDIR}" && false; }
grep -qv html ips-v6 || { echo 'Fatal, got html instead of IP list from cloudflare, failed to generate nginx/dynamic-conf/cloudflare-*.conf!' && rm -r "${CFDLDIR}" && false; }
echo "geo \$realip_remote_addr \$cloudflare_ip {" > "${NGINX_WHITELIST}"
echo -e "\tdefault\t\t0;" >> "${NGINX_WHITELIST}"
cat ips-v4 | perl -p -e "s/^(.*)$/\t\1\t1;/g" >> "${NGINX_WHITELIST}"
echo "" >> "${NGINX_WHITELIST}"
cat ips-v6 | perl -p -e "s/^(.*)$/\t\1\t1;/g" >> "${NGINX_WHITELIST}"
echo "" >> "${NGINX_WHITELIST}"
echo "}" >> "${NGINX_WHITELIST}"
echo "# CloudFlare IP Ranges" > "${NGINX_ORIGIN}"
echo "# Generated at $(date) by $0" >> "${NGINX_ORIGIN}"
echo "" >> "${NGINX_ORIGIN}"
echo "# - IPv4 (https://www.cloudflare.com/ips-v4)" >> "${NGINX_ORIGIN}"
cat ips-v4 | perl -p -e "s/^(.*)$/set_real_ip_from\ \1;/g" >> "${NGINX_ORIGIN}"
echo "" >> "${NGINX_ORIGIN}"
echo "# - IPv6 (https://www.cloudflare.com/ips-v6)" >> "${NGINX_ORIGIN}"
cat ips-v6 | perl -p -e "s/^(.*)$/set_real_ip_from\ \1;/g" >> "${NGINX_ORIGIN}"
echo "" >> "${NGINX_ORIGIN}"
echo "real_ip_header CF-Connecting-IP;" >> "${NGINX_ORIGIN}"
echo "" >> "${NGINX_ORIGIN}"
rm -rf "${CFDLDIR}"
unset CFDLDIR

# add text/csv mime type
if ! grep -q csv /usr/local/openresty/nginx/conf/mime.types; then
  sed -i '/}/i \ \ \ \ text/csv                                         csv;' /usr/local/openresty/nginx/conf/mime.types
fi

# generate .htpasswd files
NGINX_AUTHBLOCK=""
if [ -n "${NGINX_HTPASSWD:-}" ]; then
  touch /etc/nginx/.htpasswd
  user="$(echo "${NGINX_HTPASSWD}" | cut -d : -f 1)"
  pass="$(echo "${NGINX_HTPASSWD}" | cut -d : -f 2)"
  salt="gesalzen"
  echo "${user}:$(mkpasswd -m sha512 "$pass" "$salt")" > /etc/nginx/.htpasswd
  # shellcheck disable=SC2089
  NGINX_AUTHBLOCK='
        auth_basic "Login";
        auth_basic_user_file /etc/nginx/.htpasswd;'
  chown "$(id -u nobody)":"$(id -g nobody)" /etc/nginx/.htpasswd
fi
# shellcheck disable=SC2090
export NGINX_AUTHBLOCK

# write out certificate and key files
# no arrays or indirection in ash, so done for each cert
echo -e "${NGINX_CERT}" > "/etc/nginx/dynamic-conf/certificates/domain-origin.crt"
echo -e "${NGINX_CERT_KEY}" > "/etc/nginx/dynamic-conf/certificates/domain-origin.key"
chmod 400 "/etc/nginx/dynamic-conf/certificates/domain-origin.crt" "/etc/nginx/dynamic-conf/certificates/domain-origin.key"

# use envsubst to substitute variables in default.conf and healthcheck.conf
# shellcheck disable=SC2016
NGINX_ENVSUBST='$NGINX_AUTHBLOCK:$NGINX_CERT:$NGINX_CERT_KEY:$NGINX_DOMAIN:$EXPLORER_NET_IP_ADDRESS:$EXPLORER_NET_UI_PORT'
export NGINX_ENVSUBST

envsubst "${NGINX_ENVSUBST}" < /etc/nginx/templates/default.conf.tmpl > /etc/nginx/dynamic-conf/conf.d/default.conf

# check config and sleep until valid (backend might not be up yet)
while ! nginx -t; do
  sleep 10
done

# shellcheck disable=SC2046
unset $(env | grep "NGINX_" | awk -F'=' '{print $1}')

exec "$@"