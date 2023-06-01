#!/bin/sh
pnpm run build
pnpm run serve --port 37628 &
SERV_PID=$!

cleanup() {
  kill $SERV_PID
}
trap cleanup EXIT

sleep 10

mkdir -p spider && cd spider || exit
wget --spider -nv -nd -r -o wget.log -l 100 http://localhost:37628/workflow/introduction/overview

grep "broken links" wget.log
echo Check spider/wget.log
