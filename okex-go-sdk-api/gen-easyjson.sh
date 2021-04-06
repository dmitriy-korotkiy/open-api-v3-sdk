#!/bin/bash

# go get -u github.com/mailru/easyjson/...

filesForJSON=(
"client.go"
"cursor_page.go"
"futures_params.go"
"futures_results.go"
"req_with_models.go"
"swap_params.go"
"swap_results.go"
"utils.go"
"ws_base.go"
)

echo "Start easyjson"
# shellcheck disable=SC2068
easyjson ${filesForJSON[@]}
