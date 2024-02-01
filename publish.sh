#!/usr/bin/env bash


set -e
ROOT=$(cd "$(dirname "$0")";cd .;pwd)


########################################################################################################################
# 填写你的信息
########################################################################################################################
HOST=ziji.fun


bash build_linux.sh


scp -r -v ip-reporter root@${HOST}:/ziji/oss/bin/


ssh root@${HOST} << 'EOF'
echo '0 * * * * /ziji/oss/bin/ip-reporter >> /ziji/oss/bin/ip-reporter.log.$(date +\%Y\%m\%d) 2>&1' >> /var/spool/cron/crontabs/root
exit
EOF
