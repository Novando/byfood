go build -o ./bin/byfood ./cmd/main.go

GVM_PATH=
TZ=Asia/Tokyo

if [ $GVM_TAG ]; then
  GVM_PATH=~/.gvm/gos/$GVM_TAG/bin/;
fi

if [ $TZ ]; then
  TZ=$TZ;
fi

if [ $ENV ]; then
  GOOS=linux GOARCH=amd64 ${GVM_PATH}go build \
    -o ./bin \
    --ldflags="-X 'gitlab.com/novando/byfood/be/pkg/env.ENV=${ENV}'" \
    --ldflags="-X 'gitlab.com/novando/byfood/be/pkg/env.TZ=${TZ}'" \
    ./cmd/main.go;
else
  echo "ENV param not declared, using 'local' as default";
  GOOS=linux GOARCH=amd64 ${GVM_PATH}go build \
  -o ./bin \
  --ldflags="-X 'gitlab.com/novando/byfood/be/pkg/env.TZ=${TZ}'" \
  ./cmd/main.go;
fi