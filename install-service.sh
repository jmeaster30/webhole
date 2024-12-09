SYSTEMD_INSTALL_DIR=~/.config/systemd/user
WEB_HOLE_HOME=~/.webhole
echo $WEB_HOLE_HOME
echo $SYSTEMD_INSTALL_DIR

set +a
echo "[Unit]"                                          > "webhole.service"
echo "Description=Web Hole"                            >> "webhole.service"
echo "After=network.target"                            >> "webhole.service"
echo ""                                                >> "webhole.service"
echo "[Service]"                                       >> "webhole.service"
echo "Type=notify"                                     >> "webhole.service"
echo "Restart=on-failure"                              >> "webhole.service"
echo "WorkingDirectory=$(readlink -f $WEB_HOLE_HOME)"  >> "webhole.service"
echo "ExecStart=$(readlink -f $WEB_HOLE_HOME)/webhole" >> "webhole.service"
echo ""                                                >> "webhole.service"
echo "[Install]"                                       >> "webhole.service"
echo "WantedBy=multi-user.target"                      >> "webhole.service"
set -a

mkdir -p $WEB_HOLE_HOME
mkdir -p $SYSTEMD_INSTALL_DIR
cp ./*.go ./go.mod ./go.sum $WEB_HOLE_HOME
cp ./webhole.service "$SYSTEMD_INSTALL_DIR/webhole.service"
pushd $WEB_HOLE_HOME
go build -v -o webhole
popd
systemctl --user daemon-reload
systemctl --user enable webhole
systemctl --user start webhole
systemctl --user --no-pager --full status webhole.service
