## Process Management
1. Make `~/.config/systemd/user/<SERVICE_NAME>.service`
    ```
    [Unit]

    [Install]
    WantedBy=multi-user.target

    [Service]
    ExecStart=<BINARY_PROGRAM>
    WorkingDirectory=/home/<username>/<MY_GO_APP_HOME_DIR>
    Restart=always
    RestartSec=5
    StandardOutput=syslog
    StandardError=syslog
    SyslogIdentifier=%n
    ```

1. Configure services
    1. Reload services
        ```bash
        systemctl daemon-reload
        ```

    1. Start service (if not started yet)
        ```bash
        systemctl --user start <SERVICE_NAME>.service
        ```

    1. Check service status
        ```bash
        systemctl --user status <SERVICE_NAME>
        ```
