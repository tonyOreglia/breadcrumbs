# Running breadcrumbs as a systemd service on Ubuntu 

1. Update `breadcrumbs.service` ConditionPathExists, User, Group, WorkingDirectory, and ExecStart as needed.
1. Move `breadcrumbs.service` to `/lib/systemd/system/glee.service`
    ```
    $ sudo cp config/systemctl/breadcrumbs.service /lib/systemd/system/
    ```
1. Update the file permissions
    ```
    $ sudo chmod 755 /lib/systemd/system/breadcrumbs.service
    ```
1. Start the service
    ```
    $ sudo systemctl breadcrumbs start
    $ // if this fails with 'unknown operation' try `sudo systemctl start breadcrumbs`
    ```
1. Monitor the service ouput
    ```
    $ journalctl -f -u breadcrumbs
    ```
1. Enable the service to start at boot
    ```
    $ sudo systemctl enable breadcrumbs
