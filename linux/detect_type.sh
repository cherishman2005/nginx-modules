#!/bin/sh

detect_container()
{
    if which systemd-detect-virt >/dev/null 2>&1; then
        TYPE=$(systemd-detect-virt -c)
        if [ "$TYPE" = "none" ]; then
            return 1
        else
            echo "Container: $TYPE"
            return 0
        fi
    fi
    if [ -n "$container" ]; then
        echo "Container: $container"
        return 0
    fi
    if grep -qi docker /proc/1/cgroup; then
        echo "Container: Docker"
        return 0
    fi
    if test -f /.dockerenv; then
        echo "Container: Docker"
        return 0
    fi
    if grep -qi 'machine-rkt' /proc/1/cgroup; then
        echo "Container: rkt"
        return 0
    fi
    # Other container type detect here
    return 1
}

detect_physical()
{
    if ! lscpu | grep -qi 'Hypervisor vendor'; then
        echo "Physical: $(cat /sys/class/dmi/id/product_name)"
        return 0
    fi
    return 1
}

detect_virtual_machine()
{
    if lscpu | grep -qi 'Hypervisor vendor'; then
        HYPER_TYPE=$(lscpu | grep -i "Hypervisor vendor" \
            | cut -d ':' -f 2 | sed 's/^ *//g')
        if dmidecode -t system | grep -qi 'amazon'; then
            echo "Virtual Machine: AWS/$HYPER_TYPE"
        elif dmidecode -t system | grep -qi 'openstack'; then
            echo "Virtual Machine: OpenStack/$HYPER_TYPE"
        elif dmidecode -t system | grep -qi 'alibaba'; then
            echo "Virtual Machine: Aliyun/$HYPER_TYPE"
        else
            Manufacturer=$(dmidecode -t system | grep 'Manufacturer' \
                | cut -d ':' -f 2 | sed 's/^ *//g')
            ProductName=$(dmidecode -t system | grep 'Product Name' \
                | cut -d ':' -f 2 | sed 's/^ *//g')
            Version=$(dmidecode -t system | grep 'Version' \
                | cut -d ':' -f 2 | sed 's/^ *//g')
            echo "Virtual Machine: $Manufacturer $ProductName($Version)/$HYPER_TYPE"
        fi
        return 0
    fi
    return 1
}

detect_virtual_type()
{
    detect_container || detect_physical \
    || detect_virtual_machine || echo "Unknown"
}

detect_virtual_type "$@"