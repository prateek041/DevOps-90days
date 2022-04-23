# Chapter 3: Linux Basics and System Startup

## The Boot Process
1. BIOS
    - Starting an x86-based Linux system involves a number of steps. When the computer is powered on, the Basic Input/Output System (BIOS) initializes the hardware, including the screen and keyboard, and tests the main memory. This process is also called POST (Power On Self Test)
    - The BIOS software is stored on a ROM chip on the motherboard. After this, the remainder of the boot process is controlled by the operating system (OS).
2. Master Boot Record (MBR) and Boot Loader
    - Once the POST is completed, the system control passes from the BIOS to the boot loader. The boot loader is usually stored on one of the hard disks in the system, either in the boot sector (for traditional BIOS/MBR systems) or the EFI partition (for more recent (Unified) Extensible Firmware Interface or EFI/UEFI systems).
    - A number of boot loaders exist for Linux; the most common ones are GRUB (for GRand Unified Boot loader), ISOLINUX (for booting from removable media), and DAS U-Boot (for booting on embedded devices/appliances).
    - When booting Linux, the boot loader is responsible for loading the kernel image and the initial RAM disk or filesystem (which contains some critical files and device drivers needed to start the system) into memory.
3. Initial RAM Disk
    - The initramfs filesystem image contains programs and binary files that perform all actions needed to mount the proper root filesystem, like providing kernel functionality for the needed filesystem and device drivers for mass storage controllers with a facility called udev (for user device), which is responsible for figuring out which devices are present, locating the device drivers they need to operate properly, and loading them. After the root filesystem has been found, it is checked for errors and mounted.
    - The mount program instructs the operating system that a filesystem is ready for use, and associates it with a particular point in the overall hierarchy of the filesystem (the mount point). If this is successful, the initramfs is cleared from RAM and the init program on the root filesystem (/sbin/init) is executed.
    - init handles the mounting and pivoting over to the final real root filesystem. If special hardware drivers are needed before the mass storage can be accessed, they must be in the initramfs image.
4. Text-Mode Login
    - Near the end of the boot process, init starts a number of text-mode login prompts. These enable you to type your username, followed by your password, and to eventually get a command shell.

## Kernel, init, and services
1. The Linux Kernel
    - The boot loader loads both the kernel and an initial RAM–based file system (initramfs) into memory, so it can be used directly by the kernel.
    - When the kernel is loaded in RAM, it immediately initializes and configures the computer’s memory and also configures all the hardware attached to the system. This includes all processors, I/O subsystems, storage devices, etc. The kernel also loads some necessary user space applications.
2. /sbin/init and Services
    - Once the kernel has set up all its hardware and mounted the root filesystem, the kernel runs /sbin/init. This then becomes the initial process, which then starts other processes to get the system running.
    - Besides starting the system, init is responsible for keeping the system running and for shutting it down cleanly. One of its responsibilities is to act when necessary as a manager for all non-kernel processes
3. systemd features
    - Systems with systemd start up faster than those with earlier init methods. This is largely because it replaces a serialized set of steps with aggressive parallelization techniques, which permits multiple services to be initiated simultaneously.
    - Complicated startup shell scripts are replaced with simpler configuration files, which enumerate what has to be done before a service is started, how to execute service startup, and what conditions the service should indicate have been accomplished when startup is finished.
    - /sbin/init now just points to /lib/systemd/systemd; i.e. systemd takes over the init process.
4. systemctl commands
```sh
# Starting, stopping, restarting a service (using nfs as an example) on a currently running system:
$ sudo systemctl start|stop|restart nfs.service
# Enabling or disabling a system service from starting up at system boot:
$ sudo systemctl enable|disable nfs.service
```

## Linux Filesystems
1. Partitions and Filesystems
    - A partition is a physically contiguous section of a disk, or what appears to be so in some advanced setups.
    - A filesystem is a method of storing/finding files on a hard disk (usually in a partition).
    - One can think of a partition as a container in which a filesystem resides
2. The Filesystem Hierarchy Standard
    - Linux systems store their important files according to a standard layout called the Filesystem Hierarchy Standard (FHS)
    - Multiple drives and/or partitions are mounted as directories in the single filesystem.
3. Directory tree
    - `/bin/` Essential user command binaries
    - `/boot/` Static files of the boot loader
    - `/dev/` Device files
    - `/etc/` Host-specific system configuration
    - `/home/` User home directories
    - `/lib/` Essential shared libraries and kernel modules
    - `/media/` Mount point for removable media
    - `/mnt/` Mount point for a temporarily mounted filesystem
    - `/opt/` Add-on application software packages
    - `/sbin/` System binaries
    - `/srv/` Data for services provided by this system
    - `/tmp/` Temporary files
    - `/usr/` (Multi-)User utilities and applications
    - `/var/` Variable files
    - `/root/` Home directory for the root user
    - `/proc/` Virtual filesystem documenting kernel and process status as text files

## Summary
- A **partition** is a logical part of the disk.
- A **filesystem** is a method of storing/finding files on a hard disk.
- By dividing the hard disk into partitions, data can be grouped and separated as needed. When a failure or mistake occurs, only the data in the affected partition will be damaged, while the data on the other partitions will likely survive.
- The boot process has multiple steps, starting with BIOS, which triggers the boot loader to start up the Linux kernel. From there, the initramfs filesystem is invoked, which triggers the init program to complete the startup process.
- Determining the appropriate distribution to deploy requires that you match your specific system needs to the capabilities of the different distributions.
