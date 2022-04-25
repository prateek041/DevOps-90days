### Linux filesystems:

- it is a method of storing and organising collection of data in a human usable form.
- Conventional disk : ext3, ext4 etc.
- flash storage: ubits, yaffs etc.
- special purpose: fuse, procfs etc.

### Partition and filesystem:

- partition: A physically continuous section of a disk.
- Filesystem: filesystem is used to store files on a partition or hard disk.

### FHS (Filesystem heirarchy standard):

- Standard layout of storing important linux system files.
- very easily navigatable.
- '/' character is used to seperate paths.
- Removable media is mounted to "/run/media/username/disklabel", for example, if you a pendrive named DevOps. It will be found as /run/media/prateek/DevOps.


## Chapter 4: Graphical interface.

### X windows system:

- it is loaded as one of the final steps in the boot process.
- X server provides the graphical services to applications.
- X clients are the receivers of the graphical services.
- Display manager performs display management, loads X server and also manages graphical logins.
- now Wayland is working under the hood, instead of X.
